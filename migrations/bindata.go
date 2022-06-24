// Code generated by go-bindata.
// sources:
// 0001_init.down.sql
// 0001_init.up.sql
// bindata.go
// generate.go
// DO NOT EDIT!

package migrations

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

var __0001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x72\x75\xf7\xf4\xb3\xe6\xe5\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x08\x0d\x76\x0d\x0a\xc6\x25\x19\x12\xe4\xe8\x17\xec\xe8\x1c\xe2\xe9\xef\x07\x52\xe3\xec\xef\xeb\xeb\x19\x62\x0d\x08\x00\x00\xff\xff\x1b\x95\x0a\xbe\x50\x00\x00\x00")

func _0001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_initDownSql,
		"0001_init.down.sql",
	)
}

func _0001_initDownSql() (*asset, error) {
	bytes, err := _0001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_init.down.sql", size: 80, mode: os.FileMode(438), modTime: time.Unix(1655939592, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcf\x4d\x4b\xc3\x40\x10\xc6\xf1\xb3\x0b\xfb\x1d\x9e\x63\x0a\x45\x7a\xef\x69\x93\x4e\xca\xe0\x66\x22\xb3\x13\xb0\xc7\x60\x23\x14\xfa\x22\xb6\x1e\xfc\xf6\xb2\x2b\x88\x04\xec\xf5\x77\x78\xfe\x33\x35\x6d\x59\xd6\xde\x79\xd7\x28\x05\x23\x58\xa8\x23\x81\x5b\x48\x6f\xa0\x17\x4e\x96\x30\x24\xd2\x84\xca\x3b\xde\x23\x91\x72\x88\x78\x56\xee\x82\xee\xf0\x44\xbb\xa5\x77\xf5\x78\x1c\xcf\xaf\x13\xde\x8e\x97\xf1\x86\x0d\xb5\x61\x88\x86\xd5\xe3\xaa\xac\xc8\x10\xa3\x77\x8b\x52\xb9\xdb\x31\x0d\x92\x42\x63\xdc\x4b\xc9\x1d\xfe\xc9\xd9\xd7\xfb\x04\x16\xfb\x5d\x5f\x7a\x37\x5c\xa7\x0f\xde\xcf\x35\x9c\x2e\x9f\xe7\xdb\x5c\xed\x70\x9a\xb0\xc9\x57\xfc\xc1\x07\xb4\xbd\x12\x6f\x25\x47\xaa\x9f\xb9\x45\x66\xa5\x96\x94\xa4\xa1\x84\xac\xd7\xaa\x78\x7e\xa7\xe9\xbb\x8e\x6d\xfd\x1d\x00\x00\xff\xff\xa7\x22\xe9\xc5\x44\x01\x00\x00")

func _0001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_initUpSql,
		"0001_init.up.sql",
	)
}

func _0001_initUpSql() (*asset, error) {
	bytes, err := _0001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_init.up.sql", size: 324, mode: os.FileMode(438), modTime: time.Unix(1655939586, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(438), modTime: time.Unix(1655939885, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _generateGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x8e\xbd\x6a\xc4\x40\x0c\x84\x7b\x3f\x85\xb8\xe2\x3a\x9f\x21\x3f\x8d\xc1\xa4\xc8\x23\x98\x90\x66\x1b\x79\x2d\x74\x9b\x5b\xaf\x36\x92\x9c\x9f\xb7\x0f\x0e\x17\x70\x20\x6d\xea\x99\x6f\xe6\xab\x18\x2f\xc8\x04\x4b\x62\x45\x4f\x52\xac\x69\xba\x8e\xa5\x67\x2a\xa4\xe8\x04\x2c\xed\x94\xca\x8c\x8e\xd0\xd6\x0b\xef\x9a\xd0\x0a\x5c\xa3\x13\x0b\x9c\x9a\x0d\xbd\xe6\x04\xad\xc9\xaa\x91\xe0\xb1\x0f\x4f\x46\x6a\x61\x24\x7d\x09\x2c\xc1\x34\x86\x31\x2d\x35\x7f\x3e\xd3\x34\x92\xbe\x91\x86\xfd\xe8\xb6\x37\xa1\x11\x54\x31\x67\x25\xeb\xbb\x2e\x4b\xc4\x7c\x16\xf3\xfe\xfe\xee\xf6\xa6\x73\xc5\x62\x18\x37\x60\x74\x51\x7a\x30\xcb\x8b\xcc\x34\xcc\xc9\x70\xca\x74\x5c\x8d\x74\xf8\xe1\x8f\x15\xcd\xde\x45\xe7\x81\xf9\x00\x6b\xfd\x65\x59\xd1\xcf\xf0\xe7\xfd\xe1\xdf\xfe\xf7\x02\x51\xe9\xdb\x83\x3e\x1c\xec\x35\x43\x3b\x27\xdd\xfb\xa4\x92\xfc\x2b\x00\x00\xff\xff\xb4\x07\xf7\x87\xa5\x01\x00\x00")

func generateGoBytes() ([]byte, error) {
	return bindataRead(
		_generateGo,
		"generate.go",
	)
}

func generateGo() (*asset, error) {
	bytes, err := generateGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "generate.go", size: 421, mode: os.FileMode(438), modTime: time.Unix(1655939794, 0)}
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
	"0001_init.down.sql": _0001_initDownSql,
	"0001_init.up.sql": _0001_initUpSql,
	"bindata.go": bindataGo,
	"generate.go": generateGo,
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
	"0001_init.down.sql": &bintree{_0001_initDownSql, map[string]*bintree{}},
	"0001_init.up.sql": &bintree{_0001_initUpSql, map[string]*bintree{}},
	"bindata.go": &bintree{bindataGo, map[string]*bintree{}},
	"generate.go": &bintree{generateGo, map[string]*bintree{}},
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
