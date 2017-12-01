// Code generated by go-bindata.
// sources:
// templates/cf_lb.tf
// templates/network.tf
// templates/network_security_group.tf
// templates/output.tf
// templates/resource_group.tf
// templates/storage.tf
// templates/tls.tf
// templates/vars.tf
// DO NOT EDIT!

package azure

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesCf_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\xc1\x8e\xe3\x36\x0c\xbd\xe7\x2b\x04\xa1\x57\x7b\x92\xdd\x60\xd1\x8b\x2f\x45\x0f\x3d\x16\x18\x14\x3d\x0a\xb4\x2d\xc7\x42\x14\x49\x2b\x51\x99\x4d\x07\xfe\xf7\x42\x96\x9d\xc4\x8a\xed\x4d\x36\xb7\x98\x8f\x14\xf9\x1e\x49\xe9\x0c\x56\x40\x29\x39\xa1\xee\xe2\x90\x9f\x58\xad\x4f\x20\x14\x25\x9f\xdd\x66\x73\x33\x9a\xe6\x07\xab\xb8\x45\x56\x82\xe3\xdf\xf6\x73\x66\x03\xce\x7d\x68\x5b\x47\x9b\xe5\x4e\x7b\x5b\x71\x42\xe1\x3f\x6f\xb9\x3d\x31\xe7\x4b\xc5\x91\x12\xea\x7c\xb9\xa3\xe4\x73\x43\x88\x82\x13\x27\xe9\xaf\x20\xf4\xb7\xcf\x33\xd8\x9c\xab\x33\x13\x75\x97\x55\x4d\x16\x7d\x77\x74\x43\x08\xd4\xb5\xe5\xce\x31\x63\x79\x23\x7e\xdc\x7c\x76\xdb\x7c\x9b\xef\xf2\xed\xdb\x97\x7d\x80\x8d\xe7\xb3\x83\xd5\xde\xb0\x78\x52\x1f\x7a\xcc\x67\x8a\xc8\x4b\xed\xda\x3c\xc0\xba\xe0\x7e\x16\x16\x3d\x48\xa6\x38\x7e\x68\x7b\x8c\xfe\x13\xf7\x04\x31\xf1\x9f\xad\xdf\xf8\x52\x8a\x8a\x09\x43\x09\x95\xe5\x0a\x01\x6b\x44\xc8\x32\x13\x26\x24\x28\x75\x05\x28\xb4\x5a\xf7\xb4\xfc\x20\xb4\xea\x16\x09\x99\x38\x3c\x45\xcc\xb5\x0a\x36\x0a\x01\xf2\x9a\x4b\x41\x68\x7d\x51\x70\x12\xd5\x02\x07\x60\x8c\x14\x11\xcc\x0e\x80\xfc\x03\x2e\x94\xd0\x81\xc2\x65\x4a\x1e\x98\x00\x63\xb2\xd1\x7f\xa1\xb6\xe7\x4b\x9a\xa3\xb2\x20\xf4\x5f\xee\x90\xfc\xf3\x4e\x37\x1b\x42\xdc\xd1\xf7\xc9\xdd\xa5\x57\x10\xfa\x8e\xa0\x6a\xb0\x35\x7b\x3f\x81\x94\xb4\xb7\xa3\xe0\x36\xb5\x47\x4b\x05\x06\x2a\x81\x17\x52\x90\x2f\x1b\x42\xba\x10\xd7\x58\x5d\xf2\x34\xf2\x34\x8d\xbf\x03\x64\xbb\x8b\x31\x8c\xd5\xa8\x2b\x2d\x13\xcc\x5f\x88\x66\x00\x00\xb6\x33\x41\xde\xa4\x3e\x08\x15\x21\xad\x76\x38\x03\xe9\x11\x79\xa4\x79\xb2\x0a\xba\xe8\x26\x14\x72\x7b\x86\xe4\xe8\x6f\xdb\xa1\xea\x13\xd7\x1e\xc9\xac\xd1\xab\x96\x83\xc4\xf6\xc2\xb0\xb5\xdc\xb5\x5a\xd6\xa4\x20\x5f\x47\x0e\x06\x1d\x43\x4b\x55\x5a\x35\xe2\xe0\x6d\x94\x23\xa5\x65\x6e\x1e\x06\xe7\x4c\x98\x6c\xe2\x1c\x73\x8e\x6b\x83\x89\xfa\x89\xd1\x15\x75\xf7\x16\xf1\xee\xed\x06\x8d\x5f\xf2\xb0\xb2\x6e\x0d\xd3\xa7\xdd\x58\xad\x90\xab\x9a\x19\x6d\xf1\x3e\xd7\x82\xd0\xd1\x16\x4c\x2d\xa2\x71\x83\x38\x01\x59\x90\xfd\xfe\xeb\xab\x41\xa4\x3e\xa4\x31\x66\x82\xfc\x94\xc1\xb5\x91\xaa\x9a\x6c\x0c\xb4\xc0\xe6\xe3\xe8\xa7\xc4\x5e\x11\xb9\x2c\x03\xa1\x57\xb2\x4a\xa8\x8e\x21\xc3\xeb\xf2\xd6\x5a\x26\xe5\x3e\x64\x33\xf8\x64\x83\x4f\x16\x7c\x1e\x02\x06\x76\x99\xe3\x88\x42\x1d\xdc\x5a\xbd\xcf\x2e\xef\xac\xe4\x59\x8b\x0e\x87\x99\xd5\xfa\x28\x78\x7f\xe9\xd5\x0c\x9a\x46\xa8\x38\xc0\xf4\x4f\xe1\xc2\xcd\x57\xdf\x89\x32\x73\xe2\xef\xdb\xc5\xa9\x4d\xe6\xd6\xf2\xef\x9e\x3b\x64\xd3\x41\x2a\xc8\x6e\x0c\x50\x72\x96\xd6\x35\x5d\x0e\x3d\x2d\xce\xc9\xfe\x9a\x16\x4d\xd8\xb2\x0f\x9b\xa5\x20\xd4\x39\x99\x05\x44\x3c\xb6\x06\x84\x69\x3b\x24\x17\x7d\x37\xae\x95\x78\xb7\x4f\x71\xe3\xd7\x9b\xce\xbd\x1c\x52\x38\xe4\x8a\xdb\x55\x39\x5e\xd6\x25\x84\x96\x0e\x87\x5e\x5c\xec\x79\xb6\xd8\x4f\x3f\xe9\xee\xc9\x28\x3e\x70\xbd\x36\xd5\xb3\xea\xa6\x32\x0f\xe0\x44\xa0\xf4\x9c\x44\xa0\x9e\xd3\xb1\x35\xac\xf6\xa1\xcb\x99\xf5\x72\xed\xca\x78\x91\x56\xfb\x7d\xec\x85\x10\x97\xe1\xc5\xcc\xcf\xce\x1f\xe0\xc2\xad\x1e\xfe\x4d\x44\x9e\x56\xf0\x6b\x62\xce\x6d\x87\xfb\x17\xdb\x93\x8b\x61\x61\x2b\xbc\xf4\x76\xbb\x1f\xff\x2e\xbc\x61\xb4\x47\xe3\x91\x50\x30\x66\x7c\xb1\xf4\x01\xe3\x63\xe5\x0c\xd2\x27\xb1\x67\x9e\x38\xf9\x78\xce\xf5\x79\xf8\x7f\x00\x00\x00\xff\xff\x25\x9b\xa0\x98\x76\x0b\x00\x00")

func templatesCf_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_lbTf,
		"templates/cf_lb.tf",
	)
}

func templatesCf_lbTf() (*asset, error) {
	bytes, err := templatesCf_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_lb.tf", size: 2934, mode: os.FileMode(480), modTime: time.Unix(1512143621, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNetworkTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x5d\x4e\x86\x30\x10\x45\xdf\xbb\x8a\xc9\xc4\x57\xd8\x81\x2b\x31\xa6\x29\xed\x88\x8d\xd0\x92\xe9\x8f\x46\xd2\xbd\x9b\x12\x30\xa1\x42\xfc\xfa\x7c\x66\x7a\xcf\x1d\xa6\xe0\x13\x6b\x02\x54\xdf\x89\x89\x67\x99\x2d\xc7\xa4\x26\xe9\x28\x7e\x7a\xfe\x40\xc0\xc1\x87\x77\x84\x55\x00\x38\x35\x13\x34\xef\x19\xf0\x69\xcd\x8a\x7b\x72\x59\x5a\x53\xba\x8a\x77\xd9\xa1\x00\x50\xc6\x30\x85\x20\xc3\xa2\x34\xfd\xf2\x2f\xfb\xc0\xfe\x83\xd4\xd6\x70\xc1\x57\x01\x30\x79\xad\xa2\xf5\xee\x72\x3f\xd3\x68\xbd\x2b\x75\xef\x91\x5a\x8e\xec\xd3\x22\xb7\x58\x1b\x77\x48\x9c\x81\xbe\x46\xea\x2b\x55\x50\x14\x21\xfe\x4a\x87\x34\x38\x8a\xff\xba\xde\xc8\x86\x93\xec\xc2\xf4\x66\xbf\xda\x81\xb3\xec\x8d\xc3\xc3\x12\x00\xcd\x99\x2e\x3a\x68\x88\xa6\x84\x9f\x00\x00\x00\xff\xff\x2b\x4f\xd2\x88\xf9\x01\x00\x00")

func templatesNetworkTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetworkTf,
		"templates/network.tf",
	)
}

func templatesNetworkTf() (*asset, error) {
	bytes, err := templatesNetworkTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network.tf", size: 505, mode: os.FileMode(480), modTime: time.Unix(1512143589, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNetwork_security_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x96\xcd\x6e\x9b\x40\x10\x80\xef\x3c\xc5\x68\xd5\x53\x24\x5b\x29\x81\xc8\x97\x1c\x7a\xec\xbd\x77\xb4\xde\x1d\xf0\xaa\x78\x07\xcd\x2e\x4e\xdb\x88\x77\xaf\x16\x6a\x27\x38\x38\xc2\xa8\x55\x85\x05\xe7\x6f\xfe\x76\x3e\x8d\x60\x74\x54\xb3\x42\x10\xf2\x57\xcd\xc8\xfb\xcc\xa2\x7f\x26\xfe\x9e\x39\x54\x35\x1b\xff\x33\x2b\x98\xea\x4a\x80\xd8\x92\xdb\x09\x78\x89\x00\xac\xdc\x23\x9c\x7d\x4f\x20\x3e\xbd\x1c\x24\xaf\xd1\x1e\x32\xa3\x9b\x55\x8b\x47\x00\x25\x29\xe9\x0d\xd9\x41\x98\xb1\x30\x64\x9b\xc0\x1d\x3b\xe9\xea\x65\x6d\x8d\x96\x3b\x36\xd6\x07\xd6\x21\xff\x3a\x50\x8d\x88\x22\x00\x2f\x0b\xd7\x36\x07\x80\xf6\x60\x98\xec\x1e\xad\x7f\xd7\x56\xa8\xd4\x44\x4d\x14\x5d\x31\xb8\xca\xaf\x18\x5b\xe5\xf3\x1e\x9a\xeb\x12\x05\x08\xf7\xd1\xae\x2f\x0e\xef\xba\x95\x57\x6c\x28\x24\x1b\x8e\x89\xef\xef\x23\x00\x6d\x18\xd5\xf9\x13\xbd\xe6\xfd\x6a\xb7\x54\x5b\x1d\xb2\x49\xa5\xd0\xb9\x8b\x1d\x7c\x29\x4b\x7a\xee\xaa\x92\x27\x45\xe5\x05\xee\x9b\xaa\x02\xf5\xe7\x39\x2b\x62\x9f\xb1\xb4\x05\xf6\xa9\xbb\xc0\x68\x74\xde\xd8\x76\x81\xef\xc0\x27\x10\x71\xfc\x26\x91\xd4\x9a\xd1\xb9\xac\x62\xcc\xcd\x8f\x0f\x12\x9d\x83\x47\x66\x48\x81\xde\x03\x8f\x50\x01\x60\x58\xde\x01\xa1\x86\xc1\x5e\xb6\xab\x44\x09\x81\x2b\x59\xa0\xf5\x13\x7c\x79\x13\x3c\x42\x9b\xcf\xf3\xd6\xe6\x71\xf3\xb8\x59\xc4\xe9\x8b\xd3\xad\x93\x78\xaa\x3b\xa7\xf8\x11\xfa\xc4\xf3\xd6\x27\x4e\xd3\x34\x5d\xfc\x39\xf9\xa3\xad\x9b\x60\x4d\x88\x1a\xe1\xca\xc3\xff\x70\xe5\xee\x2f\x99\x92\x3e\x2c\x9a\x9c\x34\x51\x8c\x7a\x57\x6f\x27\xa8\x72\x8c\x1c\xa1\x4b\x32\xef\xd3\xb2\xd9\x24\xc9\xa2\xcc\xab\x32\xf9\x6a\xe7\x7d\xf5\x0f\xcf\xcb\xcc\xff\x64\x92\xe4\x06\x2f\x8c\xca\xa7\xca\x52\x52\x31\xe5\xbc\x74\x81\xb7\xff\xe3\x92\xdc\xbc\x2e\xbf\x03\x00\x00\xff\xff\x77\x27\x82\x78\x45\x11\x00\x00")

func templatesNetwork_security_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetwork_security_groupTf,
		"templates/network_security_group.tf",
	)
}

func templatesNetwork_security_groupTf() (*asset, error) {
	bytes, err := templatesNetwork_security_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network_security_group.tf", size: 4421, mode: os.FileMode(480), modTime: time.Unix(1512143591, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOutputTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\xdd\x6e\xa3\x30\x10\x85\xef\x79\x0a\x0b\xed\x35\x95\x22\x71\x13\x69\x9f\xc5\x32\x66\xda\xcc\xd6\xd8\xd6\xfc\xb0\xed\x56\x79\xf7\x15\x4b\x89\xe2\x26\x40\xb6\xb9\x0c\xdf\x7c\xe7\x68\xc6\x49\x25\xab\x98\xba\x4b\x7c\xb2\x11\xe4\x77\xa2\x57\x1b\xdd\x00\xb5\xf9\xa8\x8c\x19\x5d\x50\x30\x3f\x4d\xfd\xe3\xc3\xfd\x51\x02\x1a\xec\x88\x24\xea\xc2\x02\x37\xd3\x64\x33\x4d\x9c\xeb\xea\x5c\x55\x85\x90\xb5\x8b\x20\xdb\xbe\x99\xd9\xd4\x10\x70\x52\xf2\x60\x5f\x28\x69\xde\xd6\x95\xec\x76\x3b\x49\xe4\x5e\xc0\x3a\xef\x93\xc6\xbd\x9a\x25\xbc\x29\xee\xe1\xd9\x69\x10\xcb\xe0\x95\x50\xde\xe7\x2e\xab\xea\x65\xed\x25\xbe\x96\x00\x6f\x02\x14\x5d\xb0\xb8\x6e\xcc\xda\x05\xf4\x16\x3f\x25\x98\xad\xeb\x7b\x02\xe6\x52\xd5\x23\x81\x97\x44\xcb\xd7\x2f\xbe\x93\x48\xe6\xe3\xd3\xd3\x23\xde\xe3\xa1\x6d\xdb\xf6\x76\x15\xe3\xc0\x36\x13\x8e\x4e\xc0\xbe\xc2\xfb\x75\xc2\xf4\xfb\xd7\x5a\x42\xc1\x34\xcb\x60\x73\xf5\xa7\xcd\x30\x9c\xeb\xca\x18\x86\xc8\x28\x38\x4e\x0d\x85\x14\xee\x27\xce\x3d\xff\x3f\xf0\x32\x67\x53\x86\xc8\x7c\xba\xc9\x7c\x76\x81\x8b\xd0\x5f\x3a\xe4\x2e\xbd\x59\xa5\xf0\x8d\x7b\x1c\x0f\x87\x62\x69\xcb\x5b\xf0\xd8\xd3\x8d\x6e\x74\xd4\x5c\x03\xe5\x35\x31\x7e\x3e\x8c\xd5\xd9\x82\x28\x87\x59\x3b\xf6\x84\x59\x30\x45\x8b\xfd\xdd\xb5\x4d\x8a\x2f\xdc\xee\x49\x04\xa2\x8b\xb2\x65\xbc\x10\xbb\x2e\x1f\x10\xb6\x5d\x17\xe2\x51\x17\x83\x27\x90\x3d\xdf\x4c\xad\x39\xff\x06\x00\x00\xff\xff\xf8\xc9\xc8\x2d\x43\x05\x00\x00")

func templatesOutputTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesOutputTf,
		"templates/output.tf",
	)
}

func templatesOutputTf() (*asset, error) {
	bytes, err := templatesOutputTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/output.tf", size: 1347, mode: os.FileMode(480), modTime: time.Unix(1512143593, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResource_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x41\x8a\xc4\x20\x10\x45\xf7\x9e\xe2\x23\xb3\x9d\xdc\x60\xce\x22\x15\x53\x64\x84\x44\x43\xa9\x59\x4c\xf0\xee\x83\x42\xa7\x3b\x24\xdd\x0d\x5d\x4b\xf9\xf5\xeb\xf9\x84\x63\xc8\x62\x19\x9a\xfe\xb2\xb0\xcc\xe6\xf6\x62\x46\x09\x79\xd1\xd0\x7d\x88\xbf\x1a\x9b\x02\x3c\xcd\x8c\x3a\x3f\xd0\x5f\xdb\x4a\xd2\xb1\x5f\x8d\x1b\xca\x77\xcb\x28\x60\x0a\x96\x92\x0b\xfe\x9e\x10\x1e\x5d\xf0\x45\x2b\x05\x24\x1a\x63\x2b\x02\xd8\xaf\x4e\x82\x9f\xd9\xa7\x53\x5b\x2d\x2a\xaa\x28\x75\x86\x5b\x72\x3f\x39\x6b\xdc\x13\xae\xab\x79\xcf\xfa\x72\x6b\xe7\x07\x8e\x66\xcc\xf1\x6a\x5b\xb8\x76\xd8\xd5\x8b\x5d\x8d\xb7\x9a\xfd\x0f\x86\x86\x41\x38\x46\x43\xd3\xa3\xb7\x98\x28\x39\xfb\x89\xb0\xff\x00\x00\x00\xff\xff\x58\xec\xbb\xb4\xcd\x01\x00\x00")

func templatesResource_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesResource_groupTf,
		"templates/resource_group.tf",
	)
}

func templatesResource_groupTf() (*asset, error) {
	bytes, err := templatesResource_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/resource_group.tf", size: 461, mode: os.FileMode(480), modTime: time.Unix(1512143595, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesStorageTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x91\xcd\x6a\xc3\x30\x0c\xc7\xef\x7e\x0a\x61\x76\xee\x1b\xec\xbc\xfb\xfa\x00\x46\x71\x44\x66\x70\xa4\x20\x2b\x81\xad\xe4\xdd\x47\xbc\x36\x6d\x0d\xa5\x3d\x2e\xd7\xfc\xf4\xff\xb2\x52\x91\x59\x23\x81\xc7\x9f\x59\x49\xc7\x50\x4c\x14\x07\x0a\x18\xa3\xcc\x6c\x1e\x7c\x27\xe5\xcb\xc3\xc9\x01\x30\x8e\x04\xcd\xf7\x0e\xfe\xed\xb4\xa0\x1e\x4a\x1a\xa7\x4c\x81\x78\x09\xa9\x5f\xbd\x03\xb8\x88\x87\x41\x65\x9e\x42\xbd\xae\xf8\xc5\xeb\x1e\x38\x6c\x46\x87\x8d\x5a\xbd\x73\x00\x59\x22\x5a\x12\x6e\x1d\xaf\x96\x4a\x43\x12\xae\x5e\xe7\xb8\xc1\x12\x69\x0b\x1f\x0d\xb9\x47\xed\x6f\x39\xa5\x29\xa7\x3f\xfd\x60\xdf\x53\x0d\xf6\xf1\x79\xac\xc6\x86\x43\xa9\x7d\x01\x88\x97\xa4\xc2\x23\xb1\x5d\x6d\x6f\x2a\xae\x6e\x75\xee\xf1\x88\x51\xd8\x30\x31\xe9\xd3\x19\x6b\xd0\x8a\x3c\x18\x0e\x5e\x9e\x0e\xa0\x79\xc3\xb3\xc0\xdd\x7d\x83\x34\x02\x7b\xee\xed\x3f\x95\xb2\x4f\x34\x69\x5a\xd0\xc8\xbf\x5e\xbb\x18\x8d\x91\x72\x7e\x52\x7d\xc7\xfe\x75\xfd\x2e\x4b\xb7\x75\xff\x0d\x00\x00\xff\xff\x23\xe4\x2a\x58\x37\x03\x00\x00")

func templatesStorageTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesStorageTf,
		"templates/storage.tf",
	)
}

func templatesStorageTf() (*asset, error) {
	bytes, err := templatesStorageTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/storage.tf", size: 823, mode: os.FileMode(480), modTime: time.Unix(1512143597, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTlsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x41\x0a\x02\x31\x0c\x05\xd0\x7d\x4e\xf1\xc9\x09\x5c\x88\xe0\xa2\x0b\xaf\xa0\x07\x08\xad\x04\x5b\x6c\xa9\x24\xb1\x30\x0c\x73\xf7\x79\xa6\x3e\xff\xf6\x56\x70\x74\x97\x9f\xb5\x95\x43\xe5\xab\x1b\x83\xcb\xf4\x2a\x6b\x38\x63\x27\x20\xf7\xcf\xb4\x16\x75\x20\x81\x9f\xaf\x07\x13\x60\x9e\xa5\xb4\x70\x20\xe1\x7a\xb9\xdf\xe8\xa0\x33\x00\x00\xff\xff\x52\x4d\xac\xad\x51\x00\x00\x00")

func templatesTlsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesTlsTf,
		"templates/tls.tf",
	)
}

func templatesTlsTf() (*asset, error) {
	bytes, err := templatesTlsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/tls.tf", size: 81, mode: os.FileMode(480), modTime: time.Unix(1512143599, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xd1\x4a\xc5\x30\x0c\x40\xdf\xfb\x15\xa1\xf8\x3c\x9d\x88\x6f\x7e\xcb\xe8\xda\x28\xc1\x2e\x1d\x69\x57\xc1\xb1\x7f\x97\xad\x50\xb5\x94\x7b\xb7\xb7\x9c\x73\xa0\x49\x36\x42\x66\xf6\x08\x1a\x39\x4f\xe4\x34\xec\x87\x52\xbf\x53\xc1\x0f\x0a\xdc\x4e\x23\x2d\xab\xc7\xa9\x9f\xc4\x6d\x8e\x56\x68\x4d\x14\xb8\x83\x13\xb2\xe1\xd4\x01\xd6\x13\xde\x02\x11\xad\x60\x6a\x21\x63\xfa\x0a\xf2\x39\x59\x72\xa2\x61\x57\x00\x0e\xdf\xcd\xe6\x13\xbc\x81\x1e\x9f\x86\xeb\x7f\x1c\x5f\xb5\xfa\x97\x11\x27\x14\x36\xfe\x4e\xf7\xfc\x72\x75\xab\x84\x4c\x0e\x05\xb4\xf9\xde\x04\x65\x29\x45\xb3\xe9\x59\x3e\xec\xd9\xc8\xd0\x80\x43\x2b\x80\xba\x37\x94\xaf\xca\x15\x5c\x5a\xbd\x42\xab\x55\xf0\x57\x2b\x37\xe9\x68\x05\x1c\xe7\xeb\x7f\x02\x00\x00\xff\xff\xf5\xa0\x00\x27\xe3\x01\x00\x00")

func templatesVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVarsTf,
		"templates/vars.tf",
	)
}

func templatesVarsTf() (*asset, error) {
	bytes, err := templatesVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vars.tf", size: 483, mode: os.FileMode(480), modTime: time.Unix(1512143607, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/cf_lb.tf": templatesCf_lbTf,
	"templates/network.tf": templatesNetworkTf,
	"templates/network_security_group.tf": templatesNetwork_security_groupTf,
	"templates/output.tf": templatesOutputTf,
	"templates/resource_group.tf": templatesResource_groupTf,
	"templates/storage.tf": templatesStorageTf,
	"templates/tls.tf": templatesTlsTf,
	"templates/vars.tf": templatesVarsTf,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"cf_lb.tf": &bintree{templatesCf_lbTf, map[string]*bintree{}},
		"network.tf": &bintree{templatesNetworkTf, map[string]*bintree{}},
		"network_security_group.tf": &bintree{templatesNetwork_security_groupTf, map[string]*bintree{}},
		"output.tf": &bintree{templatesOutputTf, map[string]*bintree{}},
		"resource_group.tf": &bintree{templatesResource_groupTf, map[string]*bintree{}},
		"storage.tf": &bintree{templatesStorageTf, map[string]*bintree{}},
		"tls.tf": &bintree{templatesTlsTf, map[string]*bintree{}},
		"vars.tf": &bintree{templatesVarsTf, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

