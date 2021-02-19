// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// config/config.yml
package main

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

var _configConfigYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x96\x4d\x6a\x32\x41\x18\x84\xf7\x73\x8a\xa6\xf7\xca\xf4\xbf\x3d\x87\xf8\x36\x5f\x2e\xe0\xa2\x83\x81\x89\x06\x9d\x98\xeb\x87\xfc\x40\xd4\xaa\xae\x59\x6a\xe3\xc3\x43\x51\x2f\xe5\xe5\xd0\xda\x72\xdc\xbf\xb6\xc9\x3c\x9d\xf7\xd7\x36\x6f\xfe\x1f\xde\x97\x65\x6e\xc3\xdc\xae\x6d\xbe\x4c\x83\x31\x1b\xf3\x3c\x9f\x4e\xe7\xc9\xb8\xc1\x18\x63\x8e\xfb\x6b\x5b\x96\x36\x7d\x7f\xf8\x7a\xfd\xf9\xb9\xfd\xe7\x9c\x73\xf6\xf7\x5b\x63\x5e\xde\x26\x63\xdd\xb8\x75\x21\x6f\x5d\xd9\xde\x3e\x9d\x4f\x1f\x93\xb1\xd9\x22\xc1\x2b\x42\x7e\x24\xec\x90\xe0\x95\x43\x00\x07\x37\x12\x84\x92\x08\x20\xe1\x3c\x22\x82\xb2\x48\x68\x11\x09\x42\x59\x24\xb4\x20\x69\x46\x65\x51\xd0\x82\xc4\x19\x95\x45\x01\x0b\x4f\xe2\x4c\xca\xa2\x82\x85\x27\x71\x26\x65\x51\xd1\x82\xc4\x99\x75\x37\x51\x83\xe4\x99\x65\x3b\xb1\x9e\x7e\x67\x87\xdb\xf3\xf1\xab\xe7\xe3\xfb\x78\xff\x48\xff\x2b\xf3\xdd\xfd\x08\x44\x01\x44\xa0\x07\xd4\x47\x04\xb4\x48\xf4\x80\x04\x02\x2d\x0a\x3d\xa0\x3e\x22\xa1\x45\xa5\x07\x24\x10\x60\x11\x49\x9c\x51\x59\x14\xb0\x88\x24\xce\xa8\x2c\x0a\x5a\x90\x38\x93\xb2\xa8\x68\x41\xe2\x4c\xca\xa2\xa2\x05\x89\x33\xcb\x76\x62\x3d\x13\xc9\x33\xcb\x7a\x62\x3f\x53\xb8\x3f\xa0\xb0\x7a\x40\xa1\x8f\x0f\x40\xe7\x03\x24\x10\x3b\x40\xf0\x05\xea\x23\x02\x58\x64\xbe\x40\x02\x01\x16\x99\x2f\x50\x1f\x91\xd0\x82\x2f\x90\x40\xa0\x05\x5f\xa0\x3e\xa2\xa0\x05\x5f\x20\x81\x00\x8b\xc2\x17\xa8\x8f\xa8\x60\x51\xf8\x02\x09\x04\x5a\xf0\x05\x12\xd5\xc2\x7a\x16\xbe\x40\x8a\x81\x1e\x0f\x0b\x14\x57\x0f\x28\xf6\xf1\x11\xfe\x7e\xf1\x05\x12\x88\x0a\x08\xbe\x40\x7d\x44\x40\x0b\xbe\x40\x02\x81\x16\x7c\x81\xfa\x88\x84\x16\x7c\x81\x04\x02\x2c\x2a\x5f\xa0\x3e\xa2\x80\x45\xe5\x0b\x24\x10\x68\xc1\x17\xa8\x8f\xa8\x68\xc1\x17\x48\x20\xd0\x82\x2f\x90\xa8\x16\xd6\xd3\x8d\x7c\x82\x14\x04\x44\xdc\x18\xec\xf0\x19\x00\x00\xff\xff\x71\x62\x11\xa6\x25\x0d\x00\x00")

func configConfigYmlBytes() ([]byte, error) {
	return bindataRead(
		_configConfigYml,
		"config/config.yml",
	)
}

func configConfigYml() (*asset, error) {
	bytes, err := configConfigYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/config.yml", size: 3365, mode: os.FileMode(420), modTime: time.Unix(1613763017, 0)}
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
	"config/config.yml": configConfigYml,
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
	"config": &bintree{nil, map[string]*bintree{
		"config.yml": &bintree{configConfigYml, map[string]*bintree{}},
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
