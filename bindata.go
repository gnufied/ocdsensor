// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// assets/home.html
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _assetsHomeHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\x6b\x6f\xdb\x3a\x12\xfd\xde\x5f\x31\x50\xb0\xb0\x8d\x55\x64\x3b\x69\x8d\xae\x2a\x07\x68\xfc\xc0\x76\xd1\x34\x41\x1f\xd8\x0f\x41\x10\xd0\x22\x1d\xf3\x96\x22\x55\x92\xb6\xe5\x1b\xf8\xbf\x5f\x50\xd6\x83\x92\x1f\x51\xd0\x0b\xdc\x04\x81\x65\xce\xf0\xcc\x99\x33\x1c\x8a\x4c\xb0\xd0\x11\xbb\x7a\x03\x10\x2c\x08\xc2\xe6\x01\x20\x88\x88\x46\x10\x2e\x90\x54\x44\x0f\x9d\xa5\x9e\x9f\xbf\x77\x32\x93\x0a\x25\x8d\x35\x28\x19\x0e\x9d\x85\xd6\xb1\xf2\xbb\xdd\x10\x73\x2f\x66\x42\x7b\x6c\xd3\x35\x9f\x6c\x73\xce\x90\x26\x4a\x7b\x11\xe5\xde\x1f\xca\xb9\x0a\xba\xbb\x79\x39\x88\xde\x30\xb2\x7b\x06\xef\xcf\x25\x3d\xd7\x68\xc6\x08\x3c\xef\x46\x00\xd6\x14\xeb\x85\x0f\xfd\x5e\xef\x5f\x1f\xf2\xb1\x99\x90\x98\x48\x1f\x94\x60\x14\x43\x3f\x4e\xe0\x6c\x3c\x9e\x4c\x26\x93\x9a\xc7\x79\x28\x18\x43\xb1\x22\x3e\xe4\x4f\x75\x0f\x15\xa3\x90\xf2\x27\x1f\x7a\x85\x65\x2e\xb8\xf6\x81\x0b\x19\x21\x06\xfd\xcb\x38\x81\x8f\x92\x22\xe6\x82\x42\x5c\x9d\x2b\x22\xe9\x3c\xf3\xdd\xee\xd1\xf6\xc8\x8a\xf0\x92\xfc\x0c\x85\x3f\x9f\xa4\x58\x72\x6c\xa8\x08\xe9\xc3\xd9\x68\x60\x7e\x8f\x03\x08\x8c\x4f\xce\x9f\xa6\x3f\x47\xe7\x6b\x53\x3a\xd0\x8b\x93\x18\xe3\xf1\x64\x3a\x99\xbe\x42\xce\x7c\xe2\xe5\xe5\xe0\x7a\x70\x5d\x0c\xc7\x08\xe3\x54\xbc\x7e\x2f\x4e\x8a\x51\x4d\x12\x7d\x8e\x18\x7d\xe2\x3e\x30\x32\xd7\x55\x83\x5a\x20\x2c\xd6\x7e\x1a\x26\xff\x3b\x9b\xcf\x8f\x4b\xaa\x67\x02\x6f\x40\xdb\xaa\xbc\x82\xef\x65\x13\xb2\x8d\x38\x05\x5d\x6b\xa9\x06\xf6\x1a\x86\x15\x92\x10\x2e\xa5\x24\x5c\x7f\x27\x51\xfc\x61\x6f\x74\x2a\xc9\xaf\xfd\xd1\x1b\x82\xf8\x61\x7f\x63\x19\xdd\xfd\xd8\x4d\x7b\x93\xf5\x01\xc7\x62\xed\x09\xce\x04\xc2\x30\x84\xf9\x92\x87\x9a\x0a\xde\xee\x94\xc2\x18\x04\x4d\x23\x02\x43\xe0\x64\x0d\x63\xa4\x49\xbb\x93\x03\x64\x01\xe2\xa5\x09\x09\xc3\x72\x16\x40\xe2\xc3\xfd\x83\x5b\x7e\xdf\xd4\xbe\x47\x02\x13\x1f\x5a\x8c\x72\xa2\x5a\xd6\x38\x47\x11\xf1\xc1\x19\xdd\xfd\x00\x03\xea\x58\x26\xe3\xeb\xdb\x31\x8a\xa2\xb4\xce\xde\xf7\x46\x1f\xa7\x03\x1b\x08\x40\x2d\x50\x6c\x62\xa8\xd8\xcc\x6c\x95\xa6\xed\x9b\xe2\xc1\xce\x23\x42\xa3\xbb\x1f\x7f\x7b\x26\x46\x77\x78\x6d\x3a\x17\xd7\xff\x79\xdb\x1f\x57\xd3\xc1\x48\x2d\x7c\x68\x61\xa1\xed\xf1\x23\xb9\x84\xf1\xd2\x54\xfa\x75\x99\x24\x28\xa1\xca\x87\x56\x72\x61\x47\xd8\x64\xa3\x9b\xca\x68\x83\xfa\x4d\x25\x59\x12\x1e\x6e\x0e\x24\x5d\x24\x3a\x9e\xbe\x1b\x4c\xfb\xad\x32\x8b\xea\xd2\x8a\xca\x45\xfb\x8f\xa5\x52\x14\xd0\xb0\x68\x5a\xc0\xd1\xe0\xe2\x7a\xf4\xee\x37\x0a\xc8\xd0\x46\x2c\x75\x2d\xe9\x1d\xfd\x4a\x44\xbd\x49\xd7\x38\x46\x9a\xd4\xa2\x89\x08\x51\xee\xc3\x7d\xcf\x85\xfe\x43\xad\x31\xc4\x5a\xd3\xf0\x27\x43\x33\xc2\x94\x0f\x73\xc4\x14\xb1\x18\xed\x2b\xf6\x5c\xa2\x79\x03\xb7\xff\xb0\xad\x4b\x7d\xd1\x94\x16\xe2\xe1\x22\x95\xa8\x5a\x83\x3a\xdf\x13\x6c\xea\xa1\x0a\xc4\xe4\x04\x62\xcf\x7b\x6b\x91\xae\x4a\x8d\x91\x46\x30\x84\xfb\x6c\x1f\x73\xcb\x8d\xc0\xcd\xda\xc8\xb5\x17\xe2\x43\xb9\x44\xef\xd2\xd3\x48\x7a\x38\x69\xb7\x9e\x24\x8a\x17\x2d\x37\x85\x73\xb3\xfa\xd5\x77\x4a\x6e\x2a\xda\xab\x0e\x52\xae\x89\x5c\x21\x06\x43\x50\x44\x7f\xca\xbe\xb5\x2b\x9b\x71\x99\xd6\x0b\xfb\xf1\xce\x61\x19\x1b\xe1\xab\x8b\x67\xd7\x33\xf7\x66\xee\x83\x0b\x87\x3f\xab\xcb\xc4\xf4\xd4\xbd\xf5\x0e\x32\x6e\xb5\xd7\x8c\x35\x94\x4a\x53\xf5\xc8\x05\xb3\xab\x69\x31\xcd\xd4\x23\x89\x26\x1c\x7f\x97\x28\x24\xaa\x54\x71\x97\x82\x6b\x8a\xd7\x77\x2f\xdc\xcb\x87\x8e\x35\x93\xce\xdb\xa9\x94\xc3\xa1\x39\xc4\x75\x20\x64\x04\xc9\x42\xb9\x5c\xd0\x4e\xf1\x46\xde\xba\xc6\xaf\x57\xab\xc6\x5a\x15\x0e\x6b\x95\xe9\xf9\x7f\x32\xfb\x26\xc2\x9f\x44\xb7\x9d\xe7\x67\x6f\xbb\x75\x3a\x96\x8f\x27\xb8\x88\x09\xb7\xdf\x94\x64\xa5\x3b\xb6\xca\xa1\xe0\x4a\x30\xe2\x51\x3e\x17\x6d\xe7\xf6\x6e\xf2\xc5\x42\x48\x31\x14\xe1\xb8\xed\x2c\x08\x63\xc2\x32\x6d\x2b\x51\x42\x26\x14\x69\x1e\x66\xf4\xf9\xf6\xdb\xa4\x16\xc7\xe4\xb3\x64\xec\x48\x80\x88\x28\x85\x9e\x4e\x86\xc8\xba\xe2\x7f\xdf\x6e\xbf\x78\xb1\x39\xaa\x1b\x0f\xcf\x8c\xda\x81\xac\xe5\x01\xc3\x74\x8e\xa7\xcb\xe3\x87\xe5\x91\x6d\xe1\xa9\xc7\x5c\x92\x5f\xe9\x9b\x61\xdf\x2d\x5f\x58\x36\xd8\x63\x84\x2e\x7b\x87\x5d\xcb\x77\x43\xea\x6d\xba\xf4\x11\x31\xf6\x38\x2f\x0f\x47\x79\xb5\x4d\x32\x58\x84\xcb\x88\x70\xed\x3d\x11\x3d\x61\xc4\x3c\x5e\x6f\x3e\xe1\xb6\x93\x41\x3e\x9a\x70\x4e\xa7\x9c\x88\x3c\xca\x39\x91\xff\xfd\x7e\xf3\x19\x86\x76\xb2\x55\xec\xd9\x29\xec\x88\xf2\x3d\xdc\x59\x05\x77\xc7\x3d\x73\xb3\x2a\x70\x14\x11\x25\x19\xe2\x01\x98\xcc\xd6\x04\xc6\xa8\x75\x14\x27\x37\x36\x00\xca\xd5\x33\xaa\xd7\xb0\xac\xea\x37\x61\x44\xf9\x21\x90\x42\x9f\x79\x43\x18\x94\x1c\x87\xc9\x6c\x4d\xf5\x39\x8a\x93\x1b\x0f\x37\x18\x91\x52\xc8\x53\xed\x15\x4b\xca\x75\xdb\x99\x7c\xfd\x7a\xfb\xd5\x07\x07\xfe\x0d\xfb\xed\x95\x41\x6e\x3f\xe4\x97\x85\xe2\x82\x10\x74\xf3\xeb\x74\x60\xae\x32\xd9\x05\x02\xd3\x15\x50\x3c\x74\xd2\x3d\xd4\xdc\x8a\x31\x5d\xd5\x4c\xe9\xfd\xe7\x51\x69\xa4\x55\x76\xe5\x06\x08\x76\x97\xa2\x90\x21\xa5\x86\x4e\x71\x4b\x2a\xec\xc6\x43\xe6\x66\x73\x19\xb5\x2c\xc6\x86\xaf\xd2\xb3\x2d\x04\x5d\x8d\x6b\x96\x34\x66\xa5\xb9\xae\x0e\xb9\x5d\xa5\x67\xab\x74\x4b\x78\x11\x25\x2d\x48\x1d\x25\xe8\x6a\x79\x90\xae\xc0\x78\x9f\xed\x0d\xe5\xa7\x19\x17\x2d\x7b\x84\xad\x01\x78\x99\x71\xb1\x9a\x1b\xb3\x3d\x2c\xee\x0d\x4a\x5e\xa0\x9b\xef\x07\x57\x87\xd9\xa2\x84\x46\xcb\xa8\x09\xe3\xbc\x71\x7e\x4f\xdf\x8f\x2b\x22\xcd\xdb\xe5\x34\xe9\x62\xf7\x39\x22\x72\x8e\xd2\x80\x76\xd1\xa8\xa7\x78\x07\xdd\x74\x59\x67\x0d\x91\xf7\x46\xd0\xdd\x35\x50\xd0\xdd\xfd\xa7\xea\xaf\x00\x00\x00\xff\xff\x7b\x31\x75\x16\xb1\x12\x00\x00")

func assetsHomeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsHomeHtml,
		"assets/home.html",
	)
}

func assetsHomeHtml() (*asset, error) {
	bytes, err := assetsHomeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/home.html", size: 4785, mode: os.FileMode(420), modTime: time.Unix(1621910586, 0)}
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
	"assets/home.html": assetsHomeHtml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"home.html": &bintree{assetsHomeHtml, map[string]*bintree{}},
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
