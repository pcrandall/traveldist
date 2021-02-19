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

var _configConfigYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x96\x4d\x6a\x02\x51\x10\x84\xf7\x9e\xa2\x99\x7d\xc4\xf7\xdb\xbe\x39\x44\xee\xe0\x62\x02\x01\xa3\xc1\x88\xb9\x7e\x88\x04\xa2\x56\x75\xb9\xd4\x07\x1f\x1f\x45\x35\x35\xfb\xe5\xb2\xec\xbf\xe6\x95\xd9\x8b\xbd\xed\x8f\xc7\xd3\x6c\x69\x65\x66\x76\xd8\x5d\x96\xf3\x79\x99\xaf\x3f\x7e\x5f\x0f\xbb\x8f\x65\xb6\xe9\x35\xa5\x94\xa6\xbf\x7f\xcd\xde\x3f\x67\x9b\xd2\x66\x9d\x4a\x5f\x27\x5f\xdf\x3e\x9d\x8e\xdf\xb3\x4d\x7d\x42\x42\x56\x84\xfe\x48\xd8\x22\x21\x2b\x87\x02\x0e\x69\x43\x10\x4a\xa2\x80\x44\xca\x88\x28\xca\xa2\xa1\x45\x25\x08\x65\xd1\xd0\x82\xa4\x59\x95\x85\xa3\x05\x89\xb3\x2a\x0b\x07\x8b\x4c\xe2\x6c\xca\x62\x80\x45\x26\x71\x36\x65\x31\xd0\x82\xc4\xd9\x75\x37\x51\x83\xe4\xd9\x65\x3b\xb1\x9e\x79\x3b\xad\x6e\xcf\x27\x3f\x3d\x9f\x1c\xe3\xf3\x23\xfd\xbf\xcc\x77\xf7\x23\x10\x0e\x88\x42\x0f\x28\x46\x14\xb4\x68\xf4\x80\x04\x02\x2d\x9c\x1e\x50\x8c\x68\x68\x31\xe8\x01\x09\x04\x58\x54\x12\x67\x55\x16\x0e\x16\x95\xc4\x59\x95\x85\xa3\x05\x89\xb3\x29\x8b\x81\x16\x24\xce\xa6\x2c\x06\x5a\x90\x38\xbb\x6c\x27\xd6\xb3\x91\x3c\xbb\xac\x27\xf6\xb3\x95\xfb\x03\x2a\x4f\x0f\xa8\xc4\xf8\x02\x74\x3e\x40\x02\xb1\x05\x04\x5f\xa0\x18\x51\xc0\xa2\xf3\x05\x12\x08\xb0\xe8\x7c\x81\x62\x44\x43\x0b\xbe\x40\x02\x81\x16\x7c\x81\x62\x84\xa3\x05\x5f\x20\x81\x00\x0b\xe7\x0b\x14\x23\x06\x58\x38\x5f\x20\x81\x40\x0b\xbe\x40\xa2\x5a\x58\x4f\xe7\x0b\xa4\x18\xe8\xf1\xb0\x40\xf5\xe9\x01\xd5\x18\x5f\xe1\xf3\x8b\x2f\x90\x40\x0c\x40\xf0\x05\x8a\x11\x05\x2d\xf8\x02\x09\x04\x5a\xf0\x05\x8a\x11\x0d\x2d\xf8\x02\x09\x04\x58\x0c\xbe\x40\x31\xc2\xc1\x62\xf0\x05\x12\x08\xb4\xe0\x0b\x14\x23\x06\x5a\xf0\x05\x12\x08\xb4\xe0\x0b\x24\xaa\x85\xf5\x4c\x1b\x3e\x41\x0a\x02\x22\x57\xc8\x4f\x00\x00\x00\xff\xff\x0c\xee\x9b\x11\x0b\x0d\x00\x00")

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

	info := bindataFileInfo{name: "config/config.yml", size: 3339, mode: os.FileMode(420), modTime: time.Unix(1613749386, 0)}
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
