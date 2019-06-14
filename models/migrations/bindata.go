// Code generated by go-bindata.
// sources:
// models/sql/1558054487_initial_schema.down.sql
// models/sql/1558054487_initial_schema.up.sql
// models/sql/lock.json
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

var __1558054487_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x72\x75\xf7\xf4\xb3\xe6\xe2\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\x46\x15\x2a\x4a\x2d\xc8\x2f\xce\x2c\xc9\x2f\xca\x4c\x45\x93\x29\x28\xcd\xc9\x89\x2f\x4a\x2d\x2c\x4d\x2d\x2e\x89\x2f\x4a\x2d\xcb\x4c\x2d\xc7\xa7\x22\x39\x3f\x37\x37\x35\xaf\x04\x8f\x12\x34\xa9\xfc\xa2\xf4\xc4\xbc\xcc\xaa\xc4\x92\xcc\xfc\x3c\x34\xa9\xcc\xe2\xe2\xd2\x54\x1c\x26\x82\xe5\x40\x62\xce\xfe\xbe\xbe\x9e\x21\xd6\x5c\x80\x00\x00\x00\xff\xff\xa6\x0c\xa3\xe3\xe8\x00\x00\x00")

func _1558054487_initial_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1558054487_initial_schemaDownSql,
		"1558054487_initial_schema.down.sql",
	)
}

func _1558054487_initial_schemaDownSql() (*asset, error) {
	bytes, err := _1558054487_initial_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1558054487_initial_schema.down.sql", size: 232, mode: os.FileMode(509), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1558054487_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\xcd\x6e\xe3\x36\x10\x3e\xc7\x4f\xa1\x63\x0b\xe4\x0d\x72\xca\x16\x41\x11\x34\xc9\x16\x41\x7a\x58\x14\x05\x31\x12\x47\xd2\xac\x29\x52\xcb\x19\xc5\x75\x9e\xbe\xa0\x2c\x4b\x94\x44\x27\xbb\x5e\x74\x9b\xde\xac\xf9\x28\x6a\x7e\xbf\x99\xf1\x87\x9b\x5f\x6f\x1f\xae\x36\x9b\x5f\x1e\x6f\xae\x9f\x6e\xb2\xa7\xeb\x0f\x77\x37\x19\x31\x77\xc8\xd9\x4f\x9b\x0b\xd2\x19\xa3\x27\x30\xd9\xef\x8f\xb7\xf7\xd7\x8f\x9f\xb2\xdf\x6e\x3e\x5d\x6e\x2e\x6c\xd7\xe4\xe8\xb3\x9c\x2a\xb2\x72\xb9\xb9\x60\x01\xc1\x4c\xf0\xef\xf0\x60\x5c\xb1\x45\x9d\xe5\xce\x19\x04\x7b\xb9\xb9\x10\x12\x33\xa2\xb9\xd3\xfb\xe3\xef\xc2\x35\x0d\x5a\xe1\xe9\xa2\xc2\x38\x46\xad\x40\x32\xa1\x06\x59\xa0\x69\xe5\x25\xc8\x3d\x82\xa4\x80\xae\xd5\x69\xa0\x96\xc6\x74\xde\x1c\x3f\x65\x9d\x46\x45\xfa\xf8\xe8\xb1\x75\x4c\xe2\xfc\x5e\xb9\x9d\x45\xdf\xcb\xb3\x87\x8f\x4f\xd9\xc3\x1f\x77\x77\xf3\x03\x16\x1a\x5c\xe1\x06\x72\x34\xdc\x8b\xff\xfc\x2b\x06\x3a\x46\x1f\x3e\x74\x30\x69\x85\x18\x57\x91\x5d\xdd\x06\xcc\x54\x59\xc4\xf4\x8b\x23\xfa\xfa\xcb\x9c\x7d\x66\x67\xf3\x18\x1a\xdc\x99\xef\xd3\x17\x4f\x70\xfa\xe6\x86\x0c\xb2\x38\x7b\x42\xaf\x09\x9e\x22\x3c\xe2\x9b\x9f\xaf\x36\xa9\xbc\x52\x63\xd0\x5f\xcb\xaf\x79\xb4\xe2\x9c\xf1\x08\x85\x90\xb3\x83\xb5\x67\xe5\x06\x74\x52\x3b\xaf\x80\xd9\x15\x04\xe1\xb6\xe3\xed\x8b\xac\x39\x2b\x98\x07\x33\x67\x25\x72\x22\xb3\xce\x48\xbd\x84\x5b\x9d\xaf\xc0\xd2\x0b\x1c\xbc\xf2\x8a\x57\x27\x75\xd7\x2e\x86\x67\x10\xf0\x2a\xb2\x7d\x59\x40\x47\x55\x0e\x75\xdb\x82\x1d\x43\x92\x1b\x57\x45\xc5\x3f\x73\x28\x36\x40\xe3\x1d\x1a\xb9\xf0\xd4\xc6\x78\xdb\xe5\x86\x0a\xd5\xdb\x3c\xd1\xc0\x20\xad\x88\x63\x72\x28\x9d\x31\x6e\x87\x7e\x25\x22\x5b\x45\x14\xf2\xcd\xe9\x20\x4e\xc0\xa8\xd6\xd3\x33\x08\x2e\x55\x09\x31\xd2\xa7\xc0\xa3\x78\xa1\xa8\x26\xde\xaa\x8e\xa1\xc2\x48\x2d\x67\x0c\xe4\xce\x83\xb8\xd8\x80\x9c\x8c\x21\x5b\xa9\x99\x9f\x64\xdf\x8e\xbe\x96\x9d\x53\x25\x14\xe2\xbc\xf2\xf8\xa5\x23\x8f\xa1\x7a\x14\x5a\xc8\xcd\xc4\xb2\x89\xb4\x68\x3b\x63\xfa\x57\x90\xdf\x28\xb6\xd7\xc8\xfc\x34\x77\x7f\xb3\x9b\x4f\x71\x7b\x83\xbe\x4a\xc9\xb5\x87\x52\xa2\x36\x72\x38\xb7\x14\x04\x37\xa4\x64\x6a\x66\x46\x2f\xee\xa9\x87\x44\x71\x0d\xaf\x74\xa0\xfe\x4c\x24\x00\xad\xe9\x50\x5a\x53\x7c\xd1\xe0\x42\x54\xd4\x60\x83\x15\x65\xe0\xc4\x49\xbc\x28\x22\x8f\xcf\x84\x3b\xb5\xfe\x6a\x03\x64\x05\xc8\xa2\x57\x05\x58\xd5\x38\x4d\xe5\x3e\x32\xeb\x34\x63\xbd\xef\xc6\x36\xc4\xf6\x54\x03\x9a\xe0\x1f\xdb\x17\x87\xb2\x40\xad\x0e\x11\x09\x9c\xb2\x3a\xf4\x3d\xdd\x2f\x84\x1e\x41\x8f\xa9\xb6\x02\x3c\x96\x69\xa0\xf7\x7f\x1a\x0a\x7e\x3e\x75\xdb\x1b\x51\x5f\x9e\x4a\x86\x3e\x07\xc6\xa4\xc6\x3d\x90\xd2\xb8\x07\xd2\x1a\xf7\x50\x52\xe3\xe1\xb6\x37\x34\x5e\x9e\xfa\xaa\x56\x18\x73\xde\x39\x83\x06\xd9\xf0\x49\xb3\x57\xe2\x22\x8a\x8e\x58\xaf\x05\xa9\xc7\x6e\x46\x65\xa9\xea\xce\x6e\xa7\x5e\x16\x7d\x7d\x28\xf5\x31\x75\x02\x1e\x6c\x09\xa5\x3b\x35\x17\x1f\x7e\x85\xe6\xb3\x82\x06\xae\x9a\x74\x1b\xcf\xae\x90\xf7\x3f\x12\xcd\x3c\xf3\xe3\x26\xa3\x44\x40\xce\x1d\x3b\xb9\xcb\x1b\x92\x74\x5f\x5b\xc6\x63\xe1\xb1\x59\x23\xfa\x3f\xb9\x6f\x3c\x4f\x6f\x6c\x83\x73\xb7\xc5\x93\x62\x19\x34\x8f\x05\x89\x11\xb0\x76\x0d\xb6\x61\x56\x1a\x7b\xb2\x46\xe5\x4a\x55\x38\xab\xbb\x42\xc6\xa4\xd6\x58\x42\x67\x44\xe5\x1e\x6c\x31\x56\x61\x03\x2c\xe8\x17\xc2\x93\xf9\xdf\x76\x5c\x7f\xcf\x1a\x59\x98\x40\xf8\x91\xa0\x22\x89\x1f\x1b\xf2\xde\xcd\xc6\x68\xe6\x3a\x7e\x7a\xb6\xd1\x93\x01\x5b\x75\x91\xe5\xa5\xf3\xdb\xa8\xe9\x87\x47\x56\x85\xeb\xac\x4c\xbc\x60\x51\x76\xce\x6f\x97\x62\xd7\xa2\x55\x87\xc5\x7d\x09\xb1\x80\xaf\xe0\x05\xfd\x1a\xe9\xf2\x10\x8d\x3c\x01\xed\x40\x8a\x3a\xf5\x0a\xbd\x44\x43\x2d\x74\xe2\x14\x59\x8a\xc7\xb4\x16\x7d\x43\xcc\x33\x3e\x82\x30\xa5\x2b\x8f\x3d\xa9\xf7\x4d\x3f\x9e\x6d\x7a\x90\xbf\x74\xc0\xf5\x09\x30\x1e\xdf\xe2\x3f\x16\x5c\x4b\x45\x6a\x56\x01\x5f\xd4\xf4\x3c\x1b\x16\x35\xf1\x6c\x62\x0e\xee\xa7\x02\x2d\xe3\xa8\xe5\x30\xcd\x47\x27\x6a\xe0\xc1\xa9\x0b\xe1\x8e\xb6\xb4\x10\x85\x14\x5e\x1e\x6b\xbd\xfb\x8c\x85\x2c\xc5\xda\xed\xac\x71\xa0\x79\xad\x8c\x12\x6c\x5a\x13\x11\x46\x45\x42\x95\x75\x7e\x8d\x08\x42\x33\xef\x2f\xe0\xd1\x4e\xf5\xc2\xae\xf3\xc5\x64\x5d\xcf\x09\x69\xea\x39\x40\xe3\xd6\xb1\x86\xd2\xb4\x14\x2f\x9e\x27\x2e\x8e\x4f\x7c\x15\xe7\x04\x16\xfc\xf7\x96\xd8\xca\x0f\x47\xd2\x6c\x75\xee\x5e\x5b\x93\x5f\xee\x21\x39\xb9\xf7\xb4\xe2\x72\xc7\x2d\x5a\x9d\xdc\x7e\xa3\x65\x93\x49\x50\x81\x6e\xc8\xce\xaa\xec\xbf\x5b\x8f\xa3\xed\x37\x4c\x28\x68\x85\x86\x38\xcc\x17\xdf\x8f\xf7\xf7\xb7\x4f\x57\x9b\x7f\x02\x00\x00\xff\xff\x4a\xf0\x02\x79\xcf\x14\x00\x00")

func _1558054487_initial_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1558054487_initial_schemaUpSql,
		"1558054487_initial_schema.up.sql",
	)
}

func _1558054487_initial_schemaUpSql() (*asset, error) {
	bytes, err := _1558054487_initial_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1558054487_initial_schema.up.sql", size: 5327, mode: os.FileMode(509), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _lockJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\xcb\x6e\xdc\x3a\x12\xdd\xfb\x2b\x1a\xbd\xce\x17\x78\x3b\xcb\x01\x82\xc1\x20\xb3\x0a\x02\xa2\x44\x95\xa4\x8a\x29\x52\xe1\xc3\x3d\xed\x20\xff\x7e\x21\xb5\x62\x77\xdb\x91\x6f\x80\xeb\xb4\x78\x64\x6e\x0c\xc1\x82\x80\x73\x9a\x64\xbd\x78\x58\xfc\x7e\xb3\xdb\xed\x3f\x51\x65\x38\xec\x6f\x77\x9f\x6f\x76\xbb\xdd\xee\xfb\xf4\x77\xb7\xdb\x7f\xa4\x9e\xf7\xb7\xbb\xbd\x84\x90\x38\xec\x3f\xfc\xfc\xff\xbf\x9c\x49\xbd\x7d\xfa\xe0\xfc\xa3\xcb\x0f\xeb\xc7\x8f\xa6\xff\x7f\x3a\x0e\xd3\xff\x03\x7b\x21\x73\xf9\xee\x3f\x5e\x7a\xf2\xc7\x7f\xf3\x71\x7f\xbb\x8b\x3e\xf1\xc5\xdb\xff\x72\xc3\x9e\xad\x1e\x3f\xb7\xc9\x98\x8b\x97\x1f\x5d\xfc\x98\x8c\xd9\xdf\xee\x1a\x32\xe1\xf2\xc3\xff\x59\xf9\x96\xf8\xe7\xab\xc7\x37\x3f\x3e\xbc\x8e\xdc\xa6\xbe\x62\xff\x6b\xf4\x95\xb4\x62\xe3\x2b\xe8\x5f\xa2\xb8\x36\xfc\x10\x29\xf2\xaf\xd1\x47\xfe\x7f\xde\xd8\x8d\xd3\x77\xbc\x30\x71\x2a\xe7\x0c\x93\xcd\x1a\x7f\x94\x68\x50\x7f\xfb\xca\xd5\x47\x50\xe8\xda\xf5\x3d\xdb\x18\x60\xd7\xac\x36\x2e\x70\xad\x28\x2e\x0c\x80\xf4\x1c\x22\xf5\x43\x7c\xc8\x9b\x86\x67\x8a\x1b\xe0\x91\x86\x7a\x13\x3c\xba\xd8\x9b\xe4\x0d\xe8\xaa\xb6\xae\x66\xb5\x14\x46\x64\x8f\xde\xf3\xe0\x82\x44\xe7\x8f\xca\x1d\xec\x52\x3c\xf1\xa7\x68\xbc\x08\xa2\xfe\x39\x0b\x3b\xfe\x0b\x92\x84\xa1\x8a\xcd\x82\x73\x18\xb1\x7f\xfe\x92\x33\xfa\x14\xd8\x2f\x2e\x83\x3f\xe7\xdb\xde\x14\xbe\x71\xad\x58\xcc\xc9\x43\x21\x48\x6b\x79\xd9\x12\xe5\x3f\x04\x8f\x14\x36\x30\x0c\x0b\xcb\xf8\x6b\x70\xb6\xca\x19\xff\x1c\xe1\x55\x47\xe0\x79\xf4\xc4\x01\x78\x22\xf5\x62\x38\x44\x67\x91\x17\xf4\x13\x87\x15\x12\xce\xdf\x26\x31\x3f\x7d\xb9\x39\xa3\xf4\xeb\xfa\x92\x7a\x91\xc0\xbd\x8b\x3a\x13\x74\x7c\x0b\x5c\x2e\xf0\x4c\x3a\x8a\xb3\xd7\xf6\x25\x25\xcf\x7e\xc6\x63\x2b\x79\x36\xa5\xd8\x39\xaf\x28\x04\xa7\x85\xc6\xa9\x05\xba\x32\xb0\x0b\x06\x25\x53\x5a\x15\xff\xc9\x95\xaf\xb3\x7b\xf2\x07\xea\x1d\xa5\x6a\x93\x43\x7c\xe8\x7c\x4b\x56\x1e\xe8\xd2\x5f\xbf\x87\xf0\x70\x05\x4b\x50\x42\xdb\x93\x3b\xbf\xa7\x48\x5e\xe1\x3a\x42\x6c\x37\x7e\x7d\x93\xf5\x96\x1b\x91\x03\x59\xd4\xbc\xa8\x32\xae\x05\x85\x6e\x9c\x46\x8e\xbb\xb9\x27\x41\x5d\xae\x35\x07\xed\x65\x00\xfe\xf5\x87\x54\x19\xd1\x6a\x8a\x9a\x60\x25\x04\x33\x89\x56\x02\xb0\x0e\xa2\x71\xc6\xb8\x03\x7b\x74\x06\x62\x17\x0c\x29\x00\x83\x52\x5c\xca\x8b\x47\x74\x91\x8c\x1a\xbc\xdc\x53\x64\x70\x23\x35\xe6\xd4\xf5\x46\xb8\xfc\x64\x81\x6d\x71\x6b\x09\x77\x2a\x05\x6a\x17\x42\x6e\x00\x0a\xda\x19\x43\x95\xf3\x14\x1d\xb0\xe3\xa8\xc4\x18\xb1\xad\x42\x0e\x05\xe3\x08\x14\x14\xfa\xc1\xa9\x86\x74\x74\x5e\x79\xfe\x96\xc4\x73\xcf\x36\x2a\xb6\x54\x99\xac\x15\xc9\xf3\xd3\xeb\x55\xb4\x21\x19\x33\xf1\xe2\xf0\xde\x36\x59\x8b\x98\x7f\x2d\xec\x45\x0c\xbf\x0a\xf4\x12\xbf\xe7\xc5\x63\x23\xda\xfe\x9e\x7d\xbb\x01\x1a\xb5\xa7\x66\x81\x42\x16\xee\xfc\xb7\x46\x01\x1c\xff\x18\x52\x6d\x80\x82\x42\x76\xce\x13\x8b\x49\x77\x27\x51\x85\x8e\x40\x69\xe0\x9f\xfc\x9a\x46\x00\x17\x3f\xd5\xb5\xbc\xa2\xa5\x03\x60\x50\xb3\x61\x6c\x06\xba\x23\x3b\xfa\xe6\x46\xcc\x92\x3e\x1e\x80\x05\xf6\xa6\xb5\xe7\x7b\xe1\xc3\x4b\x21\x33\xda\x28\xf4\x24\x36\x92\x58\xf6\x4a\x93\x55\xbd\xab\xa5\x59\xc8\x84\x10\x9c\xf5\x66\x94\x99\xd8\x8a\x9a\x22\xab\xcb\x86\x44\x39\x0c\xb9\x36\x7c\x60\x89\xef\x5c\x04\x80\x3e\xc5\xf6\xc4\x01\x78\x20\xca\xa9\xd4\xac\x38\xa0\x9e\x4a\x9d\xb7\x84\xb8\x56\xa7\xf8\x75\x51\x79\x93\x3d\x93\x72\xac\x33\x13\x12\x1d\x53\x7d\xfd\x82\xd2\x9b\xa2\xf7\xdc\x00\xa3\x9f\x02\x3c\x60\xfc\x63\x8c\x04\x0c\x7f\x1b\xa9\xce\x73\x2a\xb8\xf9\x4e\x45\x81\x71\xed\xd1\x84\x1e\xd6\x1e\x4d\xe8\x81\xed\xd1\x84\x1f\xd7\x1e\xcd\x93\x67\x0b\xf6\xe8\x39\x95\x8c\xed\xd1\xfc\xf4\xfb\x82\xac\xd2\xfd\xe2\x6a\x83\xf8\x66\xe8\xc5\x8e\xb3\xd1\x1c\x55\x74\xb0\xa5\x7e\x60\x8d\xd3\x40\xb1\x03\x85\x5e\x4b\xd3\xa8\x2e\xd9\x3b\x50\xfc\x17\x96\x6b\xde\xf5\xba\x7e\xca\xfd\x76\x74\x46\x8f\xb2\xb8\x2d\x04\x40\xc0\xf9\x11\x23\x19\x05\xcf\x64\x16\xa3\xc0\xfa\x84\xc7\x91\x40\x27\x52\xfa\x23\x65\xc2\x63\x2b\x12\xd8\xcd\xec\xc2\x63\x6b\x54\xca\xe6\xe9\xaa\xf8\x2f\x22\xa7\xd2\x26\x29\x2b\x16\x9b\xaa\x27\x9c\xa2\xf2\x52\x4e\xc0\xb1\xcc\xc0\xa9\x78\x48\x55\x2f\x71\x0b\x51\x0a\x7a\xd8\x8e\x1d\x9c\x20\x1f\xa6\x28\x81\x55\x09\xac\x4a\x60\xb5\xc9\xc0\xea\x91\x8a\xbc\xb7\x5b\xf0\xa0\x03\x2a\xe0\x1e\x82\xcd\x68\xcd\x80\xf1\xe3\xf7\x83\xeb\x5c\xcf\xc3\x62\x47\x9c\xec\xe1\xeb\x71\xe5\xba\x46\x69\x67\xeb\xa4\x17\x82\xf2\xfc\x8b\xb8\x35\x37\x94\x4c\x54\x95\x27\xab\x51\x77\xfb\x7a\x0a\x91\x3d\x36\x87\xad\x54\xd3\x87\x14\xba\x0d\xd0\xd8\xca\xa6\x00\x76\xba\xaa\x8d\xb3\x0c\xdc\x22\xba\x95\x08\x8c\xbe\x17\xef\x1d\x72\x87\xee\x10\x3a\x60\xf0\xf7\x16\x17\xbc\x21\xdb\x26\xdc\xe0\xae\x71\x7e\x41\x39\x84\x70\x1c\x7b\x44\x1f\x94\x76\xe9\x79\x15\x06\x48\xa6\x62\x39\x1e\x9c\xbf\x03\x67\xe1\x06\xb6\xea\x74\xb9\x3f\x38\x93\x10\xc9\xb7\xf4\xc0\x1e\x9e\x48\xaa\xc6\xcc\xb9\xc2\x67\x72\xa0\xa8\x3b\x7c\x1a\x41\x1e\x70\xbb\xe2\x52\x8a\x4e\x89\x15\xe0\x66\x67\x03\xfb\x5e\x42\x40\x56\xa2\x91\x31\xee\xa0\x3c\x4f\xe7\x17\xa6\x93\xe7\xb8\xc3\x71\xe2\x12\xbe\x25\x0a\xdd\x36\xb8\x9c\x37\x41\xc3\xe5\x12\xdd\x20\x1a\xb6\xa3\x07\x79\xdd\xc9\x3d\x72\x4f\xc3\x5a\x42\xee\x4d\xa2\xff\x2e\x27\x12\xcd\x36\x2c\x2c\xe7\xfc\x8d\xec\xdc\x89\x1f\xf7\xf7\xef\x28\xcc\xb1\x38\x36\x87\x83\xdc\x09\x36\x83\x81\x5a\xf4\x41\x18\xbc\xfb\xca\x7a\xb1\x05\x1d\x08\x8b\xda\x1d\xac\x71\x54\x03\xd3\x98\xad\xaa\x8a\xdc\x0f\x06\x57\xdc\xd4\x4a\x94\xd6\x3a\x0f\x4f\x24\x32\xf5\xd0\x47\xd5\xc8\xf3\x52\x36\x9d\xbf\x93\x0e\x2e\x79\x0d\x1b\x62\x4c\x62\x2c\x60\x89\xdf\x09\xff\xf5\xef\x67\x79\x5b\xfc\xc0\x1a\xc5\xf3\x7b\x8c\x91\xe7\xd1\x39\x0d\x74\x6d\x5f\x0a\xe7\xfd\xce\xde\x83\xa8\xaf\xdc\x29\xbd\x16\xfa\x72\xa7\xf4\xba\x41\xac\x9f\x07\x00\x76\x02\x01\x2b\x42\xcb\xad\xd8\x2b\x41\x2f\xb7\x62\xaf\x86\xbd\x13\x0f\x7e\xcf\x4d\x25\x0b\xad\x8d\xb2\xff\xed\xcb\x7d\xde\xb9\x90\x28\xf7\x79\xaf\xcf\x60\x2b\xf2\xed\xad\xe8\x9e\x43\x0a\x03\xdb\x7a\x03\x4c\x80\xaf\xfb\x0d\x12\x59\x51\xdd\x2f\x65\xc3\x08\x0e\xba\x5c\x0c\x9f\x27\x97\x72\x31\x7c\x26\x14\xb6\x71\x31\xfc\xd9\xcd\xe4\x94\x62\xc7\x36\xca\x6b\x59\x5d\x16\x96\x6b\x7e\x9a\xcb\xaf\x37\xe3\xd3\x8f\xbf\x02\x00\x00\xff\xff\xd8\x73\xac\x4e\x6f\xac\x00\x00")

func lockJsonBytes() ([]byte, error) {
	return bindataRead(
		_lockJson,
		"lock.json",
	)
}

func lockJson() (*asset, error) {
	bytes, err := lockJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lock.json", size: 44143, mode: os.FileMode(509), modTime: time.Unix(1, 0)}
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
	"1558054487_initial_schema.down.sql": _1558054487_initial_schemaDownSql,
	"1558054487_initial_schema.up.sql": _1558054487_initial_schemaUpSql,
	"lock.json": lockJson,
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
	"1558054487_initial_schema.down.sql": &bintree{_1558054487_initial_schemaDownSql, map[string]*bintree{}},
	"1558054487_initial_schema.up.sql": &bintree{_1558054487_initial_schemaUpSql, map[string]*bintree{}},
	"lock.json": &bintree{lockJson, map[string]*bintree{}},
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
