// Code generated by go-bindata.
// sources:
// files/cf-mgmt.sh
// files/cf-mgmt.yml
// files/pipeline.yml
// files/security-group.json
// files/vars.yml
// DO NOT EDIT!

package generated

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

var _filesCfMgmtSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\xd4\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\x4f\xcd\x2b\x53\x48\x4a\x2c\xce\xe0\x2a\x4e\x2d\x51\xd0\x4d\xe5\x4a\x4e\x51\x48\xce\xcf\x4b\xcb\x4c\xd7\x2d\x4a\x2d\xc8\xe7\x4a\x4e\xd3\xcd\x4d\xcf\x2d\x51\x50\x71\x76\x8b\xf7\x75\xf7\x0d\x89\x77\xf6\xf7\xf5\x75\xf4\x73\xe1\x02\x04\x00\x00\xff\xff\x4b\xdb\x02\xbb\x43\x00\x00\x00")

func filesCfMgmtShBytes() ([]byte, error) {
	return bindataRead(
		_filesCfMgmtSh,
		"files/cf-mgmt.sh",
	)
}

func filesCfMgmtSh() (*asset, error) {
	bytes, err := filesCfMgmtShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/cf-mgmt.sh", size: 67, mode: os.FileMode(493), modTime: time.Unix(1495821632, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesCfMgmtYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\xcd\x41\x8a\xc3\x30\x0c\x85\xe1\xbd\x4f\x21\xb2\x1e\x4f\xf6\xbe\xcc\x60\x3c\x8a\x23\x12\x5b\x46\x92\x43\x43\xe9\xdd\xeb\x94\x52\xe8\xf6\xfd\xf0\x3d\xef\xbd\x6b\x7b\xb4\x85\xa5\x04\xd8\xa9\xf6\x9b\x73\x54\x62\xc6\x3f\x41\xe5\x2e\x09\x83\x03\xb0\xb3\x61\x80\x7f\x4e\x1b\x8a\x7f\xe5\x31\xbe\x33\xdc\x05\x1b\x2b\x19\xcb\x19\xa0\xd1\xc1\x16\x77\x45\x39\x28\xa1\xce\x69\xf1\x25\x17\xfb\x01\x8b\x39\xc0\x34\xae\x50\x6d\x7a\x8c\x93\xda\xba\xe9\x85\x7b\xa8\xb1\x0c\x27\x71\x5d\x28\xfb\x4b\x73\x4e\x7a\xbd\x5a\x8b\xb6\x7e\x95\x39\xd1\x6c\x51\xb7\x8f\xfc\xab\xab\x7b\x06\x00\x00\xff\xff\x0a\x5a\x24\x71\xc6\x00\x00\x00")

func filesCfMgmtYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesCfMgmtYml,
		"files/cf-mgmt.yml",
	)
}

func filesCfMgmtYml() (*asset, error) {
	bytes, err := filesCfMgmtYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/cf-mgmt.yml", size: 198, mode: os.FileMode(420), modTime: time.Unix(1495821632, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesPipelineYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x97\xc1\x8e\xda\x30\x10\x86\xef\x3c\xc5\xbc\x40\x16\xf5\xd0\x4b\x6e\x68\x93\x5d\x45\xda\xc0\x8a\x50\x55\x55\x55\x59\xde\xc4\xa4\x6e\xe3\xd8\xb5\x9d\xad\x10\xe2\xdd\x6b\x27\x2c\x81\x60\x28\x85\xd2\x43\xe8\x2d\xbb\x33\x99\x7f\xe6\xff\xe2\x91\x91\x44\xf1\x4a\xa6\x44\xf9\x03\x0f\x4a\xcc\x88\x0f\x29\x2f\xe7\x34\xf7\x24\x11\x7c\x00\xa0\x17\xc2\xfc\x2f\xa7\xda\x3c\x37\xa9\xbe\x79\x02\xa8\x24\xf5\x61\xb9\x34\x01\x64\x33\x91\xf9\x7b\xb5\xaa\x23\x2f\x12\x97\xe9\x57\x1f\x18\x56\x9a\xc8\x4d\xd9\x77\xef\xd9\xa6\x9c\xa6\x8c\xb4\xf5\x60\x49\x4b\x93\xf9\x8a\x8b\x3a\x6b\x35\x18\x7c\xe3\x2f\xdb\x0d\x49\x82\x35\xf1\xb8\xcc\x95\x79\x49\x14\xb8\xb4\x2d\x78\x90\x13\xdd\xed\xd6\x08\x48\x9a\xe7\x44\x1a\x0d\x59\x91\x3a\x4d\x63\xf5\xbd\x5b\x04\x60\x4e\x8b\xdd\x59\x87\x29\x1d\xda\x54\x35\x4c\xe7\x1e\xcb\x99\xbe\x5b\xb0\xa2\x4e\x15\x58\x62\xa6\x9a\xb1\x01\x92\x4f\xc9\x2c\x8c\x51\x30\x89\x47\xd1\xd8\x5a\xa0\x16\x66\x4e\x86\x32\xce\x30\x2d\xd7\x1e\x00\x7c\x48\xc2\x29\x8a\x02\x9b\x50\x29\x22\x11\xcd\x36\xa1\xe7\x51\x92\x7c\x9c\x4c\xeb\x98\xc0\x4a\xfd\xe4\xb2\x0d\xde\x4f\xc6\x0f\xd1\x23\x0a\xa2\xa9\x0f\x77\xc3\xa6\xbf\xb7\xd0\x53\x14\x8e\x67\x28\x09\xef\xa7\xe1\xcc\xbe\x9c\x16\x94\x94\x1a\x29\x62\x86\xd3\x6d\x85\x07\x14\x3f\xc6\x33\x74\x3f\x89\xe3\xd1\x38\xd8\x1d\xfd\xcd\xd3\x8c\x14\xe4\x02\x4f\xeb\xb4\x86\xe8\x61\xcb\x77\x35\x6e\xca\xf2\xed\xd1\xf7\x3f\x63\x4f\x48\xfa\x6a\x9f\x9b\x01\x4e\x20\x60\x7b\x26\x99\x0f\x9f\xb7\x58\x7e\x39\xf1\x7b\x77\xa8\xf5\x99\xc5\x53\x30\x7a\x46\xdb\x1a\x45\x86\x05\xda\x17\x3a\x7c\x4a\xf6\x0c\xeb\x10\x54\x02\x9b\x85\x79\x45\x68\x1b\x81\xff\x9c\x0e\x72\x5a\x7b\xd4\xd9\x67\xd7\x40\xf3\x27\xbb\xee\x26\xc8\x1d\xda\x76\x1d\x24\x95\xc8\x2e\x39\x2d\xcd\x7b\xc7\xcf\x4b\x57\xe2\xa6\x5c\xdf\x1d\xde\xe5\xba\x67\x5b\xfc\xeb\xd6\x9f\x74\x1e\x9c\x6d\xf4\x19\xcf\xb9\xeb\xcc\x61\x94\x13\xe5\x8f\x8a\x1b\x9f\xae\x7f\x8c\x5a\xa1\x3e\xd3\x3a\x8e\x61\x6d\x81\x93\x83\x29\x64\x7e\xea\xe8\x85\x97\x4b\x5e\x89\x7f\x01\x64\x5f\xf1\x76\xc9\x74\xbd\xe8\x20\xb2\x97\xb7\x33\x77\xde\xc5\x37\x00\x47\x0b\x7d\xe6\x74\xe1\xbe\x6b\x6d\x72\x20\x3c\x77\xd7\xfd\xfe\x82\xed\x12\xe9\x33\xa5\x63\xf6\xaf\x0d\xf8\x15\x00\x00\xff\xff\x86\xc3\x95\x8c\x02\x12\x00\x00")

func filesPipelineYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesPipelineYml,
		"files/pipeline.yml",
	)
}

func filesPipelineYml() (*asset, error) {
	bytes, err := filesPipelineYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/pipeline.yml", size: 4610, mode: os.FileMode(420), modTime: time.Unix(1499893087, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesSecurityGroupJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8a\x8e\xe5\x02\x04\x00\x00\xff\xff\x44\xd2\x68\x70\x03\x00\x00\x00")

func filesSecurityGroupJsonBytes() ([]byte, error) {
	return bindataRead(
		_filesSecurityGroupJson,
		"files/security-group.json",
	)
}

func filesSecurityGroupJson() (*asset, error) {
	bytes, err := filesSecurityGroupJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/security-group.json", size: 3, mode: os.FileMode(420), modTime: time.Unix(1495821632, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesVarsYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x8e\x41\x6e\x43\x21\x0c\x44\xf7\x9c\x62\x6e\xd0\x7d\x55\x71\x15\xe4\xf2\xfd\x53\x4b\xf9\x18\xd9\x46\x51\x6e\x5f\x28\xca\xae\xab\xac\x98\x79\x0c\xc3\xdc\x24\x8a\x71\xd7\x32\x4c\x3e\xf1\xf5\xd4\x61\x98\x0c\x8b\x61\xb2\x9c\xfc\xe9\xc1\x57\x39\xf4\x22\x69\xaf\x48\x3d\xb1\x39\x36\xcf\x69\x38\x5b\x91\x63\x06\x96\x02\xd5\xaa\xa3\x05\x1e\x12\x3f\xe8\x6c\x97\xb8\x8b\x36\x84\xa2\x1a\x53\x30\xd4\x6e\xfe\xe1\x9d\x2a\x7b\x4e\x9d\xdc\x1f\x6a\xeb\xf9\x4b\x42\x4f\xbc\x53\x55\xef\xc2\x2d\x8a\xf3\xbc\x8c\xd9\xb7\x3d\xb6\xc7\xa9\x86\x41\x54\xb7\xd8\x9b\x73\xba\x1f\xd4\xcb\x7f\x1b\xe6\x1f\xdf\xd2\xfe\xce\x95\xc9\xe9\x37\x00\x00\xff\xff\x27\x4c\x31\xa2\x30\x01\x00\x00")

func filesVarsYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesVarsYml,
		"files/vars.yml",
	)
}

func filesVarsYml() (*asset, error) {
	bytes, err := filesVarsYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/vars.yml", size: 304, mode: os.FileMode(420), modTime: time.Unix(1495821632, 0)}
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
	"files/cf-mgmt.sh": filesCfMgmtSh,
	"files/cf-mgmt.yml": filesCfMgmtYml,
	"files/pipeline.yml": filesPipelineYml,
	"files/security-group.json": filesSecurityGroupJson,
	"files/vars.yml": filesVarsYml,
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
	"files": &bintree{nil, map[string]*bintree{
		"cf-mgmt.sh": &bintree{filesCfMgmtSh, map[string]*bintree{}},
		"cf-mgmt.yml": &bintree{filesCfMgmtYml, map[string]*bintree{}},
		"pipeline.yml": &bintree{filesPipelineYml, map[string]*bintree{}},
		"security-group.json": &bintree{filesSecurityGroupJson, map[string]*bintree{}},
		"vars.yml": &bintree{filesVarsYml, map[string]*bintree{}},
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
