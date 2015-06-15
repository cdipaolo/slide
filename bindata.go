// Code generated by go-bindata.
// sources:
// www/app.js
// www/index.html
// www/styles.css
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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
	name string
	size int64
	mode os.FileMode
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

var _wwwAppJs = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x93\x31\x6f\xdb\x30\x10\x85\x77\xff\x8a\x4b\x16\x52\xb1\x61\xc7\x40\xa6\x0a\x1e\xd2\xd4\x40\x86\x36\x43\xd1\xa1\x40\xd1\x81\x11\xcf\x12\x5b\x8a\x74\xc9\x53\x9b\xa0\xf0\x7f\xef\x91\xb6\x6b\x49\x06\xd2\xa1\x63\x38\x58\x96\xfc\xee\xee\xbd\xef\x64\xb9\xe9\x5c\x45\xc6\x3b\x90\x05\xfc\x9e\x00\x1f\xd1\x45\x84\x48\xc1\x54\x24\xca\xfc\x24\x7f\xfc\x54\x01\xf0\xc9\x44\x8a\xb0\x82\x53\x55\x17\xec\xb1\xf0\xa8\x6a\x88\xb6\xac\x71\xf8\x0b\x3e\x7f\x78\x7f\xcf\x77\x1f\xf1\x47\x87\x91\x64\x51\x4e\xfe\x2a\x93\x6a\xee\xb7\xe8\xa4\xb8\x5f\xdf\xbe\x13\x33\xe0\x56\x33\xd8\x28\x1b\xb1\x28\x87\xb2\x88\x4e\x0f\x8a\x03\x52\x17\xdc\xe1\x47\x52\xd4\x45\xb8\x58\xad\xe0\xe6\xfa\x66\x5f\xb9\xeb\x19\x5f\x5c\x71\x1a\x15\x08\x1e\x9f\xc1\x7a\xff\xdd\xb8\x1a\x36\x3e\xc0\x36\x5d\x8d\xc5\x38\x03\x6a\xd0\xc1\xb7\x6d\x0d\x57\x8b\x53\x56\x8b\x2d\xa7\xd0\xbe\xea\x5a\x74\x34\xaf\x91\xd6\xfc\x88\xbf\xc6\xb7\xcf\x77\x56\xc5\xf8\xa0\x5a\x94\x22\x5a\xa3\x31\x8a\xe2\xcb\xf5\xd7\xde\x50\xb3\x01\xb9\xa7\x25\x85\x69\xeb\xc5\x72\xce\xe3\x44\x31\x46\xc5\x91\x79\x46\x4f\x71\x0a\x9e\x2c\xca\xa4\x31\xac\x58\x96\x07\xf6\x99\x77\x09\x66\x3a\xed\xb7\x3a\xb6\xcb\x5e\xfa\xa6\xab\x80\x8a\xf0\xe0\x3b\x3b\x11\x3d\xb6\xe9\xe4\x12\x06\x4c\xb7\xc4\x2b\x7f\xec\x28\x45\x0a\xd5\x7e\x1d\x23\xed\xe0\x26\xf1\x99\xab\x2d\x2f\x50\xdf\x35\xc6\x6a\x99\x3b\xbd\x54\xd1\x0b\x2b\x60\x0a\xd2\x4c\x97\x05\x5f\xc5\x28\x78\x3a\x95\x77\xd1\x5b\x9c\x5b\x5f\xcb\x4b\xae\x7b\x03\x97\x63\x43\xbb\xc9\x60\xc8\x8e\x0d\xf1\x7b\x7b\xce\x9d\xd7\xfa\x0f\xee\x49\xf1\x2a\xb9\x0f\x83\xff\x17\xf7\x13\x93\xc5\x02\x94\xd6\x80\x21\x30\xc8\x16\x63\x54\x75\x5e\x8b\xf3\xd4\xa4\xbf\x9e\x89\xd0\x2a\xaa\x1a\xd4\x83\x8d\xb0\xfe\x05\x80\xcd\xb2\xcf\x8f\xb5\x03\x02\xa3\xaa\x4f\xf8\x44\x0f\x5e\x33\xd0\x75\x36\x61\xbd\xd2\x79\x72\xcb\x56\xe2\x05\xbf\x0d\xe5\xe4\x0c\xd4\x19\x56\x1e\x72\xd0\xed\x26\x3b\xc9\x35\x7f\x02\x00\x00\xff\xff\x09\xc7\x3d\x75\x2b\x05\x00\x00"

func wwwAppJsBytes() ([]byte, error) {
	return bindataRead(
		_wwwAppJs,
		"www/app.js",
	)
}

func wwwAppJs() (*asset, error) {
	bytes, err := wwwAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/app.js", size: 1323, mode: os.FileMode(420), modTime: time.Unix(1434346084, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _wwwIndexHtml = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x92\x49\x4f\xc3\x30\x10\x85\xef\x48\xfc\x87\xa9\xef\xb1\x55\xa4\x0a\x84\xd2\x5e\x58\xae\x70\x28\x07\x8e\xae\x3d\xc8\x6e\x9d\x05\xcf\x50\xc8\xbf\x27\x8b\x2b\x5c\x50\x05\xe4\x90\xcc\x78\x79\xdf\xbc\xa7\x94\xb3\xdb\x87\x9b\xf5\xf3\xe3\x1d\x38\xae\xc2\xea\xfc\xac\x4c\x5f\xe8\x9f\xd2\xa1\xb6\xa9\x1e\xfb\x0a\x59\x83\x71\x3a\x12\xf2\x52\x3c\xad\xef\x8b\x2b\x91\xef\xb3\xe7\x80\xab\xf1\x5d\xaa\xa9\xc9\x76\x83\xaf\x77\x10\x31\x2c\x05\x71\x17\x90\x1c\x22\x0b\x70\x11\x5f\x0e\x2b\xd2\x10\x09\xe0\xae\xc5\xa5\x60\xfc\x60\x35\xf4\xbf\x4b\x7c\xbb\x90\x24\x1d\x73\x4b\xd7\x4a\x19\x5b\xcb\x2d\x59\x0c\x7e\x1f\x65\x8d\xac\xb6\xaf\x6f\x18\x3b\x49\xc1\x9b\x9d\x9a\xcb\x85\x5c\xa8\xb1\x1e\xe9\x2a\xc7\x91\x89\xbe\xe5\x5c\x7f\xab\xf7\x7a\x5a\x15\x40\xd1\x4c\x94\x01\xd2\x58\x94\x49\xd8\x34\x55\x62\x14\x73\x79\x29\x2f\x64\xe5\x87\x09\xc4\xaa\x54\xd3\xd5\x43\xbc\x2a\xcb\xb7\xdc\x34\xb6\xeb\xeb\x8c\x6e\xfd\x1e\x4c\xd0\x44\xbd\xdb\xe0\x2d\x1e\x45\x31\x9e\x98\x15\x05\xbc\xfb\x10\x60\x83\xa0\xad\x45\x0b\xdc\x80\xaf\x81\x1d\xc2\xd7\xa4\x50\x14\xb9\x29\xd5\xeb\x1e\x83\xfe\x68\xf3\x7f\x61\x9e\x30\x9d\x03\x47\x65\xdd\xb6\xc3\xb9\x13\xf4\x9f\x99\xa5\x9c\xfa\xf0\xa6\x1f\xf5\x33\x00\x00\xff\xff\xf3\xf2\x18\xb3\xc2\x02\x00\x00"

func wwwIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwIndexHtml,
		"www/index.html",
	)
}

func wwwIndexHtml() (*asset, error) {
	bytes, err := wwwIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/index.html", size: 706, mode: os.FileMode(420), modTime: time.Unix(1434345933, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _wwwStylesCss = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8d\xc1\xae\xc2\x20\x10\x45\xf7\x7c\xc5\x24\x2f\x6f\xd9\x97\xf6\x99\xc6\x84\x7e\x0d\x74\x86\x76\x22\x32\x0d\xd0\xa8\x31\xfe\xbb\xa0\x55\x57\x86\x0d\x9c\x7b\xb8\xd7\x0a\x5e\xe0\xaa\x00\xac\x19\x0f\x53\x94\x35\x60\x33\x8a\x97\xa8\xe1\x87\xb0\x9e\x41\xdd\x94\x4a\x9e\x91\xbe\x79\xce\xb9\xa1\x24\x27\xc6\x3c\x6b\xd8\xb7\xed\x72\xae\xef\x99\x78\x9a\xb3\x86\xae\x7b\x91\x45\x12\x67\x96\xa0\xc1\xd8\x24\x7e\xcd\x54\xa9\x95\x88\x14\x9b\x68\x90\xd7\x54\xf4\xff\x2a\x97\xcd\xbf\x91\x42\x2e\xc1\x67\xdb\x93\x2b\x7d\x7d\xfb\x5b\xbf\x65\x59\xde\xf7\xa3\x89\x13\x87\x66\xcb\xfb\x6d\x6e\xa3\x0f\x71\xf7\x84\xb5\x56\xdd\x03\x00\x00\xff\xff\x48\xb6\x3a\x69\xf4\x00\x00\x00"

func wwwStylesCssBytes() ([]byte, error) {
	return bindataRead(
		_wwwStylesCss,
		"www/styles.css",
	)
}

func wwwStylesCss() (*asset, error) {
	bytes, err := wwwStylesCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/styles.css", size: 244, mode: os.FileMode(420), modTime: time.Unix(1434266886, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"www/app.js": wwwAppJs,
	"www/index.html": wwwIndexHtml,
	"www/styles.css": wwwStylesCss,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"www": &bintree{nil, map[string]*bintree{
		"app.js": &bintree{wwwAppJs, map[string]*bintree{
		}},
		"index.html": &bintree{wwwIndexHtml, map[string]*bintree{
		}},
		"styles.css": &bintree{wwwStylesCss, map[string]*bintree{
		}},
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
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
                err = RestoreAssets(dir, path.Join(name, child))
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

