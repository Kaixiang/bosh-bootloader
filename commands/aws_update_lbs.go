package commands

import (
	"errors"
	"io/ioutil"
	"strings"
	"time"

	"github.com/cloudfoundry/bosh-bootloader/storage"
)

type AWSUpdateLBs struct {
	certificateManager        certificateManager
	availabilityZoneRetriever availabilityZoneRetriever
	infrastructureManager     infrastructureManager
	credentialValidator       credentialValidator
	logger                    logger
	guidGenerator             guidGenerator
	stateStore                stateStore
	environmentValidator      environmentValidator
}

func NewAWSUpdateLBs(credentialValidator credentialValidator, certificateManager certificateManager,
	availabilityZoneRetriever availabilityZoneRetriever, infrastructureManager infrastructureManager,
	logger logger, guidGenerator guidGenerator, stateStore stateStore, environmentValidator environmentValidator) AWSUpdateLBs {

	return AWSUpdateLBs{
		credentialValidator:       credentialValidator,
		certificateManager:        certificateManager,
		availabilityZoneRetriever: availabilityZoneRetriever,
		infrastructureManager:     infrastructureManager,
		logger:                    logger,
		guidGenerator:             guidGenerator,
		stateStore:                stateStore,
		environmentValidator:      environmentValidator,
	}
}

func (c AWSUpdateLBs) Execute(config AWSCreateLBsConfig, state storage.State) error {
	err := c.credentialValidator.Validate()
	if err != nil {
		return err
	}

	err = c.environmentValidator.Validate(state)
	if err != nil {
		return err
	}

	if !lbExists(state.Stack.LBType) {
		return LBNotFound
	}

	if match, err := c.checkCertificateAndChain(config.CertPath, config.ChainPath, state.Stack.CertificateName); err != nil {
		return err
	} else if match {
		c.logger.Println("no updates are to be performed")
		return nil
	}

	c.logger.Step("uploading new certificate")

	certificateName, err := certificateNameFor(state.Stack.LBType, c.guidGenerator, state.EnvID)
	if err != nil {
		return err
	}

	err = c.certificateManager.Create(config.CertPath, config.KeyPath, config.ChainPath, certificateName)
	if err != nil {
		return err
	}

	// Temporary fix for IAM propagation. Terraform should have retry logic for this, so we should remove it once we start using terraform on AWS.
	time.Sleep(9 * time.Second)

	if err := c.updateStack(certificateName, state.KeyPair.Name, state.Stack.Name, state.Stack.BOSHAZ, state.Stack.LBType, state.AWS.Region, state.EnvID); err != nil {
		return err
	}

	c.logger.Step("deleting old certificate")
	err = c.certificateManager.Delete(state.Stack.CertificateName)
	if err != nil {
		return err
	}

	state.Stack.CertificateName = certificateName

	err = c.stateStore.Set(state)
	if err != nil {
		return err
	}

	return nil
}

func (c AWSUpdateLBs) checkCertificateAndChain(certPath string, chainPath string, oldCertName string) (bool, error) {
	localCertificate, err := ioutil.ReadFile(certPath)
	if err != nil {
		return false, err
	}

	remoteCertificate, err := c.certificateManager.Describe(oldCertName)
	if err != nil {
		return false, err
	}

	if strings.TrimSpace(string(localCertificate)) != strings.TrimSpace(remoteCertificate.Body) {
		return false, nil
	}

	if chainPath != "" {
		localChain, err := ioutil.ReadFile(chainPath)
		if err != nil {
			return false, err
		}

		if strings.TrimSpace(string(localChain)) != strings.TrimSpace(remoteCertificate.Chain) {
			return false, errors.New("you cannot change the chain after the lb has been created, please delete and re-create the lb with the chain")
		}
	}

	return true, nil
}

func (c AWSUpdateLBs) updateStack(certificateName string, keyPairName string, stackName string, boshAZ string, lbType string, awsRegion, envID string) error {
	availabilityZones, err := c.availabilityZoneRetriever.Retrieve(awsRegion)
	if err != nil {
		return err
	}

	certificate, err := c.certificateManager.Describe(certificateName)
	if err != nil {
		return err
	}

	_, err = c.infrastructureManager.Update(keyPairName, availabilityZones, stackName, boshAZ, lbType, certificate.ARN, envID)
	if err != nil {
		return err
	}

	return nil
}
