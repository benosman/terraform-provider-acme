// Code generated by go-bindata.
// sources:
// templates/acme-provider-sidebar-template.tmpl
// templates/dns-provider-doc-template.tmpl
// templates/dns-provider-go-template.tmpl
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

var _templatesAcmeProviderSidebarTemplateTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\xc1\x6e\xdb\x30\x0c\xbd\xe7\x2b\x38\x0f\x01\xb6\x83\xe2\x7b\x21\x1b\x18\xba\x1d\x37\x0c\xdb\x07\x18\xaa\x45\x27\x04\x64\xca\xa0\x64\xaf\x41\x90\x7f\x1f\x92\xda\x8e\xdd\xac\xab\xd3\x1e\x25\x52\x8f\xef\x3d\x92\xd2\x6b\xf8\x23\xa6\x29\x9c\xd9\xfb\x36\xc2\x1d\x31\xa3\x80\xf5\xb0\xce\x57\x00\x7a\x0d\xa5\xe7\x88\x1c\x8b\xca\x0b\xdc\x05\xb2\xf8\x60\x2e\x71\x00\xfd\x41\x29\x30\x6d\xf4\x6a\x8b\x8c\x62\x22\x5a\xa8\xc4\xd7\x60\xca\x1a\x55\x23\xbe\x23\x8b\xa2\xfa\x87\x2a\x62\xdd\x38\x13\x71\x13\xeb\xc6\x81\x52\x3d\x88\xa5\x0e\x4a\x67\x42\xc8\x12\xeb\xcb\x30\xa4\xc3\x8e\xac\x45\x56\x8d\x10\x47\x30\x55\x45\x8f\x2a\xfa\x26\x01\xf1\x0e\xb3\xa4\xf4\x75\xe3\xb0\x46\x8e\x46\xf6\xc9\x13\x14\x80\x6e\xdd\x80\xc5\xa6\x83\x11\x8f\x4d\x37\xe6\x00\x68\x47\x7a\x9d\x41\x5f\xa8\x28\x5b\x11\xe4\xf8\xe9\xa9\xfc\xce\xd7\x98\x7c\x86\x75\x7e\xc9\x07\xd0\x06\x76\x82\x55\x96\xa4\xa7\x9c\x74\x90\x16\x52\x62\x8b\x8f\x9b\x5d\xac\x5d\x92\x7f\x71\x0e\x7e\x0e\x11\x9d\x9a\x49\xc1\xd4\x51\xbe\x5a\x56\xff\xec\xdd\x19\x76\x39\x8b\xd3\x9b\x39\x95\xfb\xef\xdf\x46\x2e\xef\xa3\x22\x18\x7c\x2b\xe5\x7f\x3c\xf9\x98\xe4\xbf\xfa\xa4\xb9\xec\xab\x7e\xb0\xe9\x54\x47\x81\x1e\x1c\x26\xd3\xb4\x5b\x78\x28\xc1\x2d\x85\x28\x26\x92\xe7\x2b\x52\xaf\xdb\x24\xe9\x14\xa0\xf7\xeb\x14\x29\xa6\xf7\xcf\x74\x0c\xbe\xbd\x91\x72\x89\x12\xa9\xa2\xd2\xc4\x6b\x1b\x97\x30\x9e\xbc\x9f\x12\x9e\x5c\xbf\xca\x57\xa7\xad\x7b\x3e\x05\xcb\x95\x58\x0e\xe3\x3e\x87\x1b\xe7\xd2\x72\x28\xfe\xbd\x30\x5f\x7f\xfc\x7e\x61\x61\x6e\x98\x9c\xc3\x41\x81\x18\xde\x22\x6c\x8e\xc7\x1b\x1b\x34\x93\xa5\x0e\x87\xcd\xbd\xb7\x78\x3c\xbe\xa1\x47\x73\x91\x23\x52\x2f\x74\x3c\x2f\x18\xab\x93\x1e\x64\x3b\xd3\xf2\x72\xf3\x2e\x11\x9d\x5a\xea\xfa\x7f\x1b\xd9\x0e\x7f\x78\x06\x7b\x42\x77\x3e\x5e\x02\xab\xbf\x01\x00\x00\xff\xff\xa7\xdf\x84\x96\xfa\x05\x00\x00")

func templatesAcmeProviderSidebarTemplateTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAcmeProviderSidebarTemplateTmpl,
		"templates/acme-provider-sidebar-template.tmpl",
	)
}

func templatesAcmeProviderSidebarTemplateTmpl() (*asset, error) {
	bytes, err := templatesAcmeProviderSidebarTemplateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/acme-provider-sidebar-template.tmpl", size: 1530, mode: os.FileMode(420), modTime: time.Unix(1565474879, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDnsProviderDocTemplateTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x41\x8f\xdb\x36\x13\xbd\xcf\xaf\x18\x78\xf7\x64\x58\xf2\xdd\x40\x0e\x86\xbf\x0d\x10\x20\x5f\x50\x24\x69\x2f\x86\x61\x8d\xa9\x91\x45\x2c\x4d\x0a\x24\xa5\xcd\x42\xd5\x7f\x2f\x46\xb6\x25\xcb\x6d\x77\xd3\x02\x3d\x59\x32\x87\xc3\xf7\xde\xbc\x47\x25\x49\x02\x86\x5e\x5d\x1d\x57\x38\x23\x75\xe2\x19\x54\x74\xe4\x7d\xd4\xd1\xf0\x0a\x67\xeb\xcd\xff\x9f\x56\xd8\xb6\xe9\x17\x3a\x71\xd7\xe1\xff\xbe\x7c\xc3\x4d\x49\xc6\xb0\x3d\x32\xfe\xe2\x5d\xa3\x73\xf6\x33\x08\x3a\xe7\x03\xf9\xbd\xaa\xbd\x67\x2b\xcd\x72\xa7\x42\x22\x1d\x93\xdc\x86\xa4\xba\x54\x86\xa4\x6d\xd3\x8d\xcb\xb9\xeb\x66\x90\x73\x50\x5e\x57\x51\x3b\xbb\xc2\xdf\x13\xc0\x6b\xc3\x80\x84\x9e\x83\xab\xbd\x62\x8c\x0e\x4f\x64\xe9\xc8\xa8\xd8\x47\x5d\x68\x45\x91\x03\x3a\x8b\x64\x51\xf0\xe1\x66\x9d\x82\x10\x81\x87\x77\x91\x02\x7c\x2f\x19\xb3\x01\x44\xd6\xd7\xa9\xa1\xee\x8a\x13\x15\x59\x3c\x30\xd6\x81\x73\x41\x50\xb1\x2f\x9c\x3f\x4d\xab\x03\x16\xce\x43\x2c\x19\xb7\x99\x30\xdd\xdf\x00\xcc\x76\xdb\x2b\x83\xb3\x0a\x37\x6b\xbb\x91\xdc\x8b\x8e\x25\x6c\x07\xd4\xbb\xed\x15\x40\x12\xd8\x37\x5a\x71\x22\xe3\xd8\xa5\x00\x6f\xb4\x5b\xe1\x52\xd4\x5e\x0e\x22\x2f\xa5\x64\xe9\x97\x37\x45\x69\x19\x4f\x06\xfe\xa6\x7d\x3f\xe1\x5f\xbf\x7e\xee\x3a\x80\x8f\xce\xa3\x72\xa7\xca\x70\x64\xd4\x56\x58\x93\x4c\x48\x04\x2f\xdd\x8b\x88\x51\x07\xc6\x58\xea\x30\xaa\x25\x34\x50\x84\x98\xe8\x20\x32\xc0\x15\xf6\x02\x03\x33\x6e\x4b\xf6\xfc\x86\x34\xbd\x59\x46\x81\xdf\x26\x7e\x5f\xfc\xd3\x3a\x3c\xd4\x41\xdb\xe3\xdd\x76\x80\x87\x07\x7c\xfa\x41\x42\x1d\x20\xcb\xb2\x52\x99\x01\xfd\x39\x1c\xb7\x03\x9e\xe1\x6c\xf2\xd6\x02\x62\x9a\xa6\x00\x88\xb9\x0d\xfb\xd1\x52\xb2\x80\xa3\x54\x1f\x70\x76\x13\x01\xc4\x0e\x3a\x39\xac\x3f\x7d\xed\x8f\xf5\x89\x6d\xc4\xaf\x5c\xb0\x67\xab\xf8\xec\xd7\xc2\x19\xe3\x5e\xb4\x3d\x22\x5d\x2a\xc2\xd5\xa0\xac\x63\xc9\x1e\x2b\x0a\xe2\x54\x0a\xc8\xb6\xd1\xde\xd9\xbe\x4d\x43\x5e\xd3\xc1\x70\x58\xa0\xf3\x90\x6b\xcf\x2a\x9a\x57\x8c\xa5\x77\xf5\xf1\x32\x30\xe5\x6c\xa1\x8f\x19\x1e\x8c\x53\xcf\xa8\xad\xfc\x0b\xdb\x6c\xc2\xe1\x2d\x33\x4f\x55\x4c\xc8\x1f\x77\x03\xca\xb1\xdd\xbf\x8b\x47\x8a\xe2\xc6\x93\xf3\x8c\x39\x47\xd2\x26\xf4\x2e\x82\xff\xd0\x45\x3d\x81\x9f\x37\xd2\x44\x26\x80\x4f\x16\x29\xcf\xb5\xe4\x65\x71\x37\x2b\x32\xc1\xc9\xc0\x42\x74\x9e\x73\x51\x86\xd0\x38\x45\x06\x0b\x6d\x78\x31\x46\xa8\xa2\x58\x42\xa8\xab\xca\x68\xce\xf1\xf0\x8a\xfd\xf3\xab\x4c\x5f\x96\x07\x6d\xc7\xcc\xed\x3f\x7e\xfa\xfc\x94\x61\xa8\x8b\x42\xff\x48\xf1\xdb\xa8\xd0\x9f\xb8\xca\x59\x42\x31\xe1\xb3\xd1\x77\x72\x85\x9d\x15\xbe\xc9\xba\xe8\xf5\xfe\xd6\x7f\x1a\xb7\xab\x1b\xfb\x4e\x21\x29\x9c\x1f\x3e\x0b\xc9\xa0\x15\xb4\xad\x27\x89\xcd\xe3\xf3\x02\x1f\x1b\x5c\x7d\xc0\x74\xd3\x5b\xb4\xf6\x67\x6c\x1b\xcf\x39\xdb\xa8\xc9\x84\xae\x83\xb9\x5c\xe6\x8f\xcf\x72\x93\x27\xd8\xb6\x8f\x4d\xd7\xa5\xd0\xb6\x09\xb2\xcd\xe5\x3a\xbb\x8b\xcf\x65\x38\x64\xd0\x55\x97\x87\x21\x24\x48\x9e\x91\x1a\xd2\x46\x5e\x57\xef\x23\x59\x0f\xdd\xde\x07\xd2\xb6\xba\xc0\xf4\xc9\x36\xbf\x91\x5f\x1b\x4d\x81\x05\xfd\x14\xdd\x14\xc9\x7c\xfe\x9d\xbd\x27\x19\x4a\x12\x2a\x56\xa2\xe8\x7c\x8e\x74\xde\xdc\xcf\xad\xf7\xc3\xc1\x35\x0c\xea\x16\x18\x36\x64\x6a\x0e\x2b\xf8\x2b\x0a\xf7\x10\x26\xb8\xfb\xe6\x7d\xeb\xec\x4c\x21\x9b\x72\xb8\xfe\x4e\xa9\xff\x11\x00\x00\xff\xff\x48\xd4\x53\x92\x45\x08\x00\x00")

func templatesDnsProviderDocTemplateTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDnsProviderDocTemplateTmpl,
		"templates/dns-provider-doc-template.tmpl",
	)
}

func templatesDnsProviderDocTemplateTmpl() (*asset, error) {
	bytes, err := templatesDnsProviderDocTemplateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dns-provider-doc-template.tmpl", size: 2117, mode: os.FileMode(420), modTime: time.Unix(1565468560, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDnsProviderGoTemplateTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x4f\x8b\xdb\x30\x10\xc5\xef\xfa\x14\xd3\x10\xa8\x03\x59\xf9\xbe\x90\x43\xd8\x64\x7b\x6a\x08\x2c\xe4\x52\x7a\x98\xb5\xc6\x8e\xb0\x3c\x32\x92\xec\x12\x8c\xbe\x7b\x91\x13\xbb\xe9\x9f\x94\x3d\xd9\x33\xd2\x7b\xf3\x7b\xf6\xe4\x39\x6c\xbb\x60\x9f\x2a\x62\x72\x18\x48\x41\xa9\x0d\x49\xd8\x59\x60\x1b\x80\x94\x0e\x52\xb4\x58\xd4\x58\x11\x60\xd1\x90\x10\xba\x69\xad\x0b\x90\x89\xc5\x30\x2c\xe5\xb1\xae\x8e\x18\xce\x31\xe6\xc5\x19\x8d\x21\xae\x68\x21\x86\xe1\x09\x1c\x72\x45\x20\x8f\xce\xf6\x5a\x91\xf3\x31\xfe\x29\x68\xa7\xa3\x5c\xb1\xcf\x87\x41\x7e\xb1\xc7\xba\x8a\xf1\xaa\x27\x56\x31\x8a\x95\x10\x79\x0e\x8a\xfd\xe4\xf3\x8a\x45\xb0\xee\xf2\xda\x71\x01\xda\x03\x42\xd9\x71\x11\xb4\x65\x08\x67\x0c\x50\xa0\x31\xa9\x3b\x79\x7f\xf6\x49\x5f\x58\xf6\xc1\x75\x49\x09\xc8\x0a\x1c\x85\xce\xb1\x87\x70\xa6\xf9\x26\x68\x0e\xe4\x4a\x2c\x48\x8a\x70\x69\xe9\xd1\xd0\x34\x2f\x5b\x41\x36\xc7\x9d\x23\xae\x81\x9c\xb3\xee\x11\xf2\x0d\xf7\x56\x94\x09\xc5\x18\xb0\xe5\x48\xd1\xa3\xd1\x0a\x76\x87\xb7\x99\x67\x04\xf7\x5d\x9b\x3e\x36\x29\x78\xbf\xc0\xf6\xe5\xeb\x7e\x3e\x96\xa2\x47\xf7\xaf\x29\x1b\x68\xb0\xfd\xe6\x83\xd3\x5c\x7d\x7f\x90\x61\xf8\xcf\x1f\x92\x2f\x56\x51\x8c\x8b\xe7\x0f\x24\xbd\x19\xe9\x12\xe4\x9e\xfb\x13\xba\xad\xd1\xe8\x29\x39\x35\xd8\xee\xb9\xd7\xce\x72\x43\x1c\x4e\xe8\x34\xbe\x1b\x3a\xa1\xe9\xc8\x67\x77\x84\xd7\xc7\x3d\xd0\xb2\x5e\xc3\xb2\x87\xe7\xcd\xdf\xa6\x69\x81\xea\x91\x2d\xbd\xf5\x31\x2e\xd6\x77\xbb\x12\x57\x77\x45\x3b\x32\x26\x97\x5f\x8b\x25\x0f\xf4\x63\x77\x78\x9b\x52\x64\x2b\xa1\xcb\xf1\xd6\xa7\x0d\xb0\x36\x30\x88\xeb\x62\xa4\x62\x94\x8b\x28\xa6\x56\xbb\x4e\x5d\x11\x7f\x1b\x28\x7e\x06\x00\x00\xff\xff\x8a\x29\xdb\x58\x3f\x03\x00\x00")

func templatesDnsProviderGoTemplateTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDnsProviderGoTemplateTmpl,
		"templates/dns-provider-go-template.tmpl",
	)
}

func templatesDnsProviderGoTemplateTmpl() (*asset, error) {
	bytes, err := templatesDnsProviderGoTemplateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dns-provider-go-template.tmpl", size: 831, mode: os.FileMode(420), modTime: time.Unix(1565476137, 0)}
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
	"templates/acme-provider-sidebar-template.tmpl": templatesAcmeProviderSidebarTemplateTmpl,
	"templates/dns-provider-doc-template.tmpl": templatesDnsProviderDocTemplateTmpl,
	"templates/dns-provider-go-template.tmpl": templatesDnsProviderGoTemplateTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"acme-provider-sidebar-template.tmpl": &bintree{templatesAcmeProviderSidebarTemplateTmpl, map[string]*bintree{}},
		"dns-provider-doc-template.tmpl": &bintree{templatesDnsProviderDocTemplateTmpl, map[string]*bintree{}},
		"dns-provider-go-template.tmpl": &bintree{templatesDnsProviderGoTemplateTmpl, map[string]*bintree{}},
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
