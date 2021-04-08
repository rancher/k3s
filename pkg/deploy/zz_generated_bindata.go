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

var _corednsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x5d\x6f\xdb\xb8\x12\x7d\xf7\xaf\x20\x74\xd1\xb7\x2b\x27\xbe\x41\x7b\xb3\x7c\x6b\x93\x6c\x1b\xa0\x71\x8d\x38\xe9\xcb\x62\x51\x8c\xa9\xb1\xc5\x0d\xc5\xe1\x92\x94\x1b\x6f\x37\xff\x7d\x41\x7d\x99\x74\x94\x22\x0d\xba\x7e\xb1\xc8\xe1\x9c\xa1\x0e\x87\x67\x46\x60\xe4\x67\xb4\x4e\x92\xe6\x6c\x3b\x9b\xdc\x49\x5d\x70\xb6\x44\xbb\x95\x02\xdf\x0a\x41\xb5\xf6\x93\x0a\x3d\x14\xe0\x81\x4f\x18\xd3\x50\x21\x67\x82\x2c\x16\xda\x75\x63\x67\x40\x20\x67\x77\xf5\x0a\x73\xb7\x73\x1e\xab\x49\x9e\xe7\x93\x18\xda\xae\x40\x4c\xa1\xf6\x25\x59\xf9\x17\x78\x49\x7a\x7a\x77\xea\xa6\x92\x8e\xb6\xb3\x15\x7a\xe8\x23\x9f\xa9\xda\x79\xb4\xd7\xa4\x30\x09\xab\x60\x85\xca\x85\x27\xd6\xc4\xb1\x1a\x3d\x36\xfe\x2b\x22\xef\xbc\x05\x63\xa4\xde\xb4\x81\xf2\x02\xd7\x50\x2b\xef\x86\xfd\xb6\xbb\xe2\xfd\xb6\x6d\xad\xd0\xf1\x49\xce\xc0\xc8\xf7\x96\x6a\xd3\x20\xe7\x2c\xcb\x26\x8c\x59\x74\x54\x5b\x81\xdd\x1c\xea\xc2\x90\xd4\x0d\x58\xce\x5c\xcb\x4c\x3b\x30\x54\xb4\x0f\x03\x09\x61\xb8\x45\xbb\xea\x7c\x95\x74\xbe\x79\xf8\x0a\x5e\x94\xcf\x8b\xa7\xa9\x38\x84\xd9\xa0\xff\x19\x84\xbe\x93\xba\x90\x7a\x93\xf0\x0a\x5a\x93\x6f\xdc\x3b\x72\xc7\x70\x13\xbe\xa1\xf6\x54\x9b\x02\x3c\x72\x96\x79\x5b\x63\xf6\xf3\x8f\x87\x14\x5e\xe3\xba\xd9\x5f\x47\xd8\x77\x5e\x78\xc2\xd8\xe3\xdc\x79\x02\xd9\xd5\xab\x3f\x50\xf8\xe6\xec\x47\x53\xfd\xc5\x09\x3e\xdc\x9d\x33\xd2\x6b\xb9\xb9\x02\xf3\x92\x6b\xd3\x2f\x3f\x23\x8b\x6b\xa9\x90\xb3\xbf\x1b\x4e\xa7\xfc\xf5\x09\xfb\xd6\x3c\x86\x1f\x5a\x4b\xd6\x0d\xc3\x12\x41\xf9\x72\x18\x5a\x84\x62\x37\x8c\xf6\xc7\xc1\x5e\x7d\x3b\xfb\x78\xbb\xbc\xb9\xb8\xfe\x72\xfe\xe9\xea\xed\xe5\xfc\xe1\x15\x93\x3a\x87\xa2\xb0\x53\xb0\x06\x98\x34\x6f\xda\x87\x7d\x24\xd6\x24\x39\x93\xda\xa1\xa8\x2d\x46\xf3\x6b\x50\xca\x97\x96\xea\x4d\x39\x8e\x32\xac\x7d\xd8\x6f\x94\x9c\x77\xec\x08\xbd\x38\xea\xa8\x38\x9a\x53\x81\x1f\x9a\xe9\x38\xa8\xf7\x8a\xbd\x39\x8e\x26\x2c\x2a\x82\x82\xcd\x5e\xbb\xf1\x2d\x8c\x04\x33\x96\x2a\xf4\x25\xd6\x8e\xf1\x5f\x66\xaf\x4f\x06\xc3\x9a\xec\x57\xb0\x05\x9b\xb6\x3b\x09\xf7\x4f\x6d\xa7\x82\xf4\x7a\x58\x22\x40\x94\xc8\x4e\xf6\x3b\x50\x44\x66\x92\x6e\x26\xb2\x41\xb1\x02\x05\x5a\xb4\xfc\x3c\x3c\x4a\x0e\x30\xc6\x1d\x0d\x19\x72\x8e\x46\xd1\xae\xc2\x97\x29\xeb\xc1\x65\x3b\x75\x39\x18\xd3\x2d\x69\x1d\x0f\xaf\x60\x0b\x9c\x85\x9c\x3a\x9f\x2f\xb3\x89\x33\x28\x82\xf7\x7f\x2c\x1a\x25\x05\x38\xce\x66\x13\xc6\xc2\x2d\xf5\xb8\xd9\xb5\xc0\x7e\x67\x90\xb3\x6b\x52\x4a\xea\xcd\x6d\x73\xdf\x5b\x7d\x88\x67\x78\xc7\x41\x05\xf7\xb7\x1a\xb6\x20\x15\xac\x42\xd2\x36\x70\xa8\x50\x78\xb2\xed\x9a\x2a\x08\xe0\xc7\x68\xe3\xe3\x5b\xf7\x58\x19\x35\x00\xc7\xec\x34\x44\x27\xfe\x4f\xbd\x7c\xff\x7a\x6d\x0e\x48\xb2\xd2\xef\xce\x14\x38\x37\x6f\x79\x68\x79\xcc\x45\xab\x16\xb9\xb0\xd2\x4b\x01\x2a\xeb\x5c\x5c\x22\x08\xf3\x83\x43\x69\xa8\x21\x85\x36\xd6\xcc\xf0\xcb\xd9\x1d\xee\x02\xcb\x1d\xdc\xdb\xa2\x20\xed\x3e\x69\xb5\xcb\xa2\x8c\x25\x13\x3c\xc9\x72\x96\x5d\xdc\x4b\xe7\x5d\xf6\x08\x20\xe8\x7f\x1e\x14\xf0\x40\x77\x05\x69\x6f\x49\xe5\x46\x81\xc6\x67\x62\x32\x86\xeb\x35\x0a\xcf\x59\x36\xa7\xa5\x28\xb1\xa8\x15\x3e\x3f\x64\x05\x81\xa1\x9f\x11\x2b\x44\x58\x26\x09\x11\x7e\xa1\x4e\x1d\x84\x24\xc7\x99\x92\xba\xbe\xef\x16\x85\xb7\x06\xa9\xd1\x0e\x54\xe7\x8f\x2e\x4a\xfb\x93\x15\x6c\x90\x33\x0b\x5a\x94\x68\x7b\x7d\xc9\xbb\x7f\x3e\x9b\x9e\x4e\x4f\xd2\xc5\x8b\x5a\xa9\x05\x29\x29\x76\x9c\x5d\xae\xe7\xe4\x17\x16\x1d\x36\x45\xa0\xbf\xe7\x51\x65\x1e\x6e\xbb\xac\xa4\x4f\x66\x42\xa2\x56\x64\x77\x9c\xcd\xfe\x7f\x7c\x25\x13\xd5\xfa\xb3\x46\x77\xb8\x5a\x98\x9a\xb3\xd9\xf1\x71\x35\x8a\x91\x40\x80\xdd\x38\xce\x7e\x63\x59\x1e\xe4\x29\xfb\x2f\xcb\x12\xf1\xec\xcb\x44\xc6\x7e\x1f\x5c\xb6\xa4\xea\x0a\xaf\x42\xf2\x26\xe9\xd9\xb3\x16\xaa\x53\xde\x2e\x8a\xe2\x57\x61\xfd\x02\x7c\xc9\x13\x79\x4e\xde\x05\x8a\x90\xce\x9c\x85\xa2\xbf\x57\x59\xb2\x69\x9c\xe1\xc4\x16\x64\x3d\x67\x91\xee\xf6\x12\x97\xe2\x1a\x4b\x9e\x04\x29\xce\x6e\xcf\x17\x3f\x8a\x93\x7b\x61\x46\xb1\x6e\xce\xbe\x83\x95\x54\x83\x1e\xad\x42\x6f\xa5\x18\xdf\x59\x8c\xd6\x14\xc2\x20\x29\xa4\x3d\xde\xfb\xf8\x68\x41\x29\xfa\xba\xb0\x72\x2b\x15\x6e\xf0\xc2\x09\x50\x8d\x4c\xf0\x50\xa9\x5c\x4c\xb7\x00\x03\x2b\xa9\xa4\x97\x78\x90\x1c\x50\x14\xe9\x44\xce\xe6\x17\x37\x5f\xde\x5d\xce\xcf\xbf\x2c\x2f\xae\x3f\x5f\x9e\x5d\x24\xe6\xc2\x92\x39\x74\x00\xa5\x46\x0e\xee\x9a\xc8\xff\x2a\x15\x76\x2d\x51\x7a\x8c\x4a\x6e\x51\xa3\x73\x0b\x4b\x2b\x8c\xf1\x4a\xef\xcd\x7b\xf4\x69\x08\xd3\x26\xca\x41\xdf\xc1\xba\x74\xe0\xec\xf4\xf8\xf4\x38\x99\x76\xa2\xc4\x40\xf2\x87\x9b\x9b\x45\x64\x90\x5a\x7a\x09\xea\x1c\x15\xec\x96\x28\x48\x17\x8e\xa7\x75\xdf\xa0\x95\x54\x0c\xb6\x59\x6c\xf3\xb2\x42\xaa\xfd\xde\x18\xd9\x5c\x2d\x04\x3a\x77\x53\x5a\x74\x25\xa9\x22\xb5\xae\x41\xaa\xda\x62\x64\x3d\x49\xba\x27\xf9\xc3\x54\xa4\x3d\x57\xc4\xc4\xec\x74\xf6\x62\x26\xbe\x43\xc4\xff\xfe\x65\x1e\x0a\xed\x7a\x69\x3c\x6f\xbb\xf5\xce\xd0\x2a\xc7\x0f\x28\x8b\xe8\xfb\xe1\x94\xb7\x71\x11\x6f\xa8\xf0\x58\xb9\xc3\x94\x6e\x0a\x55\x2f\x77\x89\xad\x3f\x82\x51\x63\xe7\x38\x34\x99\xa3\x9e\x7b\xeb\x93\x4d\x7d\xf7\x95\x30\xd2\xaf\x45\xad\xc7\x93\x0d\xdb\xa3\x8f\xac\x7d\x6b\x1a\x6a\x5e\x9b\x29\x59\x50\xa5\x6c\xc4\xec\x84\x05\xf3\xe4\xc7\xd6\x33\xfa\xbf\xbe\xd3\xe9\x3a\x9b\x08\xe9\xb9\x9d\x62\xda\xcb\x8d\xc5\xec\x62\x5c\x2e\x78\xfc\x95\x31\x5f\x3e\xbc\x9a\x44\x35\x22\x3f\xa8\x00\x26\x96\xf6\xc3\x42\x90\x8f\xc8\xfc\x13\x0e\xad\x3e\xe7\x23\x4a\x6e\x52\xc1\x4f\x5d\xfe\x09\x00\x00\xff\xff\x2e\x93\x58\xfe\xfc\x10\x00\x00")

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

var _localStorageYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x4d\x6f\xe3\x36\x13\xbe\xeb\x57\xcc\xab\x77\xf7\xd0\x62\x69\xaf\xdb\xc3\xb6\x5c\xf4\xe0\xb5\x9d\x34\x40\x62\x1b\x71\xda\x1e\x82\x85\x41\x51\x63\x9b\x1b\x8a\x24\x48\xca\x59\x6f\x9a\xff\x5e\x90\x92\x1d\xc9\x71\x12\x07\x6d\x6f\xd5\x45\xe0\x70\x9e\xf9\xfe\x20\x33\xe2\x77\xb4\x4e\x68\x45\x61\xdd\x4b\x6e\x84\xca\x29\xcc\xd0\xae\x05\xc7\x3e\xe7\xba\x54\x3e\x29\xd0\xb3\x9c\x79\x46\x13\x00\xc5\x0a\xa4\x20\x35\x67\x92\x18\xe6\x57\xc4\x58\xbd\x16\x01\x8f\x96\xb8\x0a\x47\x58\x0d\xac\xd8\x9d\x61\x1c\x29\xdc\x94\x19\x12\xb7\x71\x1e\x8b\x84\x10\x92\x34\x35\xdb\x8c\xf1\x0e\x2b\xfd\x4a\x5b\xf1\x8d\x79\xa1\x55\xe7\xe6\x27\xd7\x11\xba\xbb\xee\x65\xe8\xd9\xd6\xb0\x81\x2c\x9d\x47\x7b\xa9\x25\x1e\x6f\x95\x0d\xdc\xb6\x94\xe8\x68\x42\x80\x19\x71\x6a\x75\x69\x1c\x85\xeb\x34\xfd\x9c\x00\x58\x74\xba\xb4\x1c\x23\x45\xe9\x1c\x5d\xfa\x0e\x52\x13\x6c\x73\x1e\x95\x5f\x6b\x59\x16\xc8\x25\x13\x45\xbc\xe1\x5a\x2d\xc4\xb2\x60\xc6\x45\xf8\x1a\x6d\x16\xa1\x4b\xf4\xe1\x5a\x0a\x17\xff\xb7\xcc\xf3\x55\xfa\xf9\x65\x95\xa8\x72\xa3\x85\xf2\x07\xd5\x56\x44\x9d\xef\xe9\xfa\xfe\x28\xc1\x6b\x0c\x52\x5b\x40\x6e\x91\x79\x8c\x42\x0f\xdb\xe7\xbc\xb6\x6c\x89\x75\xfc\x1f\x0b\xad\xef\xb9\x64\xce\xe1\x91\x11\xf8\xfb\xd9\xfe\x24\x54\x2e\xd4\xf2\xf8\xa4\x67\x42\xe5\x49\xc8\xfc\x25\x2e\x02\xf3\xd6\xc7\x67\xb4\x27\x00\x8f\xab\xec\x98\xda\x72\x65\xf6\x05\xb9\x8f\xe5\x75\xb0\x81\xfe\xad\xb6\x61\xc6\xb8\xee\xae\x6b\x87\x68\xa4\xde\x14\xf8\x8a\x8e\x7d\x5a\x95\x33\xc8\x69\xcc\xbd\x91\x82\x33\x47\xa1\x97\x00\x38\x94\xc8\xbd\xb6\xe1\x06\xa0\x08\xf9\x3d\x67\x19\x4a\x57\x11\x42\x98\xcd\x33\xba\x3c\x16\x46\x32\x8f\x35\xbc\x61\x64\xf8\x64\x4b\xd2\x4b\xb2\x00\xb6\x26\x86\xcf\x58\xa1\xad\xf0\x9b\x41\x28\xcb\x71\xf4\x38\xad\x3c\x21\xa1\xa3\x09\xb7\xc2\x0b\xce\x64\x5a\xf3\xbb\x56\x82\xc6\xaf\xcb\x4e\xf8\xbc\x96\x68\x63\xf5\x34\x2c\x06\x20\x70\x83\x1b\x0a\xe9\xa0\xd6\xd7\xcf\x73\xad\xdc\x44\xc9\x4d\xda\xe0\x02\xd0\x26\xa0\xb5\xa5\x90\x8e\xbe\x0a\xe7\x5d\x7a\x40\x48\xb4\x3c\x54\x58\x27\x64\xc6\x2a\xf4\x18\xbb\x84\x6b\xe5\xad\x96\xc4\x48\xa6\xf0\x15\x72\x01\x70\xb1\x40\xee\x29\xa4\x63\x3d\xe3\x2b\xcc\x4b\x89\xaf\x51\x5c\xb0\xd0\x17\xff\x94\xc6\xe0\x06\x13\x0a\xed\x2e\x82\xe4\xa5\x62\xad\x3e\x51\xb0\x25\x52\xb0\x4c\xf1\x15\xda\xee\x61\x6e\xba\x7e\xdf\x79\xdf\xe9\xfd\xdc\x46\x4d\x4b\x29\xa7\x5a\x0a\xbe\xa1\x70\xb6\x18\x6b\x3f\xb5\xe8\x70\x97\xd5\x60\x54\x51\x30\x95\x3f\xe4\x94\xbc\x64\x0d\x01\xe7\x99\xf5\x8d\x33\x21\xd5\x8a\x68\x90\xba\xe8\x79\xb7\xa2\xd6\xbf\xce\x17\xa7\xd5\x8e\xa3\x1a\xf6\x17\xa1\xc0\x5c\x53\x77\x15\x8f\x0a\x41\x2a\xa6\x46\x78\x8b\xc0\x3f\x65\x7e\x45\x5b\x0a\x76\x1c\xa8\xd6\x8f\x85\x4d\x27\xc3\xf9\xb8\x7f\x31\x9a\x4d\xfb\x83\x51\x43\xd8\x9a\xc9\x12\x4f\xac\x2e\x68\x2b\x81\x0b\x81\x32\xaf\x87\xe8\x23\x7a\xa5\x7b\xdb\xc8\x9d\xdd\x2c\x49\x9a\x5e\xbd\xc2\xa1\x8a\x7e\xc1\x4c\x5b\xdb\xa3\xaa\xa8\xe3\xbb\x3f\x0f\xdb\xbb\xeb\x61\x32\xce\x2a\x7a\x1c\x0e\xcf\xce\xc6\xb0\x28\x94\xd2\xbe\xd9\xd8\xcd\x85\xb7\xd7\x0f\xc2\x91\x1c\x17\xac\x94\x9e\xc4\x6b\x0a\xa9\xb7\x25\xa6\x49\xb3\x0e\xb7\x75\x1a\x00\x0d\x4d\x95\xef\xf5\x5e\xbb\xd0\x39\x52\xf8\x83\x09\x7f\xa2\xed\x89\xb0\xce\x0f\xb4\x72\x65\x81\x36\xb1\xd5\xcb\x63\x5b\xb4\x43\x94\xe8\x31\x7a\x5e\x2f\xab\x6d\xc8\x92\xbd\xa7\xdc\xb3\x3b\x60\x57\xa0\x4f\x8c\xff\x2d\xb0\x51\xab\x14\xfe\x24\x31\x20\x77\x75\x6e\xe2\x98\x08\x15\x70\xc1\x4c\x4a\xaf\x6b\xea\xdd\x2e\x73\xf1\x3e\xa5\xe9\x70\x74\xd2\xff\xed\xfc\x6a\x3e\xed\x5f\xfd\x3a\x3f\x99\x5c\xce\xc7\x93\xf1\xfc\xfc\x6c\x76\x35\x1a\xce\xc7\x93\xe1\x68\x96\xbe\x7b\xc0\x04\xeb\x5c\x4a\xaf\xd3\xb7\x77\x5b\xdc\xf9\x64\xd0\x3f\x9f\xcf\xae\x26\x97\xfd\xd3\x51\x94\x72\xff\x36\xbe\x3b\xc2\x77\x5f\xff\xab\xf3\x7d\xdc\x51\x3e\xac\xf9\xda\xd8\xff\xff\xaf\x9b\x09\xd5\x75\xab\x78\xba\x5d\x09\x89\xb0\x44\xaf\x8d\x77\x90\x16\xd4\x51\x43\x53\xd0\xa6\x6a\xdf\x5c\x3f\xcc\x01\xe6\x10\xde\x68\xe3\x41\xa8\x56\x2d\x9a\xef\x5a\x47\x96\x39\x2d\x4b\x1f\xe3\xf0\xcb\x9b\xc9\xf4\xaa\x7f\x79\xda\x62\xf8\xf8\xb1\x75\x74\x6d\xb8\x13\xdf\xf0\x4c\x7d\xda\x78\x74\xc7\xa0\x8b\x36\x7a\xad\x65\xa8\x9c\x97\x90\xe8\x18\xaf\xfd\x53\x55\xb7\x15\x37\xb9\xb0\x40\x0a\x78\xff\xe1\xc3\x07\x20\x06\xde\xdc\x35\x1d\xb9\x8f\xbb\x9a\xd9\x5c\xdf\xaa\xff\x22\xf9\x6c\x24\x6d\x01\xc4\x2e\x0e\xc4\x6f\x85\xd2\xa0\x9d\xea\xbc\xb3\x61\x85\xdc\x45\x71\xaf\x49\x03\xa9\xea\xe3\xa9\xce\x0f\xbe\x8a\xaa\xd6\xad\xa4\x11\x53\x33\x35\x9f\x3e\x4f\xaf\xd1\x3d\x10\x1c\x58\x9d\x22\xb3\xcc\x6e\x48\x56\xba\x4d\xa6\xbf\xd2\x5e\xe7\xc7\x1f\x3a\xbd\xe4\xaf\x00\x00\x00\xff\xff\xee\x9c\x92\x0b\x15\x0e\x00\x00")

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

var _traefikYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x91\x5d\x8b\xdb\x3c\x10\x85\xef\xfd\x2b\x06\xc3\x5e\xca\x7e\x77\x79\x2f\xb6\xba\x4b\x43\xa0\x4b\x21\x2d\x4d\xdb\xdb\x32\x96\x27\xb1\x88\x2c\x89\x99\x71\xda\xf4\xe3\xbf\x17\xe5\x6b\x53\x58\xd8\x52\x5a\xdf\x79\x34\xe7\x99\xc3\x39\x98\xfd\x47\x62\xf1\x29\x5a\x18\x28\x8c\x8d\x43\xd5\x40\x8d\x4f\xed\xee\xb6\xda\xfa\xd8\x5b\x78\x45\x61\x9c\x0f\xc8\x5a\x8d\xa4\xd8\xa3\xa2\xad\x00\x22\x8e\x64\x41\x19\x69\xed\xb7\xc6\x71\x7f\x9a\x49\x46\x47\x16\xb6\x53\x47\x46\xf6\xa2\x34\x56\x92\xc9\x15\x89\x2b\x10\x0b\x83\x6a\x16\xdb\xb6\x37\xdf\x5e\x7f\x78\xb9\x78\xb7\x5c\xbc\x5f\xac\x3e\xcd\xde\x3e\xfc\xb8\x69\x45\x51\xbd\x6b\x0f\x8b\xd2\x5e\xc1\xcd\x8b\xe6\xf6\xff\xe6\xae\xd1\xcd\xd7\xca\x18\x53\xfd\x25\xdf\xff\xce\xf3\x95\x5f\x80\x1d\x86\x89\x64\x9e\xa2\x52\x54\x0b\xdf\x4d\x05\x00\xc0\x1d\x1e\x4e\x94\x8f\x22\x76\x81\xfa\x62\x6c\xa2\xc3\x2c\x27\x56\x39\x3f\x7f\xa6\x4e\xc8\x4d\x4c\xe7\x01\x80\x06\x79\xfc\x79\x1a\xd0\xcf\x62\x4c\xc5\x5d\x8a\x97\xdd\xcc\x69\x24\x1d\x68\x92\x92\x55\x39\x62\xa1\xbe\xff\xef\xfe\xae\x7e\x72\x41\x1c\x63\x26\x0b\x75\xc1\x1e\x57\x32\xa7\x9d\xef\x89\x2f\xc8\x12\x1b\x47\x52\x92\x87\xb8\x61\x92\x2b\x5f\x79\xea\x82\x97\x81\xfa\x15\xf1\xce\x3b\x7a\xc6\x31\xfb\xc4\x5e\xf7\xf3\x80\x22\xcb\x43\x4f\xf5\xb1\x0e\xe3\xc2\x24\x4a\x6c\x1c\x7b\xf5\x0e\xc3\xd1\x8a\x1f\x71\x73\x61\x1e\x8b\xad\x19\xa3\x1b\x88\xdb\xe0\x3b\x46\xde\x9b\x53\x21\x47\x81\xa6\x40\x7c\x1d\x88\x81\x2d\xed\x2d\xd4\xf3\x13\x77\xd6\xf7\x29\xca\x9b\x18\xf6\xe7\x3c\x52\x2e\x8a\xc4\x16\xea\xc5\x17\x2f\x2a\xf5\x2f\xc2\x98\x7a\x32\x9c\x02\x35\x8f\x31\x94\xe0\x5c\x8a\xca\x29\x98\x1c\x30\xd2\x33\x2c\x00\x5a\xaf\xc9\x95\x26\x96\x69\xe5\x06\xea\xa7\x40\xbf\x77\x66\xc4\x12\xcb\x9f\xf1\x7f\x06\x00\x00\xff\xff\xd5\xb5\xd7\x89\xfc\x03\x00\x00")

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
