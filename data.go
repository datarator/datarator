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
		size:    37531,
		modtime: 1474401743,
		compressed: `
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
`,
	},

	"/data/swagger.yml": {
		local:   "data/swagger.yml",
		size:    2521,
		modtime: 1474361246,
		compressed: `
H4sIAAAJbogA/+xWTW/cNhC981eMswXUos3KLoIedCvitHXhxG53A/Q6K85KTCiS5oz2w7++oL6sXTto
ix566Z6oGc7jm/dGKy5gfXd9Bx9/v10VagE5lzU1yPmhsfnW+3mIH2y+wTgPlbzLHx8f1UItYMdqAWoB
aC0YB94RtNHC18jg/B5MEyw15IT0N93+i4uLi3HzFRiGxkeCraWD2Vj6DhyRZhAPn4kCfGKfMO0Rvk1k
UzxSiF63JamF4j1WFcUCsu+Xl5kybusLBSBGLBVwjYIRxUf48f5GAWjiMpogxrt50jAgsKCQJWbQKAgV
OeqzeyM1/LJe38MGmTT8urr7MODtKHKH9epyebW8eqUArCnJMSUSAA4bKuD9zbp7aqMtIKtFAhd5HnG/
rIzU7aZliqV3Qk6WpW9yPRKbrRpkoZjf3rx992H1LlMLkJpA+waT5NvuiSnuTEmq9iwFTLXLmqL/3GII
CT0ZFSMeU1EyoXOUGKRGgaNvO6mA2xB8FFZDOrXzGhL1aZFc3xtrYUMQIm3NgXRyJ4EGlJpVkusepS4g
x2DUYNqAhSFYU2KyIk8eq64m5cYh6yUMqZluBSBY8bhOGP3GIcBt02A8FvBz7xwx/PH+tpNh2HHifpYN
UR/SbuPdjS4AtV7NQUvvuB36H099xrxPzNsbtwodJL0r56FDY78AyA9jJmDEhoTiCaJxBWhftumFmsLj
oL2QOOl5nYaka6+3e5hx493kZD/ibZjaSr9ID62JpAuQ2NIs0YMVswjAV5G2BWSLXNPWOJPQOe9FHSWP
xME7nmuVvbm8zIovEb9xO7RGg3GhTd39m6FY/fYfD8V8qy+F5DVLJGz+9/0vfJ8Bp4rVjIQcAxXgN5+o
7FWYqKtRwSTV9BBqL/5jtKwGlwJFMU/EjH7i1GMbJ1RRnKJbHxuULv7Dm3EwUKjy8fhU+6Iob4dtoyyd
iWfHsUTjqilIB0yf0WR1VZmxj6mL8+ruH36KHRpbPJuZsXaW2EcM4ZnbRqjhef0LBE/fwb/NQbD658e/
KOkaq1HN9B1vn7E50/NkygLJUJWuJN2XVHx8okCubeYMXgPu0FjcWDqJBnJ6fkj3b+StVmcC9M3fkyg6
CEWH9tqXHeMTWj8Zp8G30t+PcJOWq/66o+aXiSLPh1vQ0vhM/RkAAP//gb/LstkJAAA=
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
