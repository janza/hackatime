// Code generated by go-bindata.
// sources:
// static/index.html
// static/js.js
// DO NOT EDIT!

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

var _staticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\x4d\x6f\xdc\x20\x10\xbd\xf3\x2b\xa6\x54\xbd\xad\xcd\x56\xe9\xa1\x72\xb0\xa5\x7e\x49\x3d\xb5\x3d\xf4\xd2\x23\x31\x63\x33\x5b\x1b\x2c\x98\x75\xb2\xaa\xf2\xdf\x2b\xf0\x36\xd9\x6a\xa5\x5c\x60\x86\x79\xef\x0d\xf3\x40\xbf\xfa\xfc\xfd\xd3\xcf\x5f\x3f\xbe\x80\xe3\x79\xea\x84\xde\x36\x00\xed\xd0\xd8\x1c\x00\xe8\x19\xd9\x40\xef\x4c\x4c\xc8\xad\x3c\xf2\x50\xbd\x97\xa0\x2e\x8b\xde\xcc\xd8\xca\x95\xf0\x7e\x09\x91\x25\xf4\xc1\x33\x7a\x6e\xe5\x3d\x59\x76\xad\xc5\x95\x7a\xac\x4a\xf2\xcc\x64\xe2\x09\xbb\x8f\x93\xd1\x6a\x0b\x45\xee\xab\xfe\x35\xd6\x77\xc1\x9e\xce\xd0\xc4\xa7\x52\x1f\x42\x9c\xe1\x8f\x00\x18\x82\xe7\x6a\x30\x33\x4d\xa7\x06\xe4\x57\x9c\x56\x64\xea\x0d\x7c\xc3\x23\xca\x1d\x3c\x1d\xec\xe0\x43\x24\x33\xed\x20\x19\x9f\xaa\x84\x91\x86\x5b\xf1\x28\x44\x1e\x73\x07\xb9\x43\x91\x73\x48\xa3\xe3\x06\xde\xee\xf7\x6f\x6e\x05\xc0\x62\xac\x25\x3f\x36\xb0\xcf\xd9\x6c\xe2\x48\xbe\x24\x8f\x42\xa4\x75\x7c\xba\x42\x66\x2c\x0f\xff\x89\x03\x94\x31\x9f\xb5\xae\xb4\x43\x22\xa6\xe0\x1b\x30\x77\x29\x4c\x47\xc6\x22\x5b\x9b\xd8\x03\xe3\x03\xbf\x2c\x9e\x11\x95\xf1\xbd\x0b\xb1\x81\x99\xac\x9d\x2e\xe8\x8b\x61\x57\xe8\x89\x63\xf8\x8d\x0d\xbc\x1e\x86\x6d\xde\xe2\xa2\x3a\xdb\xb8\x59\xba\x8e\x9d\x56\x79\x3d\x57\x53\x1f\x69\x61\x48\xb1\x6f\xa5\x63\x5e\x52\xa3\x94\xbd\x39\xa4\x3a\xc4\x51\xd9\x9b\x7a\x7d\x57\xcf\xe4\xeb\x43\x92\x99\x57\xc0\xdd\x35\xb3\x56\x87\x74\x85\xd1\x6a\x7b\x4b\xad\xb6\xef\xf5\x37\x00\x00\xff\xff\x95\x40\xcf\x52\x76\x02\x00\x00")

func staticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticIndexHtml,
		"static/index.html",
	)
}

func staticIndexHtml() (*asset, error) {
	bytes, err := staticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/index.html", size: 630, mode: os.FileMode(420), modTime: time.Unix(1492878234, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticJsJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\x4d\x6f\xe3\x36\x13\xbe\xeb\x57\xcc\x02\x01\x44\xbd\xd1\xd2\x4a\xf2\xa6\x05\x82\x55\x8b\xc5\xb6\x3d\x14\x6d\xb7\xe8\x2e\xd0\x43\x10\x20\xb4\x38\xb6\xd9\x95\x44\x81\xa4\xed\x04\x8e\xff\x7b\x31\x43\xca\x96\x93\x2c\xd0\x43\x81\x5e\x9c\x68\x3e\x9e\x79\x38\xf3\xf0\x63\xa3\x1c\xe8\x2b\xa8\x61\x6b\x7a\x6d\xb7\x52\x5f\x65\x59\x63\x7b\x1f\xc0\xdb\xe6\x0b\x86\xdf\x9d\x0d\xf6\xe8\x6e\x6d\xa3\x82\xb1\xbd\x1c\xc8\xde\xd8\x16\xea\xba\x86\x7c\x15\xc2\xe0\x6f\x72\xf8\x1e\xf2\xad\xf7\x39\xdc\xd0\xdf\xfc\x04\x08\x6a\xe8\x71\x3b\xe2\xfc\x89\xf3\x4f\x6c\x16\xf7\x67\xbb\x49\xa5\xfd\xcd\x6c\x76\xb6\x7b\x5e\x6c\x65\x7d\xd8\xcf\xb6\xfe\xbe\xc8\x32\x22\xdc\x2a\x1f\x3e\x9b\x0e\xa1\x86\xdd\x9e\x2d\xc1\x06\xd5\x9e\x9a\x16\xc6\x4d\xa3\xb2\xc5\xba\x6f\x08\x0d\x1a\xd5\x36\x6c\x17\x83\xb3\x7f\x61\x13\x4a\x30\x01\xbb\x02\x76\x19\x40\xa4\xac\xa1\x86\x73\xa2\xfb\x83\x0a\x28\xc8\x2b\x83\xe9\xb0\xc8\x00\xcc\x02\xc4\x9b\x91\xc0\x6d\x42\xb8\x8b\xc9\x00\x2f\x1c\x50\x83\x66\xcf\x81\xce\x73\xd7\xfe\x58\xd5\x2c\x16\x64\x84\xb7\x2f\x71\xb2\x54\x9a\x63\xde\xc1\x45\x55\x55\xf0\x3f\xf8\x86\x7e\xae\xc7\xe2\x87\x26\x4c\x4b\x88\x57\xac\x4f\x4f\x50\x15\x70\xce\x05\x99\x41\xf6\x75\xe6\x0e\xc3\xda\xf5\xaf\x60\x67\xfb\x38\x8c\xf5\xa0\x55\xc0\x4f\x41\x05\xea\xb4\xc6\xb9\x5d\xf7\x0d\x0a\xee\x9a\x2f\xa0\xfe\x8e\xd9\x9d\x8e\x0c\x9e\x0f\xec\xd0\x03\xd2\x1a\x67\x4a\x87\x7a\x4d\x38\xbe\x04\x73\x80\x19\xc3\x12\x09\x0a\x96\xe3\xff\x4f\x4f\x90\xdb\xb0\x42\x97\x73\xa0\x9f\xae\x64\x9c\xf9\x64\xe4\x05\x47\xa5\xe5\x79\x6a\x43\x09\xbb\x7d\xc1\x4b\xee\x35\x3a\x91\x01\x7c\x9c\x53\xb0\xfc\x82\x8f\x5e\xd8\x42\x7a\xeb\x82\x10\xaa\x84\x39\x13\x52\xac\xd0\x16\x3f\xd8\x6e\x50\x0e\xc5\xbc\x28\x64\xa7\x06\x61\xc8\x29\x22\xdd\x8d\x6a\xd7\x78\x03\xf6\xd6\xdc\x95\x6c\xe8\x55\x87\x37\x60\xa8\x5e\x51\xc8\x85\x69\x03\xba\x98\xf1\xe6\x8d\x91\x1c\x4e\x1c\x8a\x6c\x5f\xd2\x94\x8b\x2c\x9b\xcd\xe0\x17\xe3\x03\xf6\xb0\xb0\x0e\x3a\xf4\x5e\x2d\xd1\xa7\xcd\xc5\xbd\x82\x1a\x6e\xef\xb2\xb8\x8d\xa4\xd2\xfa\xc7\x0d\xf6\x21\xe6\xa0\x13\x79\x4a\xc9\x4b\x38\xec\x01\x81\x14\x32\x95\x3c\x01\x41\x0d\x3f\x7f\xfa\xf8\x9b\x1c\x94\xf3\x18\x43\xa4\x56\x41\xb1\xee\x79\x28\xc3\xda\xaf\x78\xb2\x64\x9a\x4c\x3e\x4d\x3b\xdb\x47\xbe\x74\x8a\x90\x36\xb6\x46\x87\x15\xd4\x70\x5d\x55\xfc\xbd\x42\xb3\x5c\x85\x89\xc1\x29\x6d\xd6\xc4\xff\x57\x15\x56\xb2\x33\xbd\xe0\x94\x32\x45\x16\x30\x83\x4b\x78\x0b\x17\x55\xc4\x6b\x6c\x6b\x1d\xa9\xec\x4a\x7a\xea\xfc\x47\xa7\x4d\xaf\x5a\xc1\xdf\x2b\xec\xf0\x83\x0a\xb8\xb4\xee\xf1\xb2\x6a\xd2\x59\xe1\x37\xcb\x94\x80\x2d\x36\x41\xe4\x7e\xb3\xcc\x93\x8f\x3c\x7e\xb3\xcc\x00\xa4\x1a\x06\xec\xb5\xc8\xc9\x47\x9f\x21\x38\x91\x07\xa7\x7a\xbf\xb0\xae\xcb\x4b\x88\x1f\x2d\xad\x35\x87\xf3\xb4\x32\x62\x77\x0e\x79\x49\x96\xb4\xb6\x64\x2a\xc6\x1a\x83\xe1\x6d\x71\x45\xa8\x83\x41\xc1\xf0\x83\xd2\xef\xfb\x65\x8b\xa2\x92\xd5\x05\x5b\x78\xf0\x42\x93\x0e\xf4\x51\x05\x2f\x15\x47\xf2\x79\x2e\x3b\x36\x16\x63\x3d\xc5\x1d\x8f\x05\x95\x6b\x62\x41\xbb\x0e\xe8\xfe\xe0\x66\x8b\xd8\x73\x36\x9b\xbe\x3f\x98\xbf\x3e\x81\x8b\x8a\x83\x1b\xeb\x8e\xd1\xdf\x1e\x8e\xe2\x39\xb6\xff\xa4\x1e\xbc\x85\xff\x57\x2f\x8a\x9e\xf8\x8e\x47\xb4\x03\xc1\xba\x63\x7d\xf2\xb1\xa7\x82\x92\x2d\xf6\xcb\xb0\x82\x77\x70\x59\xa4\x7d\x4b\xe7\x06\xb1\x50\xae\x81\x1a\x96\x69\xc6\xef\xdb\x56\xe4\xc4\x25\x8f\x7b\x9c\x35\x2c\xa8\xf9\x0c\x5a\x42\x6a\x33\x83\x72\xef\x46\x1c\xec\x03\xba\xf7\x0c\xa6\x5c\x13\x93\xd9\x26\x12\xd2\xa9\x4c\x0e\x42\x69\x5a\xe5\x3d\x89\x24\x16\xcd\xe0\x80\x74\xc8\xa0\xb9\x9c\x26\xe9\xbc\xe4\x69\x9d\x18\x17\xa6\x6d\xf3\xc4\x90\xd5\x2e\xa6\x3c\x5f\x87\x0e\xf8\x10\x4e\xa1\xa7\xc2\x15\x9a\xa5\xf3\x4c\xbf\x3c\x37\xd9\x60\x1f\x9c\x35\x9a\x62\x92\x68\xa7\x04\x1f\x69\x49\x95\xbc\xba\xc6\x6e\xf4\x50\xad\x28\xd3\xfb\xb3\xdd\x84\xda\x1e\xc4\xd9\x6e\xe4\xca\xfa\x65\xe1\x54\x15\xcc\xe8\x96\xa2\x9f\x42\x06\xfb\x93\x79\x40\x2d\x2e\x8b\x7d\x71\xcf\x6b\x51\xae\x91\xf8\x60\x82\x28\xa4\xc3\xce\x6e\x68\x7f\x44\x73\x2c\x37\x6e\xda\x69\xf7\x78\x21\x86\x74\x32\x4e\x45\xaf\x1d\x3f\x14\xc4\x75\x55\xfd\x2b\x2d\x7e\x51\xff\xbf\x69\xf1\x89\x44\xa7\x4f\x98\xf1\x96\x05\xb1\xe8\x4b\xe0\xa7\x09\x6f\x15\x7e\x0a\x99\x0e\xed\x3a\x1c\xaf\xee\xe3\xa1\x2f\xa5\x54\x6e\xe9\xc7\xf7\x42\xd3\xa2\x72\x9f\x63\xb8\x48\x69\x91\x43\xfa\xa0\xe3\x11\xc3\x18\x21\x78\x91\x8b\xfe\x00\x93\x2a\xf3\x0b\x22\x3d\x06\xe2\xd5\x39\x7d\x07\xb8\x74\x8d\xfd\x1d\x00\x00\xff\xff\xb1\x15\x75\xa3\x69\x0a\x00\x00")

func staticJsJsBytes() ([]byte, error) {
	return bindataRead(
		_staticJsJs,
		"static/js.js",
	)
}

func staticJsJs() (*asset, error) {
	bytes, err := staticJsJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/js.js", size: 2665, mode: os.FileMode(420), modTime: time.Unix(1492898580, 0)}
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
	"static/index.html": staticIndexHtml,
	"static/js.js": staticJsJs,
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
	"static": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{staticIndexHtml, map[string]*bintree{}},
		"js.js": &bintree{staticJsJs, map[string]*bintree{}},
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

