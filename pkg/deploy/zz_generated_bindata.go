// Code generated for package deploy by go-bindata DO NOT EDIT. (@generated)
// sources:
// manifests/ccm.yaml
// manifests/coredns.yaml
// manifests/local-storage.yaml
// manifests/metrics-server/aggregated-metrics-reader.yaml
// manifests/metrics-server/auth-delegator.yaml
// manifests/metrics-server/auth-reader.yaml
// manifests/metrics-server/metrics-apiservice.yaml
// manifests/metrics-server/metrics-server-deployment.yaml
// manifests/metrics-server/metrics-server-service.yaml
// manifests/metrics-server/resource-reader.yaml
// manifests/rolebindings.yaml
// manifests/traefik.yaml
// +build !no_stage

package deploy

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _ccmYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\x3f\x6f\xe3\x30\x0c\xc5\x77\x7d\x0a\x21\x4b\x80\x03\x9c\xe0\xb6\x83\xc7\xeb\xd0\x3d\x40\xbb\xd3\x12\x9b\xa8\x91\x45\x81\xa4\x1c\xb4\x9f\xbe\x70\x9c\x16\x81\xdd\xfc\x71\x36\xc2\x20\x7f\xef\xd1\xd4\x83\x1c\x5e\x91\x25\x50\xaa\x2d\x37\xe0\x56\x50\x74\x47\x1c\x3e\x41\x03\xa5\xd5\xfe\x9f\xac\x02\xad\xbb\xbf\x66\x1f\x92\xaf\xed\x53\x2c\xa2\xc8\x1b\x8a\x68\x5a\x54\xf0\xa0\x50\x1b\x6b\x13\xb4\x58\x5b\x17\xa9\xf8\xca\x51\x52\xa6\x18\x91\xab\x16\x12\x6c\x91\x0d\x97\x88\x52\x9b\xca\x42\x0e\xcf\x4c\x25\x4b\x3f\x54\x59\x47\xc4\x3e\xa4\x73\x2d\x63\x2d\xa3\x50\x61\x87\xa7\xa6\x88\x20\x28\xc6\xda\x0e\xb9\x39\x7d\xdb\xa2\x0e\x00\x46\x50\x3c\x96\x25\xfb\xbe\x9c\x68\x2c\x16\x53\x24\x76\x98\x74\x84\x3c\x43\x65\x50\xb7\x9b\x0d\x4d\xe4\xc7\x36\x97\x7f\x96\x33\x66\xd7\xa2\xa0\x65\x84\x18\xbc\xdc\x05\x11\xe4\x2e\xb8\xb1\x87\x18\x44\x7f\xdf\xaa\x2f\x0f\xb3\xf1\xe0\x1c\x95\x4b\x7f\xef\x2e\x50\xee\x1f\x9c\x28\x26\xed\x28\x96\xf6\xd2\x6d\x7f\x8c\x3f\x66\x17\x93\xcf\x14\xae\x9d\x79\x22\x74\x98\xdc\xbd\xaa\xcc\xe3\x09\xf9\x1f\x92\x0f\x69\x3b\x2b\x28\x14\x71\x83\x6f\x7d\xe7\xf7\x8a\x57\x54\x8d\xb5\xd3\x58\xde\xd4\x90\xd2\xbc\xa3\xd3\x63\x1e\x87\xf1\x17\x41\xbe\x3d\x37\x34\x48\x06\x87\xb5\xdd\x97\x06\x2b\xf9\x10\xc5\xd6\x7c\x05\x00\x00\xff\xff\xff\xea\xf1\x4c\x44\x04\x00\x00")

func ccmYamlBytes() ([]byte, error) {
	return bindataRead(
		_ccmYaml,
		"ccm.yaml",
	)
}

func ccmYaml() (*asset, error) {
	bytes, err := ccmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ccm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _corednsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xdf\x6f\xdb\x36\x10\x7e\xf7\x5f\x41\x68\xe8\xdb\xe4\xd8\x0b\xda\x65\x7c\x6b\xed\xac\x0d\xd0\xb8\x46\x9c\xf4\x65\x18\x8a\x33\x75\xb6\xb8\x50\x3c\x8e\xa4\xdc\x78\x5d\xfe\xf7\x81\xfa\x65\xd1\x51\x8a\x34\xc8\xfc\x62\x91\xc7\xfb\x8e\xfa\x78\xfc\xee\x04\x46\x7e\x46\xeb\x24\x69\xce\x76\xd3\xd1\xad\xd4\x19\x67\x2b\xb4\x3b\x29\xf0\xad\x10\x54\x6a\x3f\x2a\xd0\x43\x06\x1e\xf8\x88\x31\x0d\x05\x72\x26\xc8\x62\xa6\x5d\x33\x76\x06\x04\x72\x76\x5b\xae\x31\x75\x7b\xe7\xb1\x18\xa5\x69\x3a\xea\x43\xdb\x35\x88\x31\x94\x3e\x27\x2b\xff\x01\x2f\x49\x8f\x6f\xcf\xdc\x58\xd2\xc9\x6e\xba\x46\x0f\x6d\xe4\x99\x2a\x9d\x47\x7b\x45\x0a\xa3\xb0\x0a\xd6\xa8\x5c\x78\x62\x55\x1c\xab\xd1\x63\xe5\xbf\x26\xf2\xce\x5b\x30\x46\xea\x6d\x1d\x28\xcd\x70\x03\xa5\xf2\xae\xdb\x6f\xbd\x2b\xde\x6e\xdb\x96\x0a\x1d\x1f\xa5\x0c\x8c\x7c\x6f\xa9\x34\x15\x72\xca\x92\x64\xc4\x98\x45\x47\xa5\x15\xd8\xcc\xa1\xce\x0c\x49\x5d\x81\xa5\xcc\xd5\xcc\xd4\x03\x43\x59\xfd\xd0\x91\x10\x86\x3b\xb4\xeb\xc6\x57\x49\xe7\xab\x87\xaf\xe0\x45\xfe\xb4\x78\x9a\xb2\x63\x98\x2d\xfa\x97\x20\xf4\x9d\xd4\x99\xd4\xdb\x88\x57\xd0\x9a\x7c\xe5\xde\x90\x3b\x84\x1b\xf1\x0d\xa5\xa7\xd2\x64\xe0\x91\xb3\xc4\xdb\x12\x93\x97\x3f\x1e\x52\x78\x85\x9b\x6a\x7f\x0d\x61\xdf\x79\xe1\x11\x63\x0f\x73\xe7\x11\x64\x57\xae\xff\x42\xe1\xab\xb3\x1f\x4c\xf5\x67\x27\x78\x77\x77\x66\xa4\x37\x72\x7b\x09\xe6\x39\xd7\xa6\x5d\x3e\x23\x8b\x1b\xa9\x90\xb3\x7f\x2b\x4e\xc7\xfc\xf5\x29\xfb\x56\x3d\x86\x1f\x5a\x4b\xd6\x75\xc3\x1c\x41\xf9\xbc\x1b\x5a\x84\x6c\xdf\x8d\x0e\xc7\xc1\x5e\x7d\x9b\x7d\xbc\x59\x5d\x9f\x5f\x7d\x99\x7f\xba\x7c\x7b\xb1\xb8\x7f\xc5\xa4\x4e\x21\xcb\xec\x18\xac\x01\x26\xcd\x9b\xfa\xe1\x10\x89\x55\x49\xce\xa4\x76\x28\x4a\x8b\xbd\xf9\x0d\x28\xe5\x73\x4b\xe5\x36\x1f\x46\xe9\xd6\xde\x1f\x36\x4a\xce\x3b\x76\x82\x5e\x9c\x34\x54\x9c\x2c\x28\xc3\x0f\xd5\x74\x3f\xa8\xf7\x8a\xbd\x99\xf4\x26\x2c\x2a\x82\x8c\x4d\x5f\xbb\xe1\x2d\x0c\x04\x33\x96\x0a\xf4\x39\x96\x8e\xf1\xdf\xa6\xaf\x4f\x3b\xc3\x86\xec\x57\xb0\x19\x1b\xd7\x3b\x09\xf7\x4f\xed\xc6\x82\xf4\xa6\x5b\x22\x40\xe4\xc8\x4e\x0f\x3b\x50\x44\x66\x14\x6f\xa6\x67\x83\x6c\x0d\x0a\xb4\xa8\xf9\xb9\x7f\x90\x1c\x60\x8c\x3b\xe9\x32\x64\x8e\x46\xd1\xbe\xc0\xe7\x29\xeb\xd1\x65\x3b\x73\x29\x18\xd3\x2c\xa9\x1d\x8f\xaf\x60\x0d\x9c\x84\x9c\x9a\x2f\x56\xc9\xc8\x19\x14\xc1\xfb\x27\x8b\x46\x49\x01\x8e\xb3\xe9\x88\xb1\x70\x4b\x3d\x6e\xf7\x35\xb0\xdf\x1b\xe4\xec\x8a\x94\x92\x7a\x7b\x53\xdd\xf7\x5a\x1f\xfa\x33\xbc\xe1\xa0\x80\xbb\x1b\x0d\x3b\x90\x0a\xd6\x21\x69\x2b\x38\x54\x28\x3c\xd9\x7a\x4d\x11\x04\xf0\x63\x6f\xe3\xc3\x5b\xf7\x58\x18\xd5\x01\xf7\xd9\xa9\x88\x8e\xfc\x1f\x7b\xf9\xf6\xf5\xea\x1c\x90\x64\xa5\xdf\xcf\x14\x38\xb7\xa8\x79\xa8\x79\x4c\x45\xad\x16\xa9\xb0\xd2\x4b\x01\x2a\x69\x5c\x5c\x24\x08\x8b\xa3\x43\xa9\xa8\x21\x85\xb6\xaf\x99\xe1\x97\xb2\x5b\xdc\x07\x96\x1b\xb8\xb7\x59\x46\xda\x7d\xd2\x6a\x9f\xf4\x32\x96\x4c\xf0\x24\xcb\x59\x72\x7e\x27\x9d\x77\xc9\x03\x80\xa0\xff\x69\x50\xc0\x23\xdd\x15\xa4\xbd\x25\x95\x1a\x05\x1a\x9f\x88\xc9\x18\x6e\x36\x28\x3c\x67\xc9\x82\x56\x22\xc7\xac\x54\xf8\xf4\x90\x05\x04\x86\x5e\x22\x56\x88\xb0\x8a\x12\x22\xfc\x42\x9d\x3a\x0a\x49\x8e\x33\x25\x75\x79\xd7\x2c\x0a\x6f\x0d\x52\xa3\xed\xa8\x4e\x1f\x5c\x94\xfa\x27\x0b\xd8\x22\x67\x16\xb4\xc8\xd1\xb6\xfa\x92\x36\xff\x7c\x3a\x3e\x1b\x4f\xe2\xc5\xcb\x52\xa9\x25\x29\x29\xf6\x9c\x5d\x6c\x16\xe4\x97\x16\x1d\x56\x45\xa0\xbd\xe7\xbd\xca\xdc\xdd\x76\x59\x48\x1f\xcd\x84\x44\x2d\xc8\xee\x39\x9b\xfe\x3a\xb9\x94\x91\x6a\xfd\x5d\xa2\x3b\x5e\x2d\x4c\xc9\xd9\x74\x32\x29\x06\x31\x22\x08\xb0\x5b\xc7\xd9\x1f\x2c\x49\x83\x3c\x25\x3f\xb3\x24\x12\xcf\xb6\x4c\x24\xec\xcf\xce\x65\x47\xaa\x2c\xf0\x32\x24\x6f\x94\x9e\x2d\x6b\xa1\x3a\xa5\xf5\xa2\x5e\xfc\x22\xac\x5f\x82\xcf\x79\x24\xcf\xd1\xbb\x40\x16\xd2\x99\xb3\x50\xf4\x0f\x2a\x4b\x36\x8e\xd3\x9d\xd8\x92\xac\xe7\xac\xa7\xbb\xad\xc4\xc5\xb8\xc6\x92\x27\x41\x8a\xb3\x9b\xf9\xf2\x47\x71\x52\x2f\xcc\x20\xd6\xf5\xec\x3b\x58\x51\x35\x68\xd1\x0a\xf4\x56\x8a\xe1\x9d\xf5\xd1\xaa\x42\x18\x24\x85\xb4\xc7\x3b\xdf\x3f\x5a\x50\x8a\xbe\x2e\xad\xdc\x49\x85\x5b\x3c\x77\x02\x54\x25\x13\x3c\x54\x2a\xd7\xa7\x5b\x80\x81\xb5\x54\xd2\x4b\x3c\x4a\x0e\xc8\xb2\x78\x22\x65\x8b\xf3\xeb\x2f\xef\x2e\x16\xf3\x2f\xab\xf3\xab\xcf\x17\xb3\xf3\xc8\x9c\x59\x32\xc7\x0e\xa0\xd4\xc0\xc1\x5d\x11\xf9\xdf\xa5\xc2\xa6\x25\x8a\x8f\x51\xc9\x1d\x6a\x74\x6e\x69\x69\x8d\x7d\xbc\xdc\x7b\xf3\x1e\x7d\x1c\xc2\xd4\x89\x72\xd4\x77\xb0\x26\x1d\x38\x3b\x9b\x9c\x4d\xa2\x69\x27\x72\x0c\x24\x7f\xb8\xbe\x5e\xf6\x0c\x52\x4b\x2f\x41\xcd\x51\xc1\x7e\x85\x82\x74\xe6\x78\x5c\xf7\x0d\x5a\x49\x59\x67\x9b\xf6\x6d\x5e\x16\x48\xa5\x3f\x18\x7b\x36\x57\x0a\x81\xce\x5d\xe7\x16\x5d\x4e\x2a\x8b\xad\x1b\x90\xaa\xb4\xd8\xb3\x9e\x46\xdd\x93\xfc\x61\x2a\xe2\x9e\xab\xc7\xc4\xf4\x6c\xfa\x6c\x26\xbe\x43\xc4\x2f\xff\x33\x0f\x99\x76\xad\x34\xce\xeb\x6e\xbd\x31\xd4\xca\xf1\x03\xca\x22\xda\x7e\x38\xe6\x6d\x58\xc4\x2b\x2a\x3c\x16\xee\x38\xa5\xab\x42\xd5\xca\x5d\x64\x6b\x8f\x60\xd0\xd8\x38\x76\x4d\xe6\xa0\xe7\xc1\xfa\x68\x53\xdf\x7c\x25\x0c\xf4\x6b\xbd\xd6\xe3\xd1\x86\xed\xc1\x47\xd6\xa1\x35\x0d\x35\xaf\xce\x94\x24\xa8\x52\x32\x60\x76\xc2\x82\x79\xf4\x63\xeb\x09\xfd\x5f\xdb\xe9\x34\x9d\x4d\x0f\xe9\xa9\x9d\x62\xdc\xcb\x0d\xc5\x6c\x62\x5c\x2c\x79\xff\x2b\x63\xb1\xba\x7f\x35\xea\xd5\x88\xf4\xa8\x02\x98\xbe\xb4\x1f\x17\x82\x74\x40\xe6\x1f\x71\xa8\xf5\x39\x1d\x50\x72\x13\x0b\x7e\xec\xf2\x5f\x00\x00\x00\xff\xff\xda\x91\xd1\x45\xfc\x10\x00\x00")

func corednsYamlBytes() ([]byte, error) {
	return bindataRead(
		_corednsYaml,
		"coredns.yaml",
	)
}

func corednsYaml() (*asset, error) {
	bytes, err := corednsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localStorageYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\xcd\x8e\xe3\x36\x0c\xbe\xfb\x29\x04\x01\x7b\x29\x56\xce\xa6\xbd\xb4\xba\xa5\xf9\x99\x1d\x20\x7f\x48\xa6\xed\x61\xb0\x08\x64\x99\x49\xb4\x23\x4b\x82\x24\xbb\x9b\x6e\xe7\xdd\x0b\xd9\x4e\xc6\x8e\x33\x33\x09\xda\xea\x12\x84\x22\xf9\x51\xe4\x47\xd2\xcc\x88\xdf\xc1\x3a\xa1\x15\x45\x45\x3f\x7a\x12\x2a\xa5\x68\x0d\xb6\x10\x1c\x06\x9c\xeb\x5c\xf9\x28\x03\xcf\x52\xe6\x19\x8d\x10\x52\x2c\x03\x8a\xa4\xe6\x4c\x12\xc3\xfc\x9e\x18\xab\x0b\x11\xec\xc1\x12\x57\xd9\x11\x56\x1b\x56\xea\xce\x30\x0e\x14\x3d\xe5\x09\x10\x77\x70\x1e\xb2\x88\x10\x12\x35\x91\x6d\xc2\x78\xcc\x72\xbf\xd7\x56\xfc\xc5\xbc\xd0\x2a\x7e\xfa\xd9\xc5\x42\xf7\x8a\x7e\x02\x9e\x1d\x03\x1b\xca\xdc\x79\xb0\x2b\x2d\xe1\xfa\xa8\x6c\xd0\xb6\xb9\x04\x47\x23\x82\x98\x11\x77\x56\xe7\xc6\x51\xf4\x88\xf1\x97\x08\x21\x0b\x4e\xe7\x96\x43\x29\x51\x3a\x05\x87\x3f\x22\x6c\x42\x6c\xce\x83\xf2\x85\x96\x79\x06\x5c\x32\x91\xb9\xd2\xa0\x00\x9b\x94\xca\x3b\xf0\x41\x55\x0a\x57\xfe\xfe\xc9\x3c\xdf\xe3\x2f\xef\x83\x80\x4a\x8d\x16\xca\x5f\x04\xaa\x84\x3a\x3d\xc3\xfa\xe1\x2a\xc7\x05\x04\xaf\x2d\x43\x6e\x81\x79\x28\x9d\x5e\x8e\xcf\x79\x6d\xd9\x0e\xea\x8c\x77\x9d\xd6\xf7\x5c\x32\xe7\xe0\xca\x0c\xfc\xfb\xfa\xfe\x2a\x54\x2a\xd4\xee\xfa\x32\x27\x42\xa5\x51\xa8\xf5\x0a\xb6\x41\xf9\xf8\xc6\x37\xd0\x23\x84\xba\xbc\xba\x86\x4d\x2e\x4f\xbe\x02\xf7\x25\xa1\x2e\xb6\xcc\xff\xd5\x28\xcc\x18\xd7\x3b\xf5\xe9\x08\x8c\xd4\x87\x0c\x6e\xe8\xd1\xd7\xa1\x9c\x01\x4e\xcb\xda\x1b\x29\x38\x73\x14\xf5\x23\x84\x1c\x48\xe0\x5e\xdb\x70\x83\x50\x16\xea\x3b\x65\x09\x48\x57\x09\x42\x9a\xcd\x1b\x58\x1e\x32\x23\x99\x87\xda\xbc\x11\x64\x38\xb2\xe5\xe9\x3d\x5f\x08\x1d\x43\x0c\xc7\x58\xa1\xad\xf0\x87\x61\xa0\xe5\xbc\x7c\x31\xae\x5e\x42\x42\x0f\x13\x6e\x85\x17\x9c\x49\x5c\xeb\xbb\x56\x81\xe6\xb7\x55\x27\x1c\xaf\x25\xd8\x92\x3d\x8d\x88\x11\x22\xe8\x09\x0e\x14\xe1\x61\x8d\x37\x48\x53\xad\xdc\x42\xc9\x03\x6e\x68\x21\xa4\x4d\xb0\xd6\x96\x22\x3c\xfe\x26\x9c\x77\xf8\x82\x93\x32\xf2\xc0\xb0\x38\x54\xc6\x2a\xf0\x50\x76\x09\xd7\xca\x5b\x2d\x89\x91\x4c\xc1\x0d\x7e\x11\x82\xed\x16\xb8\xa7\x08\xcf\xf5\x9a\xef\x21\xcd\x25\xdc\x02\x9c\xb1\xd0\x17\xff\x15\x62\x78\x06\x13\x0a\xec\x29\x83\xe4\x3d\xb2\x56\x47\x64\x6c\x07\x14\x59\xa6\xf8\x1e\x6c\xef\xb2\x36\x2d\x3e\xc5\x9f\xe2\xfe\x2f\x6d\xab\x65\x2e\xe5\x52\x4b\xc1\x0f\x14\xdd\x6f\xe7\xda\x2f\x2d\x38\x38\x55\x35\x04\x95\x65\x4c\xa5\x2f\x35\x25\xef\x45\x43\x90\xf3\xcc\xfa\xc6\x7f\x42\xb8\x56\x5b\xb1\x6b\x88\x7a\xe0\x79\xaf\x92\xd6\x3f\xf1\x57\xa7\xd5\x49\xa3\x1a\xf6\xb3\x40\x30\xd7\xc4\xae\xf2\x51\x59\x90\x4a\xa9\x91\xde\x2c\xe8\x2f\x99\xdf\xd3\x16\xc0\x49\x03\x54\xd1\x75\xb6\x5c\x8c\x36\xf3\xc1\x6c\xbc\x5e\x0e\x86\xe3\x86\xb3\x82\xc9\x1c\x26\x56\x67\xb4\x55\xc0\xad\x00\x99\xd6\x43\xb4\x23\xaf\xb0\x8f\x8d\x1c\x9f\x66\x49\x07\xf4\xf3\x78\xba\x1c\xaf\x36\xf7\xb3\xc1\x5d\x07\xb3\x51\x49\x91\x58\x66\x0f\x24\xc9\xdd\x21\xd1\xdf\x68\x3f\xfe\xe9\xc7\xb8\x1f\x35\x53\x74\x43\x76\x2a\xf9\x8c\x99\x76\xe8\x1d\x8a\xd5\xc5\x3a\x1f\xae\xed\x45\xf8\x32\x66\xd7\x95\xbc\x9c\x34\x6f\x0e\xda\xb0\x75\x94\xd2\xbe\x39\x25\x9a\xdb\xf3\xac\xb9\x84\x23\x29\x6c\x59\x2e\x3d\x29\xaf\x29\xc2\xde\xe6\x80\xa3\x26\xa9\x8f\xa9\x0a\x06\x0d\xa4\xea\xed\xf5\x92\x9c\xe9\x14\x28\xfa\x83\x09\x3f\xd1\x76\x22\xac\xf3\x43\xad\x5c\x9e\x81\x8d\x6c\xf5\xe1\x72\xec\x80\x11\x48\xf0\x50\xbe\xbc\xde\x7c\xc7\x94\x45\x67\x5f\x82\x6f\x2e\x94\x13\xdb\x5f\xd9\x25\x47\xc3\x06\xf1\x29\xfa\x9b\x9c\xaa\xf2\xfd\x8c\x5a\xa8\x9a\x40\x81\x5c\x33\x66\x30\x7d\xec\xdc\x77\x2d\x5a\x96\x98\xe2\xd1\x78\x32\xf8\x6d\xfa\xb0\x59\x0e\x1e\x3e\x6f\x26\x8b\xd5\x66\xbe\x98\x6f\xa6\xf7\xeb\x87\xf1\x68\x33\x5f\x8c\xc6\x6b\xfc\xf1\x75\x1f\xe1\x55\x0e\xd3\x47\xfc\xe1\xfb\xd1\xcf\x74\x31\x1c\x4c\x37\xeb\x87\xc5\x6a\x70\x37\x2e\xbd\x3e\x7f\x28\x3f\x7e\xda\xe7\xb9\x23\x79\xd1\x79\x8e\xfe\x09\x00\x00\xff\xff\x42\xe4\x9d\x62\x5f\x0b\x00\x00")

func localStorageYamlBytes() ([]byte, error) {
	return bindataRead(
		_localStorageYaml,
		"local-storage.yaml",
	)
}

func localStorageYaml() (*asset, error) {
	bytes, err := localStorageYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "local-storage.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerAggregatedMetricsReaderYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xcf\x31\x6b\xf4\x30\x0c\xc6\xf1\xdd\x9f\x42\x78\x7e\x93\x97\x6e\xc5\x6b\x87\xee\x1d\xba\x94\x1b\x94\xf8\x21\x27\xce\xb1\x83\x24\xe7\x68\x3f\x7d\xb9\x70\xdc\x58\x68\x27\x0d\x7f\x7e\x0f\xe8\x22\x35\x27\x7a\x29\xdd\x1c\xfa\xd6\x0a\x02\x6f\xf2\x0e\x35\x69\x35\x91\x4e\x3c\x8f\xdc\xfd\xdc\x54\xbe\xd8\xa5\xd5\xf1\xf2\x6c\xa3\xb4\xff\xfb\x53\x58\xe1\x9c\xd9\x39\x05\xa2\xca\x2b\x12\xd9\xa7\x39\xd6\xc4\xcb\xa2\x58\xd8\x91\x87\x15\xae\x32\xdb\xa0\xe0\x0c\x0d\x44\x85\x27\x14\xbb\x11\xfa\x61\xfd\xb1\x30\x78\x1b\x76\xc1\x35\x51\x74\xed\x88\xbf\x71\xc8\xe2\x7f\x71\x9c\x57\xa9\x0f\xa8\xbd\xc0\x52\x18\x88\x37\x79\xd5\xd6\x37\x4b\xf4\x11\xef\x7f\xdd\x7d\x3c\x05\x22\x85\xb5\xae\x33\x8e\xbe\xb5\x6c\xf1\x1f\xc5\xda\x32\xec\xc8\x3b\x74\x3a\xd2\x02\xbf\x95\x22\x76\xdc\x2b\xfb\x7c\x8e\xa7\xf0\x1d\x00\x00\xff\xff\xe5\x1d\x7a\x17\x89\x01\x00\x00")

func metricsServerAggregatedMetricsReaderYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerAggregatedMetricsReaderYaml,
		"metrics-server/aggregated-metrics-reader.yaml",
	)
}

func metricsServerAggregatedMetricsReaderYaml() (*asset, error) {
	bytes, err := metricsServerAggregatedMetricsReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/aggregated-metrics-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerAuthDelegatorYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xbb\x4e\xc4\x40\x0c\x45\xfb\xf9\x8a\xf9\x81\x09\xda\x0e\xb9\x03\x0a\xfa\x45\xa2\x77\x26\x97\xc5\x24\x19\x47\xb6\x27\x12\x7c\x3d\x5a\x11\xd1\xf0\x68\xaf\x74\xcf\x39\xa5\x94\xc4\x9b\x3c\xc3\x5c\xb4\x51\xb6\x91\xeb\xc0\x3d\x5e\xd5\xe4\x83\x43\xb4\x0d\xf3\xad\x0f\xa2\x37\xfb\x69\x44\xf0\x29\xcd\xd2\x26\xca\x0f\x4b\xf7\x80\x9d\x75\xc1\xbd\xb4\x49\xda\x25\xad\x08\x9e\x38\x98\x52\xce\x8d\x57\x50\x5e\x11\x26\xd5\x8b\xc3\x76\x18\xf9\xbb\x07\x56\xba\xd2\xcb\x84\x05\x17\x0e\xb5\x64\xba\xe0\x8c\x97\xeb\x8b\x37\x79\x34\xed\xdb\x3f\x19\x29\xe7\x1f\x01\xdf\xbe\xdf\x05\xde\xc7\x37\xd4\x70\x4a\xe5\xf8\x3e\xc1\x76\xa9\xb8\xab\x55\x7b\x8b\x3f\x72\x8f\xd9\x37\xae\xa0\x3c\xf7\x11\xe5\x8b\x9f\x3e\x03\x00\x00\xff\xff\x69\xfc\x98\x93\x34\x01\x00\x00")

func metricsServerAuthDelegatorYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerAuthDelegatorYaml,
		"metrics-server/auth-delegator.yaml",
	)
}

func metricsServerAuthDelegatorYaml() (*asset, error) {
	bytes, err := metricsServerAuthDelegatorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/auth-delegator.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerAuthReaderYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\x3d\x4e\xc4\x40\x0c\x46\xfb\x39\xc5\x5c\xc0\x41\xdb\xa1\xe9\xa0\xa1\x5f\x24\x7a\x67\xf2\x01\x26\x1b\x4f\x64\x7b\x22\xe0\xf4\x28\x68\xf9\x69\x76\x7b\xfb\x7d\xef\x11\x51\xe2\x55\x9e\x60\x2e\x4d\x4b\xb6\x91\xeb\xc0\x3d\x5e\x9b\xc9\x27\x87\x34\x1d\xe6\x5b\x1f\xa4\xdd\x6c\x87\x11\xc1\x87\x34\x8b\x4e\x25\x1f\xdb\x09\xf7\xa2\x93\xe8\x4b\x5a\x10\x3c\x71\x70\x49\x39\x2b\x2f\x28\x79\x41\x98\x54\x27\x87\x6d\x30\xda\x79\x64\xe0\x09\x76\x3e\xf1\x95\x2b\x4a\x9e\xfb\x08\xf2\x0f\x0f\x2c\xc9\xda\x09\x47\x3c\xef\x10\x5e\xe5\xc1\x5a\x5f\xaf\xe8\xa4\x9c\xff\x44\x7e\x77\xf1\x1e\xd0\x3d\x84\x78\x95\x7f\xe3\xd0\x90\xfa\xfd\xfe\xa3\xe1\x7d\x7c\x43\x0d\x2f\x89\xce\xa0\x47\xd8\x26\x15\x77\xb5\xb6\xae\x71\x21\xe5\xb2\xfe\x57\x00\x00\x00\xff\xff\xc7\x9e\x8d\xd1\x49\x01\x00\x00")

func metricsServerAuthReaderYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerAuthReaderYaml,
		"metrics-server/auth-reader.yaml",
	)
}

func metricsServerAuthReaderYaml() (*asset, error) {
	bytes, err := metricsServerAuthReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/auth-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerMetricsApiserviceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8f\x4d\x6a\xc4\x30\x0c\x85\xf7\x3e\x85\x2e\xe0\x74\xb2\x2b\xde\x75\x59\x68\x61\x20\x65\xf6\x1a\x8f\x3a\x08\xe3\x1f\x24\x39\x90\xdb\x97\x30\x71\xba\x13\x7a\xef\xfb\x90\xbc\xf7\x0e\x1b\xdf\x48\x94\x6b\x09\x80\x8d\x85\x9e\xac\x26\x68\x5c\xcb\x94\xde\x75\xe2\xfa\xb6\xce\x77\x32\x9c\x5d\xe2\xf2\x08\xf0\x71\xfd\x5c\x48\x56\x8e\xe4\x32\x19\x3e\xd0\x30\x38\x80\x82\x99\x02\x1c\xd5\x29\x93\x09\x47\x3d\x0c\x4e\x1b\xc5\xbd\xa4\x2f\x70\x1f\x07\x71\x34\xfd\x1e\x91\x9c\x81\x36\x8c\x14\x20\xf5\x3b\x79\xdd\xd4\x28\x3b\x80\xa7\xd4\xde\x4e\x64\xc8\x01\xd6\xf1\xc0\xb8\x14\x80\x8b\x52\xec\x42\x4b\xe2\xf6\xf3\xb5\xdc\x48\xf8\x77\x0b\x60\xd2\x69\x88\xae\xc2\x55\xd8\xb6\x6f\x2e\x9c\x7b\x0e\x30\x5f\x2e\xff\xb2\x91\xbe\xd6\x7f\x01\x00\x00\xff\xff\x25\x03\x92\xf5\x2a\x01\x00\x00")

func metricsServerMetricsApiserviceYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerMetricsApiserviceYaml,
		"metrics-server/metrics-apiservice.yaml",
	)
}

func metricsServerMetricsApiserviceYaml() (*asset, error) {
	bytes, err := metricsServerMetricsApiserviceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/metrics-apiservice.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerMetricsServerDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x52\xcd\x6e\x13\x41\x0c\xbe\xef\x53\x58\xcb\x79\x12\x10\x12\x42\x73\xab\x5a\x6e\x10\x90\x2a\x71\x77\x67\x1d\x32\xca\xcc\x78\x64\x3b\x81\x15\xe2\xdd\xd1\x74\xb7\x4b\xd2\x36\xa8\x12\xf8\x68\x7d\x7f\xfe\x71\xce\x75\x58\xe3\x57\x12\x8d\x5c\x3c\x1c\xdf\x74\xfb\x58\x06\x0f\xb7\x24\xc7\x18\xe8\x2a\x04\x3e\x14\xeb\x32\x19\x0e\x68\xe8\x3b\x80\x82\x99\x3c\x64\x32\x89\x41\x9d\x92\x1c\x49\xe6\xb6\x56\x0c\xe4\x61\x7f\xb8\x23\xa7\xa3\x1a\xe5\xee\xb1\x03\xd6\xaa\xeb\xc5\xe6\x86\x6a\xe2\x31\xd3\x3f\x59\x00\x24\xbc\xa3\xa4\x8d\x09\xb0\x7f\xaf\x0e\x6b\x7d\x42\xd7\x4a\xa1\x21\x94\x12\x05\x63\x99\xd0\x19\x2d\xec\x3e\x9e\xd0\x2f\x0b\x00\x18\xe5\x9a\xd0\x68\xa6\x9e\x04\x6e\x75\x21\x74\xab\x74\x66\xf0\x37\x0b\x80\x87\x9c\xad\xaa\x44\x96\x68\xe3\x75\x42\xd5\xcd\xbd\x7e\x3f\x0d\xed\x0a\x0f\xe4\x82\x44\x8b\x01\x53\x3f\xe3\xf5\xec\x6a\x9b\xcb\x81\x8c\x13\x09\x5a\xe4\x72\x92\xca\xc1\x9e\x46\x0f\xfd\xf5\xac\x7a\x35\x0c\x5c\xf4\x73\x49\x63\xbf\x60\x00\xb8\x36\x26\x8b\x87\xfe\xc3\x8f\xa8\xa6\xfd\x13\x81\xfb\x6c\xc2\x89\x56\xed\x4c\x52\xc8\x48\x57\x91\xd7\x81\x8b\x09\x27\x57\x13\x16\x7a\xa1\x26\x00\x6d\xb7\x14\xcc\x43\xbf\xe1\xdb\xb0\xa3\xe1\x90\xe8\xe5\x96\x19\xd5\x48\xfe\x87\xd7\x91\xd3\x21\xd3\xb2\xae\x57\x90\xdb\x8e\x21\x16\xb0\x5c\x41\x19\xbe\x13\x04\x2c\xa0\xb8\xa5\x34\xc2\x41\x09\xb6\xc2\xd9\x69\x90\xf6\x63\x10\x33\x7e\x23\x05\x2c\xc3\x9a\x05\x84\x70\x70\x5c\xd2\x08\x6d\x29\x18\x0b\x89\x76\x0f\x23\x4d\x9f\x64\xb9\xba\x21\xca\x92\x8e\x72\xb5\xf1\x26\x8a\x87\x9f\xbf\xe6\xe6\x1f\xae\x7f\x44\x7e\xf6\xea\x30\x85\xf0\x20\x58\xc2\x8e\x64\x7d\x8e\xf2\xc7\xd7\xab\xb7\xab\x77\x0b\x78\x9a\xf8\x53\x1b\xf3\xec\x4b\x9e\x8f\x07\xd3\x42\xbe\xa0\xed\x3c\xac\x2d\xd7\xae\xfb\x1d\x00\x00\xff\xff\x2c\xcf\xfd\xe4\x5e\x04\x00\x00")

func metricsServerMetricsServerDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerMetricsServerDeploymentYaml,
		"metrics-server/metrics-server-deployment.yaml",
	)
}

func metricsServerMetricsServerDeploymentYaml() (*asset, error) {
	bytes, err := metricsServerMetricsServerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/metrics-server-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerMetricsServerServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xcd\x4a\x04\x31\x10\x84\xef\x79\x8a\x26\xf7\x28\xe2\x1e\x24\x57\xcf\xc2\x82\xe2\xbd\x37\x5b\x48\x98\xfc\xd1\xdd\x33\xe0\xdb\xcb\xc4\x41\x10\xe6\xd6\x54\x57\x7d\x55\x21\x04\xc7\x23\x7f\x42\x34\xf7\x16\x69\x7b\x72\x4b\x6e\xf7\x48\xef\x90\x2d\x27\xb8\x0a\xe3\x3b\x1b\x47\x47\xd4\xb8\x22\x52\x85\x49\x4e\x1a\x14\xb2\x41\x0e\x59\x07\x27\x44\x5a\xd6\x1b\x82\x7e\xab\xa1\x3a\xa2\xc2\x37\x14\xdd\x93\x34\x3f\xd2\x60\xd0\x87\xdc\x1f\x7f\x49\xfe\xed\x1f\xca\x9f\x18\x53\x59\xd5\x20\xd3\x91\xf7\x06\x6f\xb2\xc2\x3b\x1d\x48\x3b\x58\x51\x90\xac\xcb\x51\xf2\xa2\x81\xc7\x38\xd9\x38\xba\xd8\x5c\x12\xe6\x19\xe9\x72\x79\x9e\x91\x21\xdd\x7a\xea\x25\xd2\xc7\xeb\x75\x2a\xc6\xf2\x05\xbb\xfe\xb9\x7e\x02\x00\x00\xff\xff\x92\x19\xf9\x3e\x23\x01\x00\x00")

func metricsServerMetricsServerServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerMetricsServerServiceYaml,
		"metrics-server/metrics-server-service.yaml",
	)
}

func metricsServerMetricsServerServiceYaml() (*asset, error) {
	bytes, err := metricsServerMetricsServerServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/metrics-server-service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerResourceReaderYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\x3f\x4f\xc4\x30\x0c\xc5\xf7\x7c\x0a\xeb\xf6\xf4\xc4\x86\xb2\x01\x03\xfb\x21\xb1\xbb\xa9\xb9\x33\x6d\xe3\xca\x76\x8a\xe0\xd3\xa3\x6b\xcb\x1f\x81\x74\x42\x62\xca\x4b\xe2\x9f\x9f\xde\x8b\x31\x06\x9c\xf8\x91\xd4\x58\x4a\x02\x6d\x31\x37\x58\xfd\x24\xca\x6f\xe8\x2c\xa5\xe9\xaf\xad\x61\xd9\xcf\x57\xa1\xe7\xd2\x25\xb8\x1b\xaa\x39\xe9\x41\x06\x0a\x23\x39\x76\xe8\x98\x02\x40\xc1\x91\x12\xd8\xab\x39\x8d\x69\x24\x57\xce\x16\x8d\x74\x26\x0d\x5a\x07\xb2\x14\x22\xe0\xc4\xf7\x2a\x75\xb2\x33\x11\x61\xb7\x0b\x00\x4a\x26\x55\x33\x6d\x6f\x93\x74\xb6\x88\x22\x1d\x7d\x53\x7b\x73\xf4\xed\x8e\x23\xd9\x84\x79\xf9\x9e\x49\xdb\x0d\x3d\x92\x2f\xe7\xc0\xb6\x8a\x17\xf4\x7c\x0a\xff\x0b\x79\xcb\xa5\xe3\x72\xfc\x7b\x56\x19\xe8\x40\x4f\xe7\xb1\x8f\xb4\x17\x2c\x03\xc0\xef\x5a\x2f\x1b\x58\x6d\x9f\x29\xfb\xd2\xe7\xca\x3e\x90\xce\x9c\xe9\x26\x67\xa9\xc5\x3f\xf1\x1f\x1c\x7c\xf5\x96\xa0\xaf\x2d\xc5\x75\x7f\x78\x0f\x00\x00\xff\xff\x39\x82\xcc\x46\x05\x02\x00\x00")

func metricsServerResourceReaderYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerResourceReaderYaml,
		"metrics-server/resource-reader.yaml",
	)
}

func metricsServerResourceReaderYaml() (*asset, error) {
	bytes, err := metricsServerResourceReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/resource-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rolebindingsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x31\x6f\xe3\x30\x0c\x85\x77\xfd\x0a\x21\xbb\x72\x38\xdc\x72\xf0\xd8\x0e\xdd\x03\xb4\x3b\x6d\xb3\x09\x6b\x59\x14\x48\x2a\x41\xfb\xeb\x0b\xa7\x6e\x82\xa4\x76\x90\xb4\xdd\x24\x41\x7c\x1f\x1f\xf9\x20\xd3\x13\x8a\x12\xa7\xca\x4b\x0d\xcd\x12\x8a\x6d\x58\xe8\x0d\x8c\x38\x2d\xbb\xff\xba\x24\xfe\xb3\xfd\xeb\x3a\x4a\x6d\xe5\xef\x63\x51\x43\x59\x71\xc4\x3b\x4a\x2d\xa5\xb5\xeb\xd1\xa0\x05\x83\xca\x79\x9f\xa0\xc7\xca\x77\xa5\xc6\x00\x99\x14\x65\x8b\x12\x86\x6b\x44\x0b\xd0\xf6\x94\x9c\x70\xc4\x15\x3e\x0f\xbf\x21\xd3\x83\x70\xc9\x17\xc8\xce\xfb\x2f\xe0\x03\x47\x5f\xd5\xb0\xaf\x0e\xfa\x99\x46\x86\x96\xfa\x05\x1b\xd3\xca\x85\x9b\x20\x8f\x8a\x32\xe3\xc2\xb9\x10\x82\xfb\xfe\xb4\x26\xc6\xf4\xd9\xfe\x3f\x0d\x0d\x27\x13\x8e\x11\xc5\x49\x89\x78\xd2\xb8\x0e\x15\xc1\x2f\x16\xce\x7b\x41\xe5\x22\x0d\x8e\x6f\x89\x5b\x54\xe7\xfd\x16\xa5\x1e\x9f\xd6\x68\x57\xd6\x42\x8f\x9a\xa1\x39\x17\x88\xa4\xb6\x3f\xec\xc0\x9a\xcd\x84\x56\x42\xdb\xb1\x74\x94\xd6\xa3\xdf\x29\xf1\x8f\x3f\x99\x23\x35\x74\x33\x61\x42\x10\x53\x9b\x99\x92\xe9\xfe\x96\xb9\x9d\xd3\x1c\xfc\x1f\xb5\x7f\xb8\xb4\xf9\x88\xcf\xec\xee\xf7\xb3\x7d\x0a\x38\x06\x7b\xf0\x78\x1d\xe3\x2c\xdc\x97\x01\xef\x01\x00\x00\xff\xff\x46\xd3\x6d\x9d\x0f\x04\x00\x00")

func rolebindingsYamlBytes() ([]byte, error) {
	return bindataRead(
		_rolebindingsYaml,
		"rolebindings.yaml",
	)
}

func rolebindingsYaml() (*asset, error) {
	bytes, err := rolebindingsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rolebindings.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _traefikYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\xcf\x8a\xdb\x4c\x10\xc4\xef\x7a\x8a\x46\xb0\x47\x49\x9f\x6f\x1f\x73\xdb\x38\x86\x84\x80\xb3\xc4\x49\xae\xa1\x35\x2a\x5b\x83\x47\x33\xa2\xbb\x65\xa2\xfc\x79\xf7\x20\xad\xe3\x0d\x21\x81\x25\xba\xa9\xe7\x57\xd5\x35\xc5\xf0\x18\x3e\x42\x34\xe4\xe4\xa8\x47\x1c\x6a\xcf\x66\x11\x75\xc8\xcd\x65\x53\x9c\x43\xea\x1c\xbd\x42\x1c\xb6\x3d\x8b\x15\x03\x8c\x3b\x36\x76\x05\x51\xe2\x01\x8e\x4c\x18\xc7\x70\xbe\xfe\xeb\xc8\x1e\x8e\xce\x53\x8b\x4a\x67\x35\x0c\x85\x8e\xf0\x0b\xee\x17\x03\x47\xbd\xd9\xa8\xae\x69\xee\xbe\xbe\xf9\xf0\x62\xf7\x6e\xbf\x7b\xbf\x3b\x7c\xba\x7f\x78\xfd\xfd\xae\x51\x63\x0b\xbe\x59\x41\x6d\xae\xc6\xd5\xa6\xfe\x7f\x53\xff\x57\xdb\xe9\x4b\x41\x74\xe1\x38\x41\xb7\x39\x19\x92\x39\xfa\x56\x15\x44\x44\xd2\xf2\xba\x62\xf9\x90\xb8\x8d\xe8\x96\x60\x13\xd6\x99\x6a\xfc\xfb\xe1\x00\x93\xe0\xf5\x27\x30\x4a\x1e\x60\x3d\xa6\xdb\xe4\x4f\xa2\xe5\x7a\x92\x60\xb8\x51\x21\x9d\x04\xaa\xbb\xd4\x8d\x39\x24\x7b\x12\x4f\x8a\x97\x38\xf2\x14\xed\x61\x6a\x63\xd0\x1e\xdd\x01\x72\x09\x1e\xbf\xf8\x8d\x12\xb2\x04\x9b\xb7\x91\x55\xf7\x6b\xad\xe5\x63\x7b\x95\x8f\x93\x1a\xa4\xf2\x12\x2c\x78\x8e\xe5\x2a\x08\x03\x9f\x16\x48\x38\xf9\x1e\xd2\xc4\xd0\x0a\xcb\x5c\x5d\x3b\x7b\x84\x2c\x47\x08\x5b\xc8\xe9\x16\xb3\xa2\x33\x66\x47\xe5\xf6\xea\x76\xdf\x75\x39\xe9\xdb\x14\xe7\xf2\x96\x38\x8f\x8b\x2a\x8b\xa3\x72\xf7\x39\xa8\x69\xf9\x9b\x38\xe5\x0e\x95\xe4\x88\xfa\xa9\x87\xe5\xb9\xf8\x9c\x4c\x72\xac\xc6\xc8\x09\xcf\xf0\x23\xc2\xf1\x08\x6f\x8e\xca\x7d\x3e\xf8\x1e\xdd\x14\xf1\xdc\x65\x03\x2f\xb5\xfc\xfb\x96\x1f\x01\x00\x00\xff\xff\x8e\x38\x58\x4b\xf7\x02\x00\x00")

func traefikYamlBytes() ([]byte, error) {
	return bindataRead(
		_traefikYaml,
		"traefik.yaml",
	)
}

func traefikYaml() (*asset, error) {
	bytes, err := traefikYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "traefik.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"ccm.yaml":           ccmYaml,
	"coredns.yaml":       corednsYaml,
	"local-storage.yaml": localStorageYaml,
	"metrics-server/aggregated-metrics-reader.yaml": metricsServerAggregatedMetricsReaderYaml,
	"metrics-server/auth-delegator.yaml":            metricsServerAuthDelegatorYaml,
	"metrics-server/auth-reader.yaml":               metricsServerAuthReaderYaml,
	"metrics-server/metrics-apiservice.yaml":        metricsServerMetricsApiserviceYaml,
	"metrics-server/metrics-server-deployment.yaml": metricsServerMetricsServerDeploymentYaml,
	"metrics-server/metrics-server-service.yaml":    metricsServerMetricsServerServiceYaml,
	"metrics-server/resource-reader.yaml":           metricsServerResourceReaderYaml,
	"rolebindings.yaml":                             rolebindingsYaml,
	"traefik.yaml":                                  traefikYaml,
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
	"ccm.yaml":           &bintree{ccmYaml, map[string]*bintree{}},
	"coredns.yaml":       &bintree{corednsYaml, map[string]*bintree{}},
	"local-storage.yaml": &bintree{localStorageYaml, map[string]*bintree{}},
	"metrics-server": &bintree{nil, map[string]*bintree{
		"aggregated-metrics-reader.yaml": &bintree{metricsServerAggregatedMetricsReaderYaml, map[string]*bintree{}},
		"auth-delegator.yaml":            &bintree{metricsServerAuthDelegatorYaml, map[string]*bintree{}},
		"auth-reader.yaml":               &bintree{metricsServerAuthReaderYaml, map[string]*bintree{}},
		"metrics-apiservice.yaml":        &bintree{metricsServerMetricsApiserviceYaml, map[string]*bintree{}},
		"metrics-server-deployment.yaml": &bintree{metricsServerMetricsServerDeploymentYaml, map[string]*bintree{}},
		"metrics-server-service.yaml":    &bintree{metricsServerMetricsServerServiceYaml, map[string]*bintree{}},
		"resource-reader.yaml":           &bintree{metricsServerResourceReaderYaml, map[string]*bintree{}},
	}},
	"rolebindings.yaml": &bintree{rolebindingsYaml, map[string]*bintree{}},
	"traefik.yaml":      &bintree{traefikYaml, map[string]*bintree{}},
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
