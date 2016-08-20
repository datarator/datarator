package main

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

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/data/json_schema.json": {
		local:   "data/json_schema.json",
		size:    5111,
		modtime: 1471670775,
		compressed: `
H4sIAAAJbogA/+xXzW7bMAw+O08xaL0t6M9hl1176g4dMPRQoGgDJWY6FfrxZDlNsPXdJ8mSrNiqFztb
0w4LiiamRfIjP5KSfkwyJDh8WaJP724mWaafswwdSTAC9P4khyXhRBHByxMFrKBYwXm5QnrZ03TH9deM
2vWT7FaroGiJ1jEGUGy5FmWokKIAqQj4VdG6IMkQ8Io57PZ5UYPTn1v7ZVEauai4ivTUpjBmEOEK7kGi
qZMzjY1Zk6dBhNdOdHZqPy3DtGK8TJjGUuJNMEw0+GjVc0mrzbkQniJUF07/bMt7gTdU4Dzynsibln4D
nOsoG0kDcy4EBexdBp8ZKqHAEiuRViuVJPw+aE2iL/u/toMkfK+IhDxw1JDYSuDEcWYVUVw8e5bEuq6/
/yWxXRKFBKU2s0LTqIYVRqw5U3he7qFeFngB5azFRYKNv1VoLru+xuJh6Mbbb4g51z8ViinYSUtDJeoc
y/yyYnMf3ygDVyZTQ9U/C8IHK11iNtzTV03feriWeLzgeVBMcVYnvm84cIO324VudPlOKrBSIE0JoLub
u5+3H4620Hq1Z3echv/WgNm1E1eYVrDLkJ3GIy1en86k838d5l/oHh96p2kCmq14etvMJjlOVSv6JHXt
6j88ixbRbKEhHfOoI0dSuu2t609LdFRrj94856RciFWz72R2nyl1UAZULF2REntGHb4/WRxjOO8n+SrK
xuugOIAeS/BrSbKd4wdP7UPYTTpnrLdwFBp00n2RIdwg2n8Q9517LqPqOFz5GCfHSyL1Pjrtyo4ZppB8
sYTEK4o7ZowoZcXKk0aWFaUJURKKkTsj/8JEqc9rB68J2RwbRyZUxoGkkHY6mhJG+m8hYRdO3Aqz+F74
8fRF50UiW6NmRd+hLRzJfW345Ij5Ayxczx2wYMTjjDRXhjfbhZFtn+itUJFOkyTzqrnYAgUGXIV5n2OF
m+HP7LvgY6L/nn4FAAD//4/RxJb3EwAA
`,
	},

	"/": {
		isDir: true,
		local: "/",
	},

	"/data": {
		isDir: true,
		local: "/data",
	},
}
