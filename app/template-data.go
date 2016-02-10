// Code generated by go-bindata.
// sources:
// templates/base.html
// templates/config-base.html
// templates/login.html
// templates/register.html
// templates/settings.html
// templates/shitbucket-import.html
// templates/tag-index.html
// templates/url-edit.html
// templates/url-index.html
// templates/url-new.html
// templates/url-partial.html
// templates/url-view.html
// DO NOT EDIT!

package app

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

var _templatesBaseHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x58\xcd\x6e\x1b\x37\x10\xbe\xe7\x29\x18\x3a\x80\x0f\xf5\x6a\xe3\x34\xa7\x40\x52\xe0\xfc\x38\x0d\x9a\xb6\x40\xe2\x9c\x0b\x6a\x49\x49\xb4\xb8\xcb\x0d\xc9\x95\x2d\x18\x7e\x8e\x3e\x43\x9f\xa0\xf7\xf6\xc5\x3a\xfc\xdb\xe5\xfe\x28\x31\x52\xf4\x60\x2f\x7f\x86\x33\xc3\xe1\xcc\xc7\x8f\xba\xbb\x43\x94\xad\x79\xc5\x10\x6e\x94\xc8\x6a\xa2\x0c\x27\x02\xa3\xfb\xfb\x47\x30\xf5\x04\xc6\xd0\x8b\x05\x9a\xd9\xfe\x9c\xf2\x3d\x2a\x04\xd1\x7a\x81\x95\xbc\xc1\xcb\x47\x08\xa5\x63\x85\x14\x99\xd8\x64\xe7\xcf\xdc\x4c\x7f\xce\xea\x2e\x64\x65\x08\x58\x52\x30\xef\x04\x40\x64\xfb\x7c\x19\x9a\xd0\x21\x68\xab\xd8\x7a\x81\x83\xe1\xd9\xe7\x8f\x1f\xc0\x2e\x46\x86\xa8\x0d\x33\x0b\xfc\xfb\x4a\x90\x6a\x87\x97\x71\xfe\x8a\x1b\xc1\x40\x62\x9e\x93\xa8\x65\x9e\x5b\x8d\xb1\x53\x47\xf3\x86\xdd\x9a\xac\x6c\x0c\xa3\x78\x39\xd7\x25\x11\x62\x99\x5a\x53\x6c\xcf\x94\x0e\x21\xd8\x73\x76\x83\x11\xe6\x14\x7b\x2b\xef\xdf\x58\x27\x5a\xa3\x97\x52\x95\xc4\x80\xa6\xd7\x8a\x11\xf8\x5c\x98\xe0\x01\xca\x50\xdf\xf1\x79\x1e\x2c\xe5\x75\xeb\x12\x48\xf0\xb5\x17\xfa\x89\xe8\x2b\xb2\xd1\x36\xb2\xc1\x5d\x1e\xdd\x5d\x13\xb4\x26\x99\x21\x1b\xf0\x36\xe7\x4b\xbb\x4a\x91\x6a\xc3\xd0\x13\x18\xb3\xe7\xe1\x14\xbc\x63\x26\x28\x98\xdc\x0b\x88\xf6\xf6\x02\xfd\x74\x2f\xb6\xfb\x2b\x29\x63\xfc\xac\x0d\x56\xd1\xce\x9b\xae\x1f\xdd\x1b\x1c\x27\x29\x0c\x97\x95\xc6\xc9\x01\xea\x9a\x54\x51\x64\xcb\x29\x65\x55\x32\xeb\x0e\x38\x4c\xae\x4c\x85\x8f\x84\x9f\x51\x6e\x26\xc2\xff\x16\x86\x93\x73\x1e\xa9\x43\xee\x88\xa9\x8d\x92\x8a\xaa\x4f\x30\xa2\xc4\x90\x6c\xda\x10\x65\x82\x19\x36\x36\xe5\xd7\x18\xb9\xd9\x08\xb6\xc0\xa5\xa4\xb6\x1c\xfc\x58\xc8\xc3\x13\xc8\xe4\x35\x57\x65\x54\xb1\x7c\xe3\xbe\x3d\xf7\xe0\xf0\x21\x18\x5d\x5a\x42\xf0\x42\x55\xc4\x66\x68\x84\x4f\x12\xee\xa4\x24\x83\xf5\x41\xf1\xb9\x51\xc8\x10\x0a\xde\x73\x6a\x0b\xaf\xe7\x0e\xd4\xcb\x8a\x57\x94\xdd\x2e\x70\x76\x8e\x91\x92\x76\x1f\x14\xaa\x5a\x6e\x46\x35\xeb\x54\x65\x7e\x12\xf9\x8e\x2e\x27\xea\xd7\x4f\xd9\x0a\x66\x95\x69\x0f\x75\x2c\xb1\x65\xe0\x95\x3a\x96\x13\xe9\x19\xf9\xa0\xf1\x6a\x83\xa0\x5a\x8e\x46\x6b\xca\xc6\x4a\xd2\x43\x6a\xa1\x4e\x93\xe2\x42\x31\x74\x90\x0d\xd2\x4d\x68\xdc\x90\xca\x20\x23\x91\x0f\x0e\x32\x5b\xae\xad\xc5\x97\xc9\x49\xd5\x0f\x34\xbc\x96\xd2\xf4\x37\xb7\x6a\x8c\x91\x90\x7b\x87\x1a\x42\xec\x3b\x38\x4d\x4a\xf8\x83\x53\x59\x93\x46\x98\x90\x43\x94\xeb\x92\xb7\x2a\xf1\xf2\x35\xa9\x0a\x26\xe6\xb9\x5f\xdc\x03\xc3\xa1\x1e\x17\x37\xd7\x94\xbb\x89\x9c\x7b\x48\x8e\x41\x8a\x0d\x32\x6c\x0d\x46\xb6\xd8\x97\x7d\x07\x33\x3b\x76\x38\x43\x4f\xf6\x44\x34\x4c\x3b\xfc\xbf\xb4\x62\x4c\x0f\xe5\x4a\xa6\x35\x81\x86\x85\xa4\x20\x1d\x00\x24\x8d\x1e\x11\x4c\x19\xe4\xfe\x67\x16\x7a\x40\xbb\xab\xe9\x47\xfd\x28\xc6\x6b\x44\x48\xcd\x86\xd1\x72\x8b\xf1\xf2\x9f\x3f\x86\x91\xb2\xfa\xa2\x17\xd1\x74\xdc\x7d\x0a\x69\x5d\x7b\xba\xd2\x8c\xbd\x4b\x6c\x1c\xd0\x2f\xbc\xe2\x80\xdb\x5c\x1b\x5e\xa0\x57\x52\xee\x4a\xa2\x76\x90\xa7\xd3\xeb\x56\x44\x33\x5f\xa0\x8f\xa9\x2c\x6c\x22\xa0\xad\x29\x05\xc4\xdc\x7f\xec\x1d\x07\x35\x11\xce\xc5\x59\x59\x7e\xfa\x7c\xf9\xd1\x3a\x64\x58\x59\x0b\xb8\x43\x5a\xeb\x33\x07\xc6\x5e\xc8\x2f\x28\x99\x21\xa8\x02\x98\x5e\x60\x8b\xe5\xb5\x84\x28\xa0\x50\x86\x0b\x7c\xc3\xa9\xd9\x2e\x28\xdb\xf3\x82\x65\xae\x73\x86\xc0\x7b\x7b\x83\x67\xba\x80\x98\x2d\xce\x63\x2d\x0b\x5e\xed\x3c\x2c\x9e\x6e\x8d\xa9\x5f\xe4\xf9\x1a\xb4\xe8\xd9\x46\x4a\xc0\x39\x52\x73\x3d\x2b\x64\x99\x17\x5a\xbf\x5c\x93\x92\x8b\xc3\xe2\xb7\x9a\x55\x3f\x7c\x22\x95\x7e\xf1\xfc\xe9\xd3\xb3\x1f\xfd\x1f\x37\x10\x99\xe2\xec\x79\x6c\x9d\x02\xa6\x8a\xc5\xa9\x36\x07\xc1\x20\x41\x98\x39\xf5\xc5\x70\x6a\x4b\xdd\x6a\x3b\x4d\xed\xfb\x3a\x89\x53\xd8\xad\xc5\xdd\xda\x88\xdb\xb9\x36\x04\x82\x6f\x65\xf2\x15\xd4\x9c\x36\x8a\xd4\xb3\x92\x57\x33\xbb\xea\x3f\x2a\x24\x75\x3d\x52\xf3\x80\x65\x36\x5a\x19\xb9\x61\x5a\x96\xac\xe7\x0a\x30\x8e\x70\xbe\x73\x0b\x4c\x41\xed\xe3\x2c\x43\x97\xfc\x96\x51\x38\xbc\xfd\x8a\x28\x94\x65\x09\xa8\x5a\xd0\x36\xb2\x6e\xa1\x22\xc8\xf8\x4f\x0b\x18\x53\x40\x94\x72\xa8\x16\x2a\x92\xf9\xa0\x62\x04\xc3\xdf\xc0\xaa\xb0\xcc\x5f\x7a\x83\x1b\x10\x68\x9d\x20\x75\x5b\x92\xf1\x12\x9c\x85\x35\xed\x74\x6a\x6c\x80\xfb\x1c\xdc\xce\x40\xd8\xd2\x99\x14\xe8\xff\x3f\xd1\x31\xa0\xa6\x90\x1a\x5c\x5f\x01\x8a\xd1\x63\x3c\xc4\xdd\xa0\xd8\xe1\x94\xad\xd7\xc1\xe5\x9e\xdc\x13\x23\x12\xec\xc2\x81\xbe\x16\x9e\x79\x23\x12\x57\xa2\x28\x7c\x06\x41\xb4\xe0\x66\x69\xd6\xde\x21\xec\xec\xc2\x35\xaf\xc8\xaa\x63\x69\x41\x1f\xa4\x5b\x7f\x40\xf0\x8e\xe0\x9e\x40\x80\x06\xd4\x92\x02\x9a\xaf\x24\x51\xd4\x13\x4c\xbb\xb9\x79\x0e\x6b\x7a\x4a\x2a\xb8\x34\xa3\x1c\x3a\x30\xd3\xe6\x70\x62\x05\x5c\x04\x3e\xcb\xbe\xb4\x7e\xda\xd8\x69\x7c\x7f\x1f\xa1\xdf\x8d\xe2\x70\xf3\x74\x3e\xe5\x63\x9f\x2c\xe2\x66\xc4\xe6\xfd\x71\x97\x26\x2d\x02\x93\x7d\x80\xc5\x21\x35\x4e\xce\x77\x82\x78\xeb\xc0\xbc\x2d\xc1\x9e\x72\x65\x9e\x37\xe2\x41\x27\x1a\x9b\x8a\x6f\xb6\x66\x7c\xbc\xb0\x95\xd9\x07\x28\x34\x46\xdf\x57\xa3\x53\x4d\x0f\x71\x98\x9e\x95\x25\xf6\x53\xce\xd7\xa2\x89\xce\x5f\x50\xea\xc9\xd5\xb1\x50\x86\x85\x54\xc9\x9a\xca\x9b\x6a\xe0\x5e\xf2\x22\x3b\xc1\x43\xd9\x69\xac\x68\x35\x05\xd6\x19\x61\x86\x28\x0e\x34\x9c\xe8\x5a\xd6\x4d\x0d\xd8\xa7\x1a\x16\x06\xd9\x2d\x54\x2e\x65\xd4\xba\x2f\x6c\x99\xc0\x46\x67\x9f\x35\x53\xb3\xb7\x25\xe1\xc2\x5e\xc8\xbd\xa2\x2f\x88\x62\xa6\xad\xf8\xc1\x8b\x60\x70\x12\xad\xaf\x25\xab\x9a\xd1\xe6\xbe\x12\x60\xcd\x8c\x65\xa6\x3a\x94\x7f\xe8\x4d\xc7\x31\xc6\xd2\x6f\x58\x33\x78\x41\x13\x23\x55\x17\x30\xbe\xe7\x0e\x8b\x8f\x2e\x9d\xf6\x01\x08\xb9\x6c\x8c\xf7\xe0\x83\x6b\x1f\xb3\x3f\xcc\x45\x3b\x32\x94\x1a\xbe\xee\xd2\x04\xac\xa4\xf9\x8e\x24\x04\xff\x78\xd5\xba\xc7\xab\x69\xef\xa6\xec\xf6\xfd\xf5\x58\x6a\x11\x2c\xb7\x57\x4a\x0b\x98\x09\xda\x8c\x99\xed\x88\x62\x76\xf7\x62\x7c\x11\x85\x87\xca\xe4\x1d\x2a\xb2\x5b\xdd\xfd\x4e\xd1\x45\xa2\xfd\x45\x21\xf5\xb6\xfd\x11\x01\x68\x12\x45\xe9\x2f\x09\x77\x77\x7e\x81\x25\x6d\x75\x4f\xd7\x60\xcf\x3d\xae\x17\x18\xf7\xec\xa8\x40\x74\x3e\x15\x39\x16\x02\xff\x24\x69\x11\xa0\xff\x40\xf9\x26\x71\x98\xfc\x7d\xa4\x9d\xfe\xfb\x4f\xf4\x33\x50\x22\x74\xc5\x94\x3a\x74\xa3\x7f\x8d\x7f\xb0\xb1\x8c\x52\x03\xa5\xdc\x70\xb3\x6d\x56\x8e\x48\xee\x0e\xf6\x79\x02\x0b\x73\xdd\xac\xd5\x18\xa6\xbc\x68\x00\xaa\x77\xae\xd3\x2b\xe6\xef\x32\x93\xc3\x5b\x01\xde\x20\x78\xf9\xde\x7d\xbf\xa9\xf0\xc4\x92\xb0\xe5\x2b\x52\xec\xec\x0b\x11\x3a\x83\x5b\x7e\xea\x61\x38\xcf\x7d\x94\x7d\xaf\x77\x70\xe1\xb5\xde\x1e\xdb\x5c\x17\x8a\xd7\x06\x69\x55\x74\xce\x93\x6b\x72\x3b\x64\xdd\x76\x0c\x4a\x67\xa5\xf3\xeb\x2f\x0d\x83\xcd\x9c\xcf\xce\xcf\x67\xcf\x42\xcf\xd1\xcd\x6b\x07\xea\x5e\xe1\x72\xac\x3d\x52\xd5\xeb\x21\x63\x1e\xaf\x3b\x72\xeb\x1c\x53\x67\xf9\xf2\xa4\x92\x36\xc7\x81\x6b\x39\xe2\x0b\x4c\xd8\x3d\x79\xba\xb9\x7f\x03\x00\x00\xff\xff\xe9\xea\xf2\x73\x5d\x14\x00\x00")

func templatesBaseHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesBaseHtml,
		"templates/base.html",
	)
}

func templatesBaseHtml() (*asset, error) {
	bytes, err := templatesBaseHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/base.html", size: 5213, mode: os.FileMode(420), modTime: time.Unix(1455071029, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConfigBaseHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x54\x4d\x6e\xdb\x3a\x10\xde\xfb\x14\xf3\xf8\x02\x78\xf1\x2c\x31\xce\xcb\x2a\x90\xf4\x80\xb7\x48\x51\x74\x51\xa0\x49\x0f\x40\x4b\x23\x8b\x0e\xf5\x53\x71\xec\xc4\x08\x72\x8e\x9e\xa1\x27\xe8\xbe\xbd\x58\x87\xa4\x65\xd9\x4e\x8a\x06\x08\x90\x44\x33\xe4\x37\x1f\xbf\x7c\x9c\xe1\xe3\x23\x14\x58\xea\x06\x41\x94\x46\xd9\x4a\xc0\xd3\xd3\x04\x80\x97\x7b\xd5\x2c\x11\xce\xee\x70\x3b\x83\xb3\x8d\x32\x6b\xb4\x70\x95\x42\x7c\xed\x60\x1c\x9f\xe0\x6a\xb4\x56\x71\xc0\x90\x01\xed\x11\x00\x49\xa1\x37\x90\x73\x95\x4d\x85\x32\xd8\x13\xf8\xbf\x11\xd7\x3a\x76\x86\x89\xcc\x03\x19\xba\x58\x13\xb5\xcd\x80\xce\x4d\x6b\x51\x40\xa1\x48\x45\x85\xb6\xb5\xde\x53\x88\xec\xe7\xd7\x44\x06\xf4\x50\xec\xf8\x06\x15\xc3\xd1\x92\xcf\xce\x82\x4e\x6c\x8a\xbd\xe6\x5d\x3c\x46\x93\x03\x1f\x48\x93\x41\xe7\x03\xaf\xe9\x12\xe2\x5b\x97\x87\xf4\x30\x0e\x95\x2f\x53\x2c\x94\xf5\x0c\x93\xe4\xaf\xa2\xcd\x69\xdb\x21\x54\x54\x9b\x6c\x92\x84\x0f\x2b\xab\x50\x15\x41\x79\xe2\x0f\xcc\x6e\x3e\x5f\x7f\x72\xda\x08\xeb\xce\x28\x1a\x85\xc4\x4c\x94\xc8\x00\x0a\x05\x35\x92\x82\x46\xd5\x98\x8a\x8d\xc6\xfb\xae\x65\x43\x20\x6f\x1b\xc2\x86\x52\x71\xaf\x0b\xaa\xd2\x02\x37\x3a\xc7\xc8\x27\x33\xd0\x8d\x26\xad\x4c\x64\x73\xb6\x2f\x9d\xef\x0c\x4f\x8c\x6e\xee\xa0\xea\xb1\x4c\xa7\x15\x51\x77\x25\x65\xc9\x2c\x36\x5e\xb6\xed\xd2\xa0\xea\xb4\x8d\xf3\xb6\x96\xb9\xb5\xff\x95\xaa\xd6\x66\x9b\x7e\xec\xb0\xf9\xe7\x46\x35\xf6\xea\xf2\xfc\x7c\xf6\x6f\xf8\xd5\xa4\x8c\xce\x67\x97\x43\x34\x85\x1e\x4d\x3a\xb5\xb4\x35\xc8\xbd\x82\x34\x05\xe7\x41\x3a\x25\x7c\x20\xc7\x36\x3d\x3c\xdf\x6f\x89\x61\x4b\xf8\x5a\x31\xd6\x8a\x20\x50\x48\x4b\x8a\x74\xee\x30\x72\xd1\xb6\x64\xa9\x57\x5d\x5c\xeb\x26\x76\x55\x6f\x24\x54\x5d\xf7\x8c\xe6\x15\x65\xce\xad\x48\xdd\xa3\x6d\x6b\x3c\x92\x92\xc8\xe1\x7e\x93\x45\x5b\x6c\xb3\x67\x83\xe0\x6e\x4b\x71\xaf\xf4\x02\x74\x11\x52\xbe\xbc\x71\x10\x8e\xa0\x26\x7a\xb0\xd1\xfc\x62\xbf\x0b\xc7\x7d\xb2\x1b\xdc\x78\xe8\xfa\x67\x80\x81\xfd\x10\xb2\x1f\x8d\x7d\x18\xe2\x92\xad\xc5\x7e\x38\x3b\x64\xbf\x91\x35\xfc\x07\xa3\xac\xa4\x1b\x76\x9d\xfb\x51\xbd\x26\x2c\xc0\x87\x39\x9f\x7f\x04\xfd\xf1\x0d\x3e\xb0\xb5\x70\x8b\x7d\xbf\x1d\x57\xbf\x8f\x5c\x6a\x67\xb8\xeb\x4c\xcb\xad\xb9\xd4\x54\xad\x17\xbe\x21\xef\xb8\x92\x5c\xa1\xb4\xeb\x92\x59\x13\xbd\x17\xac\xa0\x54\x51\x80\xf2\xba\xd4\x19\xbc\xf3\x49\x22\x55\xf6\xb6\x63\x24\x3f\x3f\xfc\xac\x89\xec\xbd\xff\xfe\x91\xf0\x6f\x6a\x3b\x91\xfd\xaf\x72\xee\xc8\x96\x7f\xba\xa3\x8a\x44\x76\xd9\x4b\x57\x11\x1c\xdf\x65\x36\xef\x75\x47\x60\xfb\x7c\xd4\xa7\x56\xea\xe1\x74\x40\xdd\x9a\x34\x7a\x61\xe5\xea\xcb\x1a\x59\xef\x3c\x9e\xcf\xe3\x8b\x5d\xe6\x3b\x73\x65\x9d\x1d\x81\xf0\x05\xf6\xa1\xab\x57\xa7\xc3\xf5\xfa\x3a\x37\x43\xa7\x68\x7e\xa3\x7d\xfb\xf3\x3c\xf8\x87\x6f\x7c\x2e\x7f\x05\x00\x00\xff\xff\x6a\xa1\x5c\x28\x79\x06\x00\x00")

func templatesConfigBaseHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesConfigBaseHtml,
		"templates/config-base.html",
	)
}

func templatesConfigBaseHtml() (*asset, error) {
	bytes, err := templatesConfigBaseHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/config-base.html", size: 1657, mode: os.FileMode(420), modTime: time.Unix(1455062019, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLoginHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x53\x4d\x8f\xd4\x30\x0c\xbd\xef\xaf\xb0\x72\x0f\x15\x20\x71\x6a\x2b\x2e\x70\x42\x62\xc5\xc2\x0f\x48\x13\x77\x36\x22\x5f\x4a\xdd\xdd\x59\x56\xf3\xdf\x71\xda\x4e\x3b\x15\xea\x85\x3d\x8c\xc6\x8e\xfd\x9e\xdf\xb3\xdc\xd7\x57\x30\xd8\xdb\x80\x20\x74\x0c\x84\x81\x04\x5c\x2e\x77\xb5\xb1\x4f\xa0\x9d\x1a\x86\x46\xe4\xf8\x2c\xda\x3b\x80\x3a\x5d\x5f\x1c\x2a\x03\x84\x67\x92\x9a\x01\x98\xe7\xd8\x8f\x84\x46\xb4\x0f\xbf\xbe\xfe\x80\x6f\xf1\x64\x43\x5d\xa5\x09\x77\xc3\xa5\xa3\x93\xe7\x41\xbe\xff\x00\x25\xf2\x46\x7e\xba\x06\xb1\xef\x07\x24\xf9\x71\x1a\xc5\xa0\x3e\x66\x0f\x4a\x93\x8d\xa1\x11\xac\x32\xe3\x13\xe6\x81\x65\xba\x42\x5d\x44\x0a\xf0\x48\x8f\xd1\x34\xe2\xfe\xfb\xc3\xcf\x05\x57\x90\x16\x9d\x61\xae\xeb\xc3\x5e\x41\xe1\x95\xa7\x1c\xc7\x24\xb6\x06\x6e\x71\xaa\x43\x07\x5c\x6d\x04\x7a\x65\x9d\xd8\x24\x07\xca\xac\x71\x6a\x10\xed\x97\x52\xac\xab\x29\xdb\x11\xd8\x90\x46\xda\x4d\x59\x90\x30\x55\xa4\x3b\x09\xb0\x66\x65\xa7\x97\x84\x6b\x12\x94\xdf\x92\xe4\x94\xc6\xc7\xe8\x0c\xb2\x96\xdf\x2f\x0e\x3f\xe3\x59\xf9\xe4\xf0\x9d\x8e\x5e\x80\x1a\x29\x6a\x95\x2c\x29\x67\xff\x30\x2a\xc4\x80\x37\x56\xea\x8a\xcd\xfe\xb7\xf5\xc4\x8d\xcf\x31\x9b\x23\xf7\xf7\x4b\xfd\x78\x01\xc5\xe2\x3f\x2c\x07\xeb\x98\x77\xb0\x75\xcf\x6b\x58\xf3\x37\x99\xea\x46\xa2\x18\x96\x11\xc3\xd8\x79\x4b\xab\x9c\x8e\x02\xf0\x4f\xf2\xe5\xab\xd1\x91\x68\x97\x73\x9d\x31\x07\x63\xeb\x6a\x7f\x58\x9c\xf3\xe8\xe9\xc0\xe7\xb6\xe5\x8f\x6f\x15\x83\x29\x1f\xd1\xdf\x00\x00\x00\xff\xff\x1b\x4d\xb6\x09\x5c\x03\x00\x00")

func templatesLoginHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesLoginHtml,
		"templates/login.html",
	)
}

func templatesLoginHtml() (*asset, error) {
	bytes, err := templatesLoginHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/login.html", size: 860, mode: os.FileMode(420), modTime: time.Unix(1455052737, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRegisterHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x56\xcb\x72\xeb\x36\x0c\xdd\xf7\x2b\x30\x5a\x57\x71\xd3\xb5\xed\xe9\xb4\x93\x2e\x5b\x4f\x73\x9b\x3d\x25\x42\x16\x6b\x8a\xe4\x90\x90\x1f\xc9\xe4\xdf\x0b\x52\x0f\x3f\x93\x38\xad\x73\x37\x1a\x91\x3c\x04\x81\x03\x1c\x12\x2f\x2f\x20\xb1\x52\x06\x21\x2b\xad\x21\x34\x94\xc1\xeb\xeb\x0f\x00\x53\xa9\xd6\x50\x6a\x11\xc2\x2c\xf3\x76\x93\xcd\x79\xee\x78\xb6\xb4\x3a\xdf\x86\xfc\xfe\xe7\x7e\x8d\x57\xdd\xb0\xa6\x51\x48\x20\xdc\x52\x5e\xb2\x49\xf4\xdd\x7f\xd3\x12\xca\x6c\xfe\xf8\xf7\xef\x7f\xc1\x23\x52\xeb\xa6\x13\x77\xbe\x57\x68\xf4\x04\xe9\x9b\x2b\x53\x59\x88\xc6\xb2\xf9\xb7\x5a\x05\xd0\xd6\xae\xf8\xab\x56\x08\x02\x0c\x6e\x40\x99\x40\xc2\x94\x08\xb6\x82\x68\xf7\x0e\x16\x0c\x0f\x08\x1e\x97\x2a\xc4\x93\x77\xb6\xf5\x20\x64\xa3\x0c\xb4\x81\xc7\xc2\x48\xe0\x50\x2b\xb5\x6c\x3d\x02\xd5\x6c\xc9\x39\xad\x4a\x41\xca\x1a\x28\x50\xdb\xcd\xdd\xa1\x5f\x95\xf5\xcd\xe0\x5a\xfc\xcf\x6b\xeb\xd5\x33\x73\x25\x74\x06\xa2\x8c\xbb\x66\x19\xd3\xe8\x71\x8d\x3e\x74\x3c\xb2\xf1\x48\x63\x06\x0d\x52\x6d\xe5\x2c\x5b\xfc\xf9\xf8\x6d\xa4\x29\x1a\x55\xa8\x65\x40\xda\x4f\x1d\x73\x9b\x0e\x5a\x7a\xdb\xba\xec\x10\xc2\x20\x2d\xd8\x45\xe0\xf5\x59\x86\x8d\x50\xec\xc3\x69\x3a\x20\xfe\x35\x32\x8f\x3f\x86\x3c\x0f\xd2\x9e\x6c\xfe\x10\xf1\xd3\x49\x1a\x9d\x58\xbd\x94\xd6\xc1\xce\xfd\x4f\x27\x3e\x30\x5e\x19\xd7\xd2\x91\xb3\xfd\x59\x19\x28\x39\x7a\x46\x3b\x87\xe3\xc0\x88\x66\x3f\x70\x5a\x94\x58\x5b\x2d\x91\xe3\x58\xed\x34\xfe\x82\x5b\xd1\x38\x8d\x77\xa5\x6d\x98\xd6\x96\x6c\x29\x9c\x62\x8e\xd5\x33\xef\x32\xd6\xe0\x29\x11\x13\xf6\xf9\x88\xbe\xb3\x89\xcf\xf2\xe9\x18\xba\xb1\x5e\x7e\x82\xd2\x45\xbf\xe5\x96\xac\x46\x02\xcf\x5c\x39\xe6\xb8\x23\x76\x0f\xea\xb8\x1d\xc7\x1f\x30\xc5\x53\xc1\x09\x93\x0e\xaa\x51\xbb\x5f\xb5\x2d\x57\x6f\x07\x6d\xab\x8a\x2b\x95\x63\x8f\xd8\xbc\x48\xe0\xf9\x1f\x16\x02\x0b\x58\x25\x2d\x05\xf2\x42\x19\x0a\x3f\xc2\x3f\x6d\xa0\x28\x33\x56\xe7\xd2\x5a\x09\x21\xaa\x3f\xca\x93\x55\x08\x25\x9f\xe9\xb1\xc1\xa6\x40\xcf\x12\x8b\x3e\xdc\x2c\x81\x57\x67\xec\x49\x05\x55\x28\xad\x68\x77\x9b\x9c\x1d\x5e\x93\x42\x2a\x7b\x86\x18\x5c\x3c\x9f\x3f\xca\xf8\x7a\xf4\x2b\x77\x6d\xc1\xf7\xd1\x90\xe5\xce\x6a\x9f\xe2\x3d\x2a\x83\xb5\xd0\x6d\xcc\x7a\x8f\x2e\x6b\x2c\x57\x28\x2f\x1d\xb3\x48\x90\x73\xbf\x2e\x11\x70\xb1\x5e\xbe\x28\x4e\xaf\xd6\x82\xf0\xea\x40\x7b\xf8\xc5\x08\xbb\xb5\xff\x17\xe2\x3b\xa2\x38\x2c\xfd\x07\x23\x0a\x8d\xd0\xbb\x33\xbe\x40\x01\x54\x95\xca\x7c\x23\x0c\x01\x59\xa8\x95\x44\x88\x8f\xc2\x8e\x6a\x65\x96\x50\x79\xdb\xa4\xf7\xa6\xcb\xd8\x05\x05\xdc\xfa\x56\xbb\x5a\x14\xbf\x75\x4f\x3f\x3c\xb0\x34\xa5\x64\x67\xaf\xd6\xc6\xfd\x95\xda\x48\xd5\x59\xd8\xed\x7f\x2d\x9b\x78\x69\x48\x57\x5b\xb2\x61\xa8\x97\xd1\xe4\xf8\xb4\xec\x21\x97\xac\xa5\xe0\x60\x91\x00\x37\xd3\xc2\x4d\xe2\x5a\x73\xa5\x7c\x10\x57\x0f\x79\x3b\xae\xa7\x04\xf8\x3e\x02\x18\xaa\x04\xfa\x96\x11\x36\x4a\x6b\xee\xd8\xb8\x75\x02\xa7\x4a\xe2\xc6\x2a\xa4\x3e\xab\xf3\x3a\x8a\xa1\x40\x48\x92\xd6\x51\x30\x5d\x4f\x96\xda\xc0\x0a\x51\x7e\xa9\x12\xde\xbb\xcf\x87\x97\xed\xbc\x72\x8b\x96\x88\xbb\xc1\x2e\x21\xa1\x2d\x1a\x45\x23\x15\x05\x71\x9b\x48\x26\xe7\xbe\x59\xb4\x9a\x92\x7a\xba\x6e\x72\x3a\xe9\xf6\x7d\x2e\x92\xe9\xe4\xb4\x1b\xe4\x19\x8e\xa7\xef\xb8\x07\x70\xff\xc3\x9d\x26\x32\xb3\xdc\xa3\xff\x1b\x00\x00\xff\xff\xc1\x79\xcc\x3e\xbb\x0b\x00\x00")

func templatesRegisterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesRegisterHtml,
		"templates/register.html",
	)
}

func templatesRegisterHtml() (*asset, error) {
	bytes, err := templatesRegisterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/register.html", size: 3003, mode: os.FileMode(420), modTime: time.Unix(1455052725, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSettingsHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x56\x41\x6f\xdb\x3a\x0c\xbe\xf7\x57\x10\xc2\xbb\xba\x79\x7d\xc7\x87\x24\x87\x0d\x1d\xb0\xd3\x8a\x65\xeb\x5d\xb6\xe9\x5a\xa8\x2c\x79\x32\xed\x34\x0b\xf2\xdf\x47\xc9\x96\x63\xa7\x2d\x9a\xa2\x2d\x76\x08\x62\x91\x9f\xc9\x8f\x9f\x48\xc9\xfb\x3d\xe4\x58\x28\x83\x20\x32\x6b\x08\x0d\x09\x38\x1c\x2e\xd8\xfc\x4f\x83\x44\xca\xdc\x35\xf0\xff\x0a\x2e\x37\x71\xc1\xce\x65\xae\x3a\xc8\xb4\x6c\x9a\x95\x70\x76\x2b\xd6\x17\x00\x53\x5b\x66\x75\xf2\xd0\x24\x57\xff\x05\x0f\xfb\x0a\xeb\xaa\xe8\xf4\xcf\x49\x69\x9d\xfa\xcd\xe9\xa4\x16\x20\x33\x52\xd6\xac\x04\xa7\x74\xd8\xa1\x6b\x98\x4a\x4c\xed\xb9\x08\xa8\x90\x4a\x9b\xaf\xc4\xcd\xb7\xcd\x8f\x21\xe4\x3c\x61\x88\x79\xe7\x6c\x5b\x8f\x6e\x06\x68\x99\xa2\x7e\xc4\x09\xfc\x53\x95\x27\xfe\xc1\x90\xe3\x45\xc0\x89\xf5\xad\x6a\x54\xaa\xb4\xa2\xdd\x72\x11\x4c\x93\x50\x4f\x15\x17\x03\x5d\xfd\x3b\x49\x3a\xc7\x3a\x99\x2b\x3b\xf3\x46\x5a\x73\x1b\x5b\x95\xa9\x5b\x02\xc5\x55\x76\x23\x8f\xa4\x6e\x53\xad\x32\x01\xb4\xab\x31\x46\x03\x23\x2b\x9c\xa2\x04\x74\x52\xb7\x6c\x8a\x68\x16\x52\x15\x80\xbf\x8e\x3b\x78\x79\xac\x0d\x46\xd8\xe1\x90\x95\x98\xdd\x63\xbe\xdf\xa3\xc9\x0f\x87\x35\x9c\x70\xba\x09\xc0\x39\xf9\x53\x65\xbc\x89\x0b\xfe\x30\x01\x9c\xea\x24\xe1\xd9\x0a\x44\xf8\x4b\x12\x44\xdc\x23\x0d\x4e\x25\xe8\x81\xaf\xd7\xa0\xa9\xa5\x09\xc5\x94\xa8\xeb\x4f\xda\x66\xf7\x22\x6a\xe2\x2d\x49\x1a\x4c\xeb\x6b\x23\x53\x8d\x30\xf0\x01\x65\x1a\x92\x26\xc3\xc6\xb3\xdf\xd9\x16\xb6\xd2\x10\x90\x85\x52\xe5\x08\x7e\x38\x76\x54\x72\x35\x50\x38\x5b\x01\x95\xfc\x66\xd8\xa3\xcb\xe5\xc2\x27\x9c\x34\xec\x94\xd0\x7c\xf1\xde\x63\xf3\xb9\x3f\x35\xe0\xba\x4a\x31\xcf\x99\xdc\x59\xd3\x73\x75\xc6\xf4\x84\xad\x49\xed\xc3\x6b\xfb\x07\x3d\x93\xba\xb4\x64\x9b\xd8\x38\x63\xa8\xa1\x77\x66\x90\xbe\x59\x8e\x9d\x12\x2a\xb9\x09\xce\x17\x1b\x24\x60\xa1\x07\xbf\x69\x52\xde\x54\x6c\xc7\xed\xf1\x42\xb1\x11\xf2\x64\xb1\xb7\xc1\x79\x66\xb1\x3d\xf8\xe3\x46\x22\xf6\x11\x0c\xf7\x11\x6c\x95\xd6\x20\xb5\xb6\x5b\xa8\x55\x46\xad\xe3\xf9\x90\x26\x87\xbe\x24\x3f\x1e\x29\x42\x38\x0d\xb4\x1f\x21\x3f\x38\x0e\x36\x3f\xbf\x7c\x87\x02\x31\xff\x9b\xb3\xf1\xb5\xaa\xad\xa3\xb7\x0d\x44\x3d\x63\x14\x13\xf0\x39\x41\x7c\x88\x9f\x34\x8b\x84\xd2\x61\x71\x72\x95\x96\x8a\xd2\x96\x77\x95\x12\x15\xe8\x84\x3b\x75\xa0\xd6\x1f\x24\x9b\x11\xb2\x5c\xc8\xf9\x26\xd6\xef\xa3\xdb\xb3\xf7\x67\x53\x4d\xf4\xb3\x45\xc1\x7d\x99\xf4\x6b\x7d\x37\xae\xe7\x82\xa4\x2d\x91\x35\x43\xaf\x37\x6d\x5a\x29\x1a\x1b\x29\x25\x03\xfc\x4b\xf8\x93\x46\xb6\x9a\xc4\x7a\x23\x3b\x5c\x2e\xfa\x57\xd6\xac\xcf\x11\x27\x9e\xd0\xaa\x75\x3a\x51\x26\xc7\x87\x5e\xa3\xcc\x9f\xc5\x7a\xa6\xc9\x73\x22\x2c\x17\xbe\xf0\xf0\x21\xd4\x1b\x87\x3f\x8e\xce\xa3\xe4\x3f\x9a\xfe\x04\x00\x00\xff\xff\x29\xe6\x20\x78\x69\x09\x00\x00")

func templatesSettingsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesSettingsHtml,
		"templates/settings.html",
	)
}

func templatesSettingsHtml() (*asset, error) {
	bytes, err := templatesSettingsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/settings.html", size: 2409, mode: os.FileMode(420), modTime: time.Unix(1455071338, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesShitbucketImportHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x92\xc1\x6e\x1b\x21\x10\x86\xef\x79\x8a\x11\x97\x9e\x36\xa8\x3d\x56\xde\x3d\x54\xbd\x54\xaa\xd4\xaa\x69\x1f\x80\x85\x49\x16\x85\x65\x10\x0c\xa9\xd3\xc8\xef\xde\x01\x6f\xd6\xf6\xc1\xb2\x99\xdf\x7c\xff\xfc\xc3\xbc\xbd\x81\xc3\x47\x1f\x11\x94\xa5\xc8\x18\x59\xc1\xe9\x74\x77\x70\xfe\x05\x6c\x30\xa5\x8c\x2a\xd3\x5f\x35\xdd\x01\x5c\xd7\x2c\x85\xe1\x58\x86\x8f\x9f\xba\x22\xda\x23\xe5\x15\x8c\x65\x4f\x71\x54\x02\xcd\xf8\x82\xb9\x08\xb5\x2c\x9e\xe7\x6a\x9f\x91\x07\xbf\x26\xca\x9d\xaf\x60\x45\x5e\xc8\x8d\xea\xe7\x8f\x87\xdf\x1b\xe3\xd6\xa1\x01\x87\xa7\x4c\x35\xed\xb2\xfc\x21\x98\x19\x03\x88\x36\xaa\x9a\x83\xba\xb4\x13\x39\x4b\x4b\x5d\x56\xd3\xc3\xee\x09\x7f\x7e\x7d\x3f\xe8\x5e\xbe\xa2\xf8\x98\x2a\xdf\x18\x6d\x00\x05\xde\x75\xf2\x50\xea\xbc\x7a\xe9\x95\x5f\x13\x8e\x8a\xf1\x28\xbf\xa3\x59\x71\xf3\x4d\xc1\x58\x5c\x28\x38\x94\x4e\x16\xe6\x54\x3e\x6b\x7d\x89\x7a\x8f\x47\xb3\xa6\x80\xf7\x96\x56\x6d\x92\xd7\x72\xa9\x28\x30\x95\xc9\xca\x91\x4d\xf0\xff\x04\x15\x29\xe2\x75\xba\x92\x4c\xec\x1d\x2c\x18\xd2\x97\x40\xf6\x79\x4f\xd8\x2a\xc3\xdc\x4b\xd3\x57\x8a\x1f\xb8\x0d\xe1\x49\x02\x32\x81\x71\x0e\x4c\x7c\x85\xd9\x14\x6f\x9b\xc9\x02\x36\xa3\x2b\x4d\xe3\x05\xcf\x33\x68\xec\x7d\xce\x5a\x06\xbd\x1f\xe6\xca\x4c\x71\x4b\xfa\x9e\x7b\xb3\x9d\x39\x82\x7c\x06\xd9\x11\x53\x03\xab\xe9\x5b\x7f\xc3\x83\x3e\x5f\xda\x1e\x5f\xb7\x19\xf6\x15\x39\x83\xb7\x2f\x59\x03\x8c\xae\xad\xd3\xff\x00\x00\x00\xff\xff\x61\x87\x0f\x00\x66\x02\x00\x00")

func templatesShitbucketImportHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesShitbucketImportHtml,
		"templates/shitbucket-import.html",
	)
}

func templatesShitbucketImportHtml() (*asset, error) {
	bytes, err := templatesShitbucketImportHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/shitbucket-import.html", size: 614, mode: os.FileMode(420), modTime: time.Unix(1455071795, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTagIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8f\x3d\x4b\xc6\x30\x10\xc7\xf7\x7e\x8a\xe3\x70\xd0\xa1\x7d\xd0\x51\xd2\x2e\xba\x08\xf2\x0c\xa2\x1f\x20\x36\xf7\xa4\x81\x78\x81\x24\xb6\x43\xc8\x77\xf7\x1a\x5a\x70\x0a\x77\xbf\xdc\xff\xa5\x14\x30\x74\x73\x4c\x80\x73\xe0\x4c\x9c\x11\x6a\xed\x94\x71\x2b\xcc\x5e\xa7\x34\x62\x0c\x1b\x4e\x1d\xc0\xff\xdd\x1c\x7c\xef\x6d\xff\xf8\xd4\x08\x40\x29\x51\xb3\x25\xb8\xcb\xda\xc2\xf3\x08\xc3\xa7\xb6\x49\x74\x76\xa6\xf4\x79\xf5\x9d\x19\x61\x89\x74\x1b\x51\x7c\x23\xad\x14\x93\x18\xcb\x4d\xbf\x3a\xda\x10\xd0\x19\x6c\x1a\xc3\xdb\xab\xc4\xc0\x49\xbe\xb5\xf1\xaa\x7f\x48\x16\x70\x5f\x4a\x9b\xbf\x3e\xde\x5f\xc2\x2f\xe7\x5a\x1f\xd4\x45\x9f\x19\xc8\x27\x3a\x4c\xaf\x21\x2f\x8e\xed\x70\x12\x36\x0d\xa8\x8b\x94\x98\xba\xe3\x11\x75\x01\x7b\xdf\xbf\x00\x00\x00\xff\xff\x61\xf7\x9e\x87\x07\x01\x00\x00")

func templatesTagIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTagIndexHtml,
		"templates/tag-index.html",
	)
}

func templatesTagIndexHtml() (*asset, error) {
	bytes, err := templatesTagIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/tag-index.html", size: 263, mode: os.FileMode(420), modTime: time.Unix(1454988699, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUrlEditHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x54\xc1\x6e\xdb\x30\x0c\xbd\xf7\x2b\x08\x61\x57\x37\xe8\x8e\x83\xed\xc3\xd0\x6d\x18\x30\x60\xc3\xda\x7d\x00\x2d\xd3\xb5\x30\x59\x32\x24\x2a\x4b\x5a\xf4\xdf\x47\x29\x9e\xe1\x36\x1d\x8a\x61\x87\x20\x12\xf9\xc2\xf7\xf8\x44\xe6\xe1\x01\x7a\x1a\x8c\x23\x50\xda\x3b\x26\xc7\x0a\x1e\x1f\x2f\x24\xfc\x26\x05\x0b\xef\x1a\xb8\xfc\xf1\xfd\x4b\x0e\xd5\xbd\xd9\x83\xb6\x18\x63\xa3\x82\xff\xa5\xda\x0b\x80\x6d\x4c\x7b\x5b\x1d\x62\x75\xf5\xb6\x64\x24\x37\xf8\x30\xfd\x49\xe6\x73\x35\xfa\x60\xee\x85\x04\xad\x02\xd4\x6c\xbc\x6b\x94\x10\x05\xda\x53\x88\x22\x40\x08\xab\x88\x7b\x52\xa0\x4c\xaf\x8a\x80\xcb\xcf\xd7\xc2\xad\x60\x22\x1e\x7d\xdf\xa8\x6f\x5f\x6f\x6e\x97\xfa\x4f\xd9\x0b\xc1\x5d\xf0\x69\x5e\xd3\x02\xb0\xd8\x91\x05\xc9\x35\x8a\x0d\x5b\xa9\xfc\x4c\x2c\xe4\xa6\x83\x5c\x0a\x52\xb5\xb7\x19\x55\xef\xca\x6d\x53\xe7\x85\x36\xaf\x36\x3c\x82\x30\x6e\x4e\x0c\xa6\x7f\x4e\x54\x74\x2d\x24\x0a\xf8\x38\x93\x20\xe8\x20\x2e\x3b\x9c\x68\x45\xef\xd1\x26\xca\x6e\x94\xa6\x8b\x8a\xdc\xf6\x6c\x51\xd3\xe8\x6d\x4f\xd2\xc1\x87\x03\x4e\xb3\x25\xc8\xef\x51\x10\xdb\x4e\x77\x22\x71\xf5\xe5\xc9\xe5\x9f\x4c\xc2\xbb\xf8\xba\x47\x02\xfa\x4f\x8b\xb6\x34\xaf\x3a\x54\xc0\xab\x41\xa7\xb1\xf8\x44\x9c\x65\x7c\xf4\xe1\xda\x44\x71\xe9\x08\x67\x76\x69\x3f\x09\x1f\x85\x2a\x6a\x43\x4e\x13\x58\x41\x42\x8e\x1a\x2b\xf3\x26\x23\x98\xd8\x6b\x9c\x8d\xcc\xa3\xb9\x97\xda\xce\xbb\xbf\x5a\x2a\xd7\x38\xa3\x2b\xea\x47\xb2\xf3\x7b\xeb\xf5\xcf\x33\xa7\xb2\x55\xe5\xe4\x87\x21\x12\x8b\x75\x19\x5b\x75\x05\x5c\x7c\x83\x29\x45\x86\x8e\x20\xd2\x8c\x01\x99\x7a\xe8\x8e\x80\x20\xc5\xb5\xcc\x5d\xe6\x78\xf9\x15\xbb\xc4\xec\xdd\x62\x4f\x4c\xdd\x64\x78\xa5\xef\xd8\x81\x7c\x2a\x59\x64\x4c\x96\x55\x7b\x23\x4b\x54\xef\x4e\x3f\x69\x6b\xdc\xe0\x14\x8c\x81\x86\xf3\xb5\x33\xae\xa7\x43\xde\x7c\xd5\x6a\x14\xb3\x6c\xbd\xc3\x65\x8d\x77\xf9\x81\xca\xb2\x9f\x04\x2d\x5f\x52\x81\x5c\x9f\xff\x18\x7e\x07\x00\x00\xff\xff\x9a\x0d\xa3\x53\x43\x04\x00\x00")

func templatesUrlEditHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesUrlEditHtml,
		"templates/url-edit.html",
	)
}

func templatesUrlEditHtml() (*asset, error) {
	bytes, err := templatesUrlEditHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/url-edit.html", size: 1091, mode: os.FileMode(420), modTime: time.Unix(1455002156, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUrlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x51\xc1\x4e\xeb\x30\x10\xbc\xf7\x2b\x46\xd6\xd3\xbb\x25\x15\x1c\x21\xcd\x05\x89\x13\x27\x04\x07\x8e\xc6\xd9\x34\x96\x1c\x3b\x5a\x6f\x4a\xa5\xa8\xff\xce\xb6\x69\x4a\xa9\x38\xd9\x33\xe3\x9d\x9d\x5d\x4f\x13\x1a\x6a\x7d\x24\x18\x97\xa2\x50\x14\x83\xc3\x61\x55\x35\x7e\x07\x17\x6c\xce\x1b\xc3\xe9\xcb\xd4\x2b\xe0\x9a\x73\x29\x14\xfb\x5c\xdc\xdd\x9f\x14\x60\x9a\x7c\x8b\xf2\x29\x8d\x51\xb4\x7a\x66\xd8\xc6\x2d\xe1\xdf\xc8\x01\x0f\x1b\x94\xef\xaf\x2f\x19\x67\xf1\x28\x43\xa8\x1f\x82\x15\xed\xac\x4f\x8a\xc1\xb2\x78\x1b\xcc\x5c\x70\x31\x01\xc5\x66\x41\x7f\x85\xba\xe5\x6f\x83\x9d\xf4\x61\x51\x85\xf6\x52\xf4\xa3\x50\x63\xea\x67\x0d\xdb\x1c\x3b\xcc\xb1\xb5\xc9\xff\x8e\x42\xf0\xc3\x63\xb5\x1e\x2e\xd6\x6b\xf5\x9e\xc1\xd5\x75\x9a\x28\x64\x5a\x52\xfd\x76\x77\xba\x42\x62\x53\xbf\x75\xc4\x04\x9f\x11\x93\x74\x3e\x6e\x71\xc2\x9f\xa3\x20\xfb\x40\xd1\x51\x89\x8f\x34\xc2\xd9\x08\xdf\x0f\x89\x05\x8d\x15\x8b\x96\x53\x0f\xad\x20\x46\x26\xde\x79\x47\x19\x3e\x42\x09\x54\x16\x1d\x53\xbb\x31\x9a\x99\x69\x47\x9c\x75\x75\x99\x44\xd4\x3d\x1f\x7f\xcd\xd4\x0b\xaa\xd6\xb6\x2e\x2f\x63\x5c\x6f\xf1\x3c\xc5\xf9\xf8\x51\xbe\x03\x00\x00\xff\xff\xfc\xa1\xc2\xf7\x08\x02\x00\x00")

func templatesUrlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesUrlIndexHtml,
		"templates/url-index.html",
	)
}

func templatesUrlIndexHtml() (*asset, error) {
	bytes, err := templatesUrlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/url-index.html", size: 520, mode: os.FileMode(420), modTime: time.Unix(1455071672, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUrlNewHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x54\x4d\x6f\xd4\x30\x10\xbd\xf7\x57\x8c\x7c\xe1\x14\xa2\xf6\x88\x92\x3d\x70\x46\x02\x41\x39\x70\x74\x9c\xc9\xc6\xc2\x5f\xb2\xc7\x65\xdb\xaa\xff\x9d\xb1\x37\xbb\x24\x6c\x25\xb4\xa8\x87\x28\x1e\xcf\xf8\xbd\x37\xcf\x1f\xcf\xcf\x30\xe2\xa4\x1d\x82\x50\xde\x11\x3a\x12\xf0\xf2\x72\xd3\x8d\xfa\x01\x94\x91\x29\xf5\x22\xfa\x5f\x62\x77\x03\xb0\x9e\x53\xde\x34\x87\xd4\xdc\xde\xd5\x0c\xe7\x26\x1f\xed\x29\x59\xc6\xcd\xec\xa3\x7e\x62\x44\x69\x04\x48\x45\xda\xbb\x5e\x30\x59\xc4\x07\x8c\x89\xd9\x72\x34\x4d\xca\x83\xd5\x95\x50\x80\x45\x9a\xfd\xd8\x8b\x2f\x9f\xbf\xdd\x2f\xa0\x5b\xca\x8a\xba\x8f\x3e\x87\x73\x9a\x0b\x8c\x1c\xd0\x00\xe7\xfa\x82\x28\xd6\xfa\xec\xd8\xdc\xc2\x49\x68\x19\x39\x8a\x1c\xd5\x15\x62\xf7\xfd\xeb\xa7\xae\xad\xe3\x15\xda\x5f\x1d\x16\x84\x3f\x10\x77\x2b\x5e\xae\xd5\x2e\x64\xda\x88\x5b\x18\x04\xe8\xb1\xdf\xf4\x47\x8f\x01\x7b\x41\x78\xe0\xb1\x93\x16\x17\xad\xc1\x48\x85\xb3\x37\x23\xb2\xfa\x99\x28\xa4\x0f\x6d\x8b\x07\x69\x83\xc1\xf7\xca\x5b\x36\x2e\x93\x57\x32\x68\x76\x51\x3f\xf1\x32\xe7\x1d\xae\xbb\x6f\x59\xf0\xd9\xab\x4d\x70\x8d\x71\x24\xf7\xe9\x0a\xe7\xee\xb9\xfc\x2d\xac\x2b\x2e\x6d\xa8\xb7\x26\x5e\x9a\x76\x2c\xde\xb8\xc6\x2e\x31\x14\xc6\x26\x29\x8d\x4e\x21\x18\x9d\x02\x94\x59\x6d\xf8\x9c\x5d\xe7\x20\x87\x29\x48\x57\x85\xcd\x68\xc2\x47\xe3\xd5\x4f\x71\x71\xe4\x61\xe9\xcf\x4f\x53\x42\x62\xa7\x4a\x6d\x33\xd4\xe2\x6a\x0e\xd8\x9c\x08\x06\x84\x84\x41\x46\x49\x38\xc2\xf0\x08\x12\x18\x5c\x61\xd7\x16\x8e\xff\xdc\xb4\xd7\x6e\x60\x95\x93\xec\x6b\xc2\x4a\x6c\xf6\xe7\x78\xbb\x09\x43\x26\xf2\x6e\xb1\xf9\x74\x52\x17\xf0\x81\x1c\xf0\xd7\xf0\xd3\x20\xb3\x21\xb1\xfb\xe1\xf3\xbb\xc8\x0d\xcd\x9a\x48\xbb\x3d\xdf\xd6\xae\x3d\x02\xec\x3a\xb9\x5a\x25\x60\x8e\x38\x5d\xde\x74\xed\x46\x3c\xd4\x8b\xbe\x53\x92\xf7\xc9\x74\xad\xfc\xf7\x39\xee\xda\x62\x43\x7d\x7b\x8e\x93\xcb\x8f\xd1\xd1\x8d\xe5\x9d\xfa\x1d\x00\x00\xff\xff\xd0\x26\x4f\x40\xbf\x04\x00\x00")

func templatesUrlNewHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesUrlNewHtml,
		"templates/url-new.html",
	)
}

func templatesUrlNewHtml() (*asset, error) {
	bytes, err := templatesUrlNewHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/url-new.html", size: 1215, mode: os.FileMode(420), modTime: time.Unix(1455001798, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUrlPartialHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x52\xcd\x6e\xb3\x30\x10\xbc\xe7\x29\x56\x56\xae\x04\x7d\x5f\x6f\x15\x20\x55\xfd\x51\x2f\xbd\xa5\xe7\x6a\x83\x17\x62\xd5\x31\xc8\x6c\xda\x4a\x11\xef\xde\x35\xb1\x03\x89\x7a\xc2\xeb\x99\xdd\x19\xef\x70\x3a\x81\xa6\xc6\x38\x02\x75\xf4\x36\xeb\xd1\xb3\x41\xab\x60\x1c\x57\x02\xad\xe5\x0e\xee\x4b\xd8\x84\xba\xd0\xe6\x0b\x6a\x8b\xc3\x50\x2a\xdf\x7d\xab\x6a\x05\xb0\xbc\xab\x3b\x9b\xd9\x36\xfb\xf7\x7f\x42\xae\xb1\x30\xbb\xee\x1c\xa3\x28\x79\xc1\x27\x82\x50\xf6\x77\x55\x3c\x4a\x81\xb0\xf7\xd4\x94\x2a\x0a\x6f\xde\x45\x7c\x1c\x15\x30\xfa\x96\xb8\x54\x1f\x3b\x8b\xee\x53\x55\x0b\x7c\x6b\xd8\x92\x90\x8a\x1c\xd3\xa0\x22\x0f\x43\x53\xd1\x27\x07\x4c\x3f\x9c\x1d\x8e\x4c\x5a\x55\xc5\x70\x40\x6b\xab\x3f\x04\x4d\x10\xbc\x08\xbc\x74\xfe\x80\x2c\x2d\x8f\x9e\x50\x3e\x0f\x1c\xa5\x20\x83\x6b\x93\x45\x1e\x47\xe6\xfd\x45\x5b\x18\xa6\x39\x93\x5e\x71\xd8\x62\x3b\x84\x2d\x46\x5f\x26\xf9\x6a\x10\x1a\xcc\x18\x5b\xb1\x95\x9b\x2a\x74\x79\x74\x2d\xc1\x5a\xee\xc2\xee\xa7\x01\xb1\xfb\xe2\x38\x17\x30\x0f\x16\x02\x29\x59\x3e\x9f\x27\x7f\x52\x92\xd3\xb3\xde\x5c\x27\x03\x37\xe1\x60\xcd\xa6\x73\x83\x5a\xc4\x31\xf4\xe8\x12\x65\x6f\xb4\x26\xb7\x40\xa7\xb8\x22\xb8\x63\xa7\x6e\x36\xf9\x86\x0e\x5b\x0a\xae\xd3\x4e\x9f\xb5\xe1\xc9\xda\x55\x1f\x4c\xb1\xe8\xf0\x60\x7f\x3b\xe3\x89\x2c\x31\xa5\xfe\x73\xb5\xc8\x39\x24\x1d\x2c\xce\xb9\xcb\x93\xe2\x9f\x97\x8e\xf1\x10\x3f\xf3\x12\x7e\x03\x00\x00\xff\xff\x94\xed\xdc\x0e\xf6\x02\x00\x00")

func templatesUrlPartialHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesUrlPartialHtml,
		"templates/url-partial.html",
	)
}

func templatesUrlPartialHtml() (*asset, error) {
	bytes, err := templatesUrlPartialHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/url-partial.html", size: 758, mode: os.FileMode(420), modTime: time.Unix(1454376085, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUrlViewHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x56\x48\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\xce\xcf\x2b\x49\xcd\x2b\x51\x52\xa8\xad\xe5\x02\x0a\x97\xa4\xe6\x16\xe4\x24\x96\x00\x25\x4a\x8b\x72\x74\x0b\x12\x8b\x4a\x32\x13\x73\x94\x14\xf4\x42\x83\x7c\x20\x2a\x52\xf3\x52\x80\x34\x20\x00\x00\xff\xff\x74\xa2\xa0\xcb\x41\x00\x00\x00")

func templatesUrlViewHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesUrlViewHtml,
		"templates/url-view.html",
	)
}

func templatesUrlViewHtml() (*asset, error) {
	bytes, err := templatesUrlViewHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/url-view.html", size: 65, mode: os.FileMode(420), modTime: time.Unix(1455001107, 0)}
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
	"templates/base.html":              templatesBaseHtml,
	"templates/config-base.html":       templatesConfigBaseHtml,
	"templates/login.html":             templatesLoginHtml,
	"templates/register.html":          templatesRegisterHtml,
	"templates/settings.html":          templatesSettingsHtml,
	"templates/shitbucket-import.html": templatesShitbucketImportHtml,
	"templates/tag-index.html":         templatesTagIndexHtml,
	"templates/url-edit.html":          templatesUrlEditHtml,
	"templates/url-index.html":         templatesUrlIndexHtml,
	"templates/url-new.html":           templatesUrlNewHtml,
	"templates/url-partial.html":       templatesUrlPartialHtml,
	"templates/url-view.html":          templatesUrlViewHtml,
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
		"base.html":              &bintree{templatesBaseHtml, map[string]*bintree{}},
		"config-base.html":       &bintree{templatesConfigBaseHtml, map[string]*bintree{}},
		"login.html":             &bintree{templatesLoginHtml, map[string]*bintree{}},
		"register.html":          &bintree{templatesRegisterHtml, map[string]*bintree{}},
		"settings.html":          &bintree{templatesSettingsHtml, map[string]*bintree{}},
		"shitbucket-import.html": &bintree{templatesShitbucketImportHtml, map[string]*bintree{}},
		"tag-index.html":         &bintree{templatesTagIndexHtml, map[string]*bintree{}},
		"url-edit.html":          &bintree{templatesUrlEditHtml, map[string]*bintree{}},
		"url-index.html":         &bintree{templatesUrlIndexHtml, map[string]*bintree{}},
		"url-new.html":           &bintree{templatesUrlNewHtml, map[string]*bintree{}},
		"url-partial.html":       &bintree{templatesUrlPartialHtml, map[string]*bintree{}},
		"url-view.html":          &bintree{templatesUrlViewHtml, map[string]*bintree{}},
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