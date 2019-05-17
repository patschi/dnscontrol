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
		size:    19917,
		modtime: 0,
		compressed: `
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
`,
	},

	"/": {
		isDir: true,
		local: "pkg/js",
	},
}
