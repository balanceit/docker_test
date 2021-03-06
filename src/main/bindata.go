// Code generated by go-bindata.
// sources:
// db/migrations/1_create_list_table.sql
// db/migrations/2_create_show_table.sql
// db/migrations/3_create_bob_table.sql
// db/migrations/4_create_sue_table.sql
// db/migrations/5_create_jerry_table.sql
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

var _dbMigrations1_create_list_tableSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe2\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x4a\x2e\x4a\x05\xb1\x4a\x12\x93\x72\x52\x15\x72\x32\x8b\x4b\x14\x34\x92\x4b\x8b\x8a\x52\xf3\x4a\x4a\x32\x73\x81\xe2\x40\xa2\xb8\x24\x31\xb7\x40\x21\x2f\xbf\x44\x21\xaf\x34\x27\x47\x21\x25\x35\x2d\xb1\x34\xa7\x44\x01\xaa\x2c\x1e\xae\x44\xd3\x9a\x0b\xc5\x74\x97\xfc\xf2\x3c\x2e\x97\x20\xff\x00\x85\x10\x47\x27\x1f\x57\xb0\xe9\xd6\x5c\x80\x00\x00\x00\xff\xff\xd5\xe1\x36\x00\x81\x00\x00\x00")

func dbMigrations1_create_list_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations1_create_list_tableSql,
		"db/migrations/1_create_list_table.sql",
	)
}

func dbMigrations1_create_list_tableSql() (*asset, error) {
	bytes, err := dbMigrations1_create_list_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/1_create_list_table.sql", size: 129, mode: os.FileMode(420), modTime: time.Unix(1470474717, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations2_create_show_tableSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe2\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x4a\x2e\x4a\x05\xb1\x4a\x12\x93\x72\x52\x15\x8a\x33\xf2\xcb\x15\x34\xb8\x38\x93\x4b\x8b\x8a\x52\xf3\x4a\x4a\x32\x73\x81\x32\x40\xa2\xb8\x24\x31\xb7\x40\x21\x2f\xbf\x44\x21\xaf\x34\x27\x47\x21\x25\x35\x2d\xb1\x34\xa7\x44\x01\xaa\x2c\x1e\xae\x84\x4b\xd3\x9a\x0b\xc5\x02\x97\xfc\xf2\x3c\x2e\x97\x20\xff\x00\x85\x10\x47\x27\x1f\x57\xb0\x05\xd6\x5c\x80\x00\x00\x00\xff\xff\x04\x17\x4d\xe5\x84\x00\x00\x00")

func dbMigrations2_create_show_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations2_create_show_tableSql,
		"db/migrations/2_create_show_table.sql",
	)
}

func dbMigrations2_create_show_tableSql() (*asset, error) {
	bytes, err := dbMigrations2_create_show_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/2_create_show_table.sql", size: 132, mode: os.FileMode(420), modTime: time.Unix(1470130969, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations3_create_bob_tableSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe2\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x4a\x2e\x4a\x05\xb1\x4a\x12\x93\x72\x52\x15\x92\xf2\x93\x14\x34\xb8\x38\x93\x4b\x8b\x8a\x52\xf3\x4a\x4a\x32\x73\x81\x12\x40\xa2\xb8\x24\x31\xb7\x40\x21\x2f\xbf\x44\x21\xaf\x34\x27\x47\x21\x25\x35\x2d\xb1\x34\xa7\x44\x01\xaa\x2c\x1e\xae\x84\x4b\xd3\x9a\x0b\xc5\x7c\x97\xfc\xf2\x3c\x2e\x97\x20\xff\x00\x85\x10\x47\x27\x1f\x57\x90\xf9\xd6\x5c\x80\x00\x00\x00\xff\xff\xd1\x2a\x8c\x6b\x82\x00\x00\x00")

func dbMigrations3_create_bob_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations3_create_bob_tableSql,
		"db/migrations/3_create_bob_table.sql",
	)
}

func dbMigrations3_create_bob_tableSql() (*asset, error) {
	bytes, err := dbMigrations3_create_bob_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/3_create_bob_table.sql", size: 130, mode: os.FileMode(420), modTime: time.Unix(1470130969, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations4_create_sue_tableSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe2\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x4a\x2e\x4a\x05\xb1\x4a\x12\x93\x72\x52\x15\x8a\x4b\x53\x15\x34\xb8\x38\x93\x4b\x8b\x8a\x52\xf3\x4a\x4a\x32\x73\x81\x12\x40\xa2\xb8\x24\x31\xb7\x40\x21\x2f\xbf\x44\x21\xaf\x34\x27\x47\x21\x25\x35\x2d\xb1\x34\xa7\x44\x01\xaa\x2c\x1e\xae\x84\x4b\xd3\x9a\x0b\xc5\x7c\x97\xfc\xf2\x3c\x2e\x97\x20\xff\x00\x85\x10\x47\x27\x1f\x57\x90\xf9\xd6\x5c\x80\x00\x00\x00\xff\xff\x16\xa1\x15\x61\x82\x00\x00\x00")

func dbMigrations4_create_sue_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations4_create_sue_tableSql,
		"db/migrations/4_create_sue_table.sql",
	)
}

func dbMigrations4_create_sue_tableSql() (*asset, error) {
	bytes, err := dbMigrations4_create_sue_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/4_create_sue_table.sql", size: 130, mode: os.FileMode(420), modTime: time.Unix(1470130969, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations5_create_jerry_tableSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe2\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x4a\x2e\x4a\x05\xb1\x4a\x12\x93\x72\x52\x15\xb2\x52\x8b\x8a\x2a\x15\x34\xb8\x38\x93\x4b\x8b\x8a\x52\xf3\x4a\x4a\x32\x73\x81\x52\x40\xa2\xb8\x24\x31\xb7\x40\x21\x2f\xbf\x44\x21\xaf\x34\x27\x47\x21\x25\x35\x2d\xb1\x34\xa7\x44\x01\xaa\x2c\x1e\xae\x84\x4b\xd3\x9a\x0b\xc5\x06\x97\xfc\xf2\x3c\x2e\x97\x20\xff\x00\x85\x10\x47\x27\x1f\x57\x88\x0d\xd6\x5c\x80\x00\x00\x00\xff\xff\xc1\xf2\xea\xf9\x86\x00\x00\x00")

func dbMigrations5_create_jerry_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations5_create_jerry_tableSql,
		"db/migrations/5_create_jerry_table.sql",
	)
}

func dbMigrations5_create_jerry_tableSql() (*asset, error) {
	bytes, err := dbMigrations5_create_jerry_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/5_create_jerry_table.sql", size: 134, mode: os.FileMode(420), modTime: time.Unix(1470130969, 0)}
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
	"db/migrations/1_create_list_table.sql": dbMigrations1_create_list_tableSql,
	"db/migrations/2_create_show_table.sql": dbMigrations2_create_show_tableSql,
	"db/migrations/3_create_bob_table.sql": dbMigrations3_create_bob_tableSql,
	"db/migrations/4_create_sue_table.sql": dbMigrations4_create_sue_tableSql,
	"db/migrations/5_create_jerry_table.sql": dbMigrations5_create_jerry_tableSql,
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
	"db": &bintree{nil, map[string]*bintree{
		"migrations": &bintree{nil, map[string]*bintree{
			"1_create_list_table.sql": &bintree{dbMigrations1_create_list_tableSql, map[string]*bintree{}},
			"2_create_show_table.sql": &bintree{dbMigrations2_create_show_tableSql, map[string]*bintree{}},
			"3_create_bob_table.sql": &bintree{dbMigrations3_create_bob_tableSql, map[string]*bintree{}},
			"4_create_sue_table.sql": &bintree{dbMigrations4_create_sue_tableSql, map[string]*bintree{}},
			"5_create_jerry_table.sql": &bintree{dbMigrations5_create_jerry_tableSql, map[string]*bintree{}},
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

