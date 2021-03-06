// Code generated by go-bindata.
// sources:
// cmd/templates/app.yml.tpl
// cmd/templates/d-docker/Dockerfile.tpl
// cmd/templates/d-docker/docker-compose.yml.tpl
// cmd/templates/web/.gitignore
// cmd/templates/web/conf/config.json
// cmd/templates/web/conf/configPrd.json
// cmd/templates/web/deploy/start.sh
// cmd/templates/web/deploy/stop.sh
// cmd/templates/web/deploy/validate.sh
// cmd/templates/web/handler/handler.go
// cmd/templates/web/main.go
// cmd/templates/web/schema/db.sql
// DO NOT EDIT!

package templates

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

var _cmdTemplatesAppYmlTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x2c\x28\xf0\xcc\x4b\xcb\xb7\xe2\xe5\x52\x50\x50\x50\xf0\x4b\xcc\x4d\xb5\x52\xa8\xae\xd6\x73\x2c\x28\x00\xb1\x6b\x6b\x21\xe2\x21\x95\x05\x70\x71\x10\x1b\x26\xee\x92\x5a\x90\x93\x5f\xe9\x9b\x5a\x92\x91\x9f\x02\x93\x47\x16\xab\xad\x05\x04\x00\x00\xff\xff\xb0\xe1\x2a\x25\x60\x00\x00\x00")

func cmdTemplatesAppYmlTplBytes() ([]byte, error) {
	return bindataRead(
		_cmdTemplatesAppYmlTpl,
		"cmd/templates/app.yml.tpl",
	)
}

func cmdTemplatesAppYmlTpl() (*asset, error) {
	bytes, err := cmdTemplatesAppYmlTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/templates/app.yml.tpl", size: 96, mode: os.FileMode(504), modTime: time.Unix(1504170429, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmdTemplatesDDockerDockerfileTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xc8\x34\xb6\x30\xe3\xe5\x72\x0b\xf2\xf7\x55\xa8\xae\xd6\x73\x4a\x2c\x4e\xf5\xcc\x4d\x4c\x4f\xad\xad\xb5\x42\xe6\x86\x24\xa6\xd7\xd6\xf2\x72\x85\xfb\x07\x79\xbb\x78\x06\x29\xe8\xe7\x17\x94\xf0\x72\x39\xfb\x07\x44\x2a\xe8\xe9\x27\x95\x66\xe6\xa4\xe8\x57\x57\xeb\x39\x16\x14\xf8\x25\xe6\xa6\xd6\xd6\xea\x15\x95\xe6\x81\xd5\xa0\x88\xf2\x72\x39\xfb\xba\x28\x44\x2b\xe9\xa1\x88\x2a\xe9\x28\x28\xe9\x26\x83\x48\xb0\x86\xe4\xfc\xbc\xb4\xcc\x74\xbd\xac\xe2\xfc\x3c\xa5\x58\x40\x00\x00\x00\xff\xff\x59\x26\x5e\xdb\x9e\x00\x00\x00")

func cmdTemplatesDDockerDockerfileTplBytes() ([]byte, error) {
	return bindataRead(
		_cmdTemplatesDDockerDockerfileTpl,
		"cmd/templates/d-docker/Dockerfile.tpl",
	)
}

func cmdTemplatesDDockerDockerfileTpl() (*asset, error) {
	bytes, err := cmdTemplatesDDockerDockerfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/templates/d-docker/Dockerfile.tpl", size: 158, mode: os.FileMode(504), modTime: time.Unix(1504170429, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmdTemplatesDDockerDockerComposeYmlTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8d\x31\xce\xc2\x30\x0c\x85\xf7\x5f\xfa\xef\x60\x65\x6f\xcb\x84\x54\x6f\x5c\x80\x2b\x54\x51\x6b\x42\x00\xdb\x51\x12\x60\xa8\x7a\x77\x94\x64\xc9\x62\xd9\xef\xf3\x7b\x6f\xdf\xc7\x4b\x08\x57\xcb\x74\x1c\xf8\xff\x07\xf0\xd5\xf8\xf4\xe2\x96\xcd\x47\x84\x49\x43\x2e\xa2\x67\xeb\x08\xa1\xac\xab\x4a\xb6\x5e\x28\x2e\x62\x99\x10\xfa\x80\xc6\x99\xad\x6c\xed\x59\x28\x23\x98\xbb\xa6\x6c\xca\x19\x34\xe6\x54\x5b\x06\x30\xf3\x69\x3e\x63\x19\x15\x7d\xf4\xf5\x66\x4a\xcd\x36\x80\x29\xcd\x53\x9f\x3d\xad\x2a\x37\xef\xc6\x47\x52\xc1\x4a\x3b\xc1\xfc\x02\x00\x00\xff\xff\x3c\x91\xcf\xb5\xc7\x00\x00\x00")

func cmdTemplatesDDockerDockerComposeYmlTplBytes() ([]byte, error) {
	return bindataRead(
		_cmdTemplatesDDockerDockerComposeYmlTpl,
		"cmd/templates/d-docker/docker-compose.yml.tpl",
	)
}

func cmdTemplatesDDockerDockerComposeYmlTpl() (*asset, error) {
	bytes, err := cmdTemplatesDDockerDockerComposeYmlTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/templates/d-docker/docker-compose.yml.tpl", size: 199, mode: os.FileMode(504), modTime: time.Unix(1504170429, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmdTemplatesWebGitignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x2a\xcd\xcc\x49\xd1\xe7\xe5\xd2\xd2\x2b\x2a\xcd\x03\x51\xa9\x15\xa9\x20\xaa\xb8\xbc\x80\x97\x4b\xaf\xac\x38\x39\x3f\x25\x15\x10\x00\x00\xff\xff\x74\xf2\xfe\x5d\x24\x00\x00\x00")

func cmdTemplatesWebGitignoreBytes() ([]byte, error) {
	return bindataRead(
		_cmdTemplatesWebGitignore,
		"cmd/templates/web/.gitignore",
	)
}

func cmdTemplatesWebGitignore() (*asset, error) {
	bytes, err := cmdTemplatesWebGitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/templates/web/.gitignore", size: 36, mode: os.FileMode(504), modTime: time.Unix(1504170384, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmdTemplatesWebConfConfigJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\xe5\x52\x50\x50\x50\x50\x2a\x4e\x2d\x2a\x4b\x2d\x52\xb2\x52\x80\x0a\x80\x05\x33\x0b\x94\xac\x94\x0c\xf4\xc0\x50\x49\x07\x49\xa2\x20\xbf\xa8\x44\xc9\x4a\xc1\xc2\xc0\xc2\x00\x22\x5a\xcb\xcb\x55\x0b\x08\x00\x00\xff\xff\xdb\xb2\x79\xfa\x4b\x00\x00\x00")

func cmdTemplatesWebConfConfigJsonBytes() ([]byte, error) {
	return bindataRead(
		_cmdTemplatesWebConfConfigJson,
		"cmd/templates/web/conf/config.json",
	)
}

func cmdTemplatesWebConfConfigJson() (*asset, error) {
	bytes, err := cmdTemplatesWebConfConfigJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/templates/web/conf/config.json", size: 75, mode: os.FileMode(504), modTime: time.Unix(1504170384, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmdTemplatesWebConfConfigprdJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\xe5\x52\x50\x50\x50\x50\x2a\x4e\x2d\x2a\x4b\x2d\x52\xb2\x52\x80\x0a\x80\x05\x33\x0b\x94\xac\x94\x0c\xf4\xc0\x50\x49\x07\x49\xa2\x20\xbf\xa8\x44\xc9\x4a\xc1\xc2\xc0\xc2\x00\x22\x5a\xcb\xcb\x55\x0b\x08\x00\x00\xff\xff\xdb\xb2\x79\xfa\x4b\x00\x00\x00")

func cmdTemplatesWebConfConfigprdJsonBytes() ([]byte, error) {
	return bindataRead(
		_cmdTemplatesWebConfConfigprdJson,
		"cmd/templates/web/conf/configPrd.json",
	)
}

func cmdTemplatesWebConfConfigprdJson() (*asset, error) {
	bytes, err := cmdTemplatesWebConfConfigprdJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/templates/web/conf/configPrd.json", size: 75, mode: os.FileMode(504), modTime: time.Unix(1504170384, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmdTemplatesWebDeployStartSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func cmdTemplatesWebDeployStartShBytes() ([]byte, error) {
	return bindataRead(
		_cmdTemplatesWebDeployStartSh,
		"cmd/templates/web/deploy/start.sh",
	)
}

func cmdTemplatesWebDeployStartSh() (*asset, error) {
	bytes, err := cmdTemplatesWebDeployStartShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/templates/web/deploy/start.sh", size: 0, mode: os.FileMode(504), modTime: time.Unix(1504170384, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmdTemplatesWebDeployStopSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func cmdTemplatesWebDeployStopShBytes() ([]byte, error) {
	return bindataRead(
		_cmdTemplatesWebDeployStopSh,
		"cmd/templates/web/deploy/stop.sh",
	)
}

func cmdTemplatesWebDeployStopSh() (*asset, error) {
	bytes, err := cmdTemplatesWebDeployStopShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/templates/web/deploy/stop.sh", size: 0, mode: os.FileMode(504), modTime: time.Unix(1504170384, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmdTemplatesWebDeployValidateSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func cmdTemplatesWebDeployValidateShBytes() ([]byte, error) {
	return bindataRead(
		_cmdTemplatesWebDeployValidateSh,
		"cmd/templates/web/deploy/validate.sh",
	)
}

func cmdTemplatesWebDeployValidateSh() (*asset, error) {
	bytes, err := cmdTemplatesWebDeployValidateShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/templates/web/deploy/validate.sh", size: 0, mode: os.FileMode(504), modTime: time.Unix(1504170384, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmdTemplatesWebHandlerHandlerGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x51\x3f\xef\x9b\x30\x10\x9d\xb1\xe4\xef\x70\xf2\x64\x7e\x8a\xcc\x8e\x94\xb1\x6d\x96\x74\x88\x2a\x75\x88\x32\x38\xe4\x02\xa8\x60\xd3\xe3\x5c\x12\x45\xf9\xee\x95\xc1\x34\x69\xd5\x05\xc9\xef\x71\xf7\xfe\xdc\x60\xab\x1f\xb6\x46\x68\xac\xbb\x74\x48\x52\x48\xd1\xf6\x83\x27\x06\x2d\x45\xa6\x1c\x72\xd1\x30\x0f\x2a\x12\x99\xaa\x5b\x6e\xc2\xd9\x54\xbe\x2f\x6a\x4f\x6d\xd7\xd9\xa2\x0f\x37\x25\x45\x1e\xf9\x5f\x96\xe0\xe0\x03\x23\xc1\x16\xfa\x70\x33\x5f\x71\x5a\xde\x7a\xe6\xf9\x3e\x20\xec\x66\xa1\xcf\xc1\x55\x30\x32\x85\x8a\xe1\x21\x45\xb6\xa0\x04\xd7\xe0\x2a\x1d\x05\xcd\x01\xc7\xc1\xbb\x11\xbf\x53\xcb\x48\x1b\xf8\x48\xe8\xcf\x80\x23\xe7\x52\x64\x7b\xe4\xc6\x5f\x46\x38\x9e\x46\xa6\xd6\xd5\x52\x3c\x57\x13\x29\xcc\xde\x0e\xd1\x88\x1d\x8e\xcb\x1f\xa7\x97\x76\xd4\x54\x85\x2a\xe1\x6f\x68\xf5\x51\x42\x83\x5d\xe7\x37\x11\x4a\x3a\xe5\x1f\xa1\x87\xfa\xf2\xe9\x9b\x7a\x46\x32\x7e\x66\xd5\xe8\x7b\x99\xd1\x13\xfc\xd7\x3f\xfd\x93\x60\x8e\x3d\x99\x99\xd6\xc7\xd3\xf9\xce\xa8\xd5\x2e\x6e\x80\xc9\x53\x77\x51\x79\xfe\xbe\xfb\x80\x75\x3b\x32\x52\x32\xa8\x73\xf8\x88\x0d\xa7\xba\xe3\xae\xab\x27\x08\xd4\x6d\xd6\xf4\x50\x6e\x81\xac\x7b\xdd\x36\xd6\x31\x87\x5c\x86\xcc\x2b\xba\x7e\x9f\x4b\x38\xe5\x26\x25\xd7\x2b\x91\xde\xc6\x98\xd8\xff\x53\x8a\x8c\x90\x03\xb9\x74\xf4\xd9\xee\xef\x00\x00\x00\xff\xff\xc2\xc2\x0e\x8e\x50\x02\x00\x00")

func cmdTemplatesWebHandlerHandlerGoBytes() ([]byte, error) {
	return bindataRead(
		_cmdTemplatesWebHandlerHandlerGo,
		"cmd/templates/web/handler/handler.go",
	)
}

func cmdTemplatesWebHandlerHandlerGo() (*asset, error) {
	bytes, err := cmdTemplatesWebHandlerHandlerGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/templates/web/handler/handler.go", size: 592, mode: os.FileMode(504), modTime: time.Unix(1504170384, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmdTemplatesWebMainGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\x4f\x8b\xdb\x30\x10\xc5\xcf\x36\xf8\x3b\xa8\x3a\x49\x60\xd4\x4b\xe9\xa1\x25\x97\xa6\x4d\x03\x6d\x69\xf0\xe6\xbe\x08\x7b\xec\xd5\xae\xfe\x98\xb1\x1c\x1c\x96\x7c\xf7\x65\x64\x87\x8d\x6f\x4b\x0e\x86\x19\xf9\xe9\xf7\xf4\x66\x7a\x5d\xbf\xe8\x0e\x98\xd3\xc6\x17\x79\x91\x1b\xd7\x07\x8c\x4c\x14\x79\xc6\x5b\xab\x3b\x4e\xc5\xd4\xa2\x76\xf0\xb9\x76\xcd\xaa\x0d\xbe\x35\x2b\x81\x0d\xab\x76\x00\x3c\x01\xf2\x22\x97\x44\x3e\x69\x4c\x58\xba\xb6\x33\x16\xd8\x86\x91\x81\x7a\x88\x68\x7c\x27\x78\xcd\x4b\xc6\xe9\x9b\xb9\x23\xea\x68\x82\x67\xad\xb1\x50\xb2\xe7\x81\xca\x80\x4e\x47\x2e\x17\x60\x3b\xfa\x3a\xbd\x5b\x48\xf6\x4a\x60\xd7\xa8\x83\xc6\x01\xb6\xc1\x39\xed\x1b\x21\x97\xc3\x9f\xa3\xeb\x6f\xcf\x42\x4f\xe4\x92\x01\x22\xfb\xb6\x61\x24\xf9\x0d\xf1\xaa\xe0\x35\x39\x64\xa6\x4d\xff\x3f\x6d\x98\x37\x36\xf1\x33\x1b\x3a\xf5\xab\xaa\xfe\x57\x02\x10\x49\x93\xf5\xda\x9b\xfa\xda\x5d\xde\x6f\x11\x35\xa5\x50\x7f\x83\x6e\xb6\xa9\xdc\x61\x70\x94\x5b\xcc\xf6\xf2\xfb\x3d\xfc\x85\x3a\x27\xa2\x72\x1b\x7c\x04\x1f\x45\x9a\x48\xa6\x9b\x06\x4b\xf6\x78\xe3\x9f\x82\x51\xf5\xe3\xfc\x07\xce\x82\xcf\x3b\x51\xa6\x4f\x21\x69\xd7\x1f\xd2\x93\x90\xcf\x1e\x6e\x9c\x48\xff\xa4\x7d\x63\x01\x55\x05\x9d\x19\x22\xe0\x7e\xee\x85\x5c\x4d\x61\xb9\x5e\x8d\x7e\x7f\x3c\x1e\xfe\x8d\x93\xa0\x27\x2a\x31\xa4\xa5\xcb\x92\x19\x1f\x05\xc1\x95\x68\x6d\xd0\xf1\xeb\x17\x29\x4b\xe6\xc6\xe9\xae\xf1\x5c\x8a\xfc\x2d\x00\x00\xff\xff\x8a\xf9\xd5\x74\xd0\x02\x00\x00")

func cmdTemplatesWebMainGoBytes() ([]byte, error) {
	return bindataRead(
		_cmdTemplatesWebMainGo,
		"cmd/templates/web/main.go",
	)
}

func cmdTemplatesWebMainGo() (*asset, error) {
	bytes, err := cmdTemplatesWebMainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/templates/web/main.go", size: 720, mode: os.FileMode(504), modTime: time.Unix(1504170384, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmdTemplatesWebSchemaDbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func cmdTemplatesWebSchemaDbSqlBytes() ([]byte, error) {
	return bindataRead(
		_cmdTemplatesWebSchemaDbSql,
		"cmd/templates/web/schema/db.sql",
	)
}

func cmdTemplatesWebSchemaDbSql() (*asset, error) {
	bytes, err := cmdTemplatesWebSchemaDbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/templates/web/schema/db.sql", size: 0, mode: os.FileMode(504), modTime: time.Unix(1504170384, 0)}
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
	"cmd/templates/app.yml.tpl": cmdTemplatesAppYmlTpl,
	"cmd/templates/d-docker/Dockerfile.tpl": cmdTemplatesDDockerDockerfileTpl,
	"cmd/templates/d-docker/docker-compose.yml.tpl": cmdTemplatesDDockerDockerComposeYmlTpl,
	"cmd/templates/web/.gitignore": cmdTemplatesWebGitignore,
	"cmd/templates/web/conf/config.json": cmdTemplatesWebConfConfigJson,
	"cmd/templates/web/conf/configPrd.json": cmdTemplatesWebConfConfigprdJson,
	"cmd/templates/web/deploy/start.sh": cmdTemplatesWebDeployStartSh,
	"cmd/templates/web/deploy/stop.sh": cmdTemplatesWebDeployStopSh,
	"cmd/templates/web/deploy/validate.sh": cmdTemplatesWebDeployValidateSh,
	"cmd/templates/web/handler/handler.go": cmdTemplatesWebHandlerHandlerGo,
	"cmd/templates/web/main.go": cmdTemplatesWebMainGo,
	"cmd/templates/web/schema/db.sql": cmdTemplatesWebSchemaDbSql,
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
	"cmd": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"app.yml.tpl": &bintree{cmdTemplatesAppYmlTpl, map[string]*bintree{}},
			"d-docker": &bintree{nil, map[string]*bintree{
				"Dockerfile.tpl": &bintree{cmdTemplatesDDockerDockerfileTpl, map[string]*bintree{}},
				"docker-compose.yml.tpl": &bintree{cmdTemplatesDDockerDockerComposeYmlTpl, map[string]*bintree{}},
			}},
			"web": &bintree{nil, map[string]*bintree{
				".gitignore": &bintree{cmdTemplatesWebGitignore, map[string]*bintree{}},
				"conf": &bintree{nil, map[string]*bintree{
					"config.json": &bintree{cmdTemplatesWebConfConfigJson, map[string]*bintree{}},
					"configPrd.json": &bintree{cmdTemplatesWebConfConfigprdJson, map[string]*bintree{}},
				}},
				"deploy": &bintree{nil, map[string]*bintree{
					"start.sh": &bintree{cmdTemplatesWebDeployStartSh, map[string]*bintree{}},
					"stop.sh": &bintree{cmdTemplatesWebDeployStopSh, map[string]*bintree{}},
					"validate.sh": &bintree{cmdTemplatesWebDeployValidateSh, map[string]*bintree{}},
				}},
				"handler": &bintree{nil, map[string]*bintree{
					"handler.go": &bintree{cmdTemplatesWebHandlerHandlerGo, map[string]*bintree{}},
				}},
				"main.go": &bintree{cmdTemplatesWebMainGo, map[string]*bintree{}},
				"schema": &bintree{nil, map[string]*bintree{
					"db.sql": &bintree{cmdTemplatesWebSchemaDbSql, map[string]*bintree{}},
				}},
			}},
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

