// Code generated by "esc "; DO NOT EDIT.

package js

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/helpers.js": {
		local:   "pkg/js/helpers.js",
		size:    21112,
		modtime: 0,
		compressed: `
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
QN182kiw+H8BAAD///2Gg6R4UgAA
`,
	},

	"/": {
		isDir: true,
		local: "pkg/js",
	},
}
