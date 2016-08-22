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
		size:    5954,
		modtime: 1471896060,
		compressed: `
H4sIAAAJbogA/+xYQU/bMBQ+p79i8rgNFTjssiuXsQOT0A5ICCq3eQUjx85sp7Ta+O+zHdtxGtM1YaPt
tApR9cXv+Xvve/5s58coQ5zB1zn69O5mlGX6d5ahIwHGgN6f5DAnjCjCmTxRUJQUKziXC6SHPR9vOf66
oHb8KLvVLigaon1MABRHrk0ZKgUvQSgCflQ0LlgyBKwqHHb7e1aD059b+2VRGjuvmIr81Ko0YRBhCu5B
oGNnLzS2woY8DSa8dKazU/tZC0yrgslEaCwEXoXARIOPRr1UtDqcS+E5QnXh/M9as5d4RTnOo9kTddPW
B8C5zrKxNDCnnFPAfsowZ4YklFhgxdNuUgnC7oPXKPqy/+s4SMD3igjIA0cNiWsFHDnOrCOKm+eVLbGs
++9/S7RbohSg1GpSahpVv8aIPScKT+Ur3GWJZyAna1wk2Phbjeaq63ssFkMnb78h5pxT7uC5HLf2+gzL
AY5Mqv5eujJEnWORX1bFFAbgDQG+GWL6un/hhPV2usRF/5mudLf0r+oVf7pgeXBMtUjN8yYtYgZvd9E7
pfQLt8RKgTAdh+5u7n7efjhqofVuL25wTbut6dm2C3/ZSOqGurho10E8+yw9W4k4nxcLapbAftR0/ODp
H1hX+cBFTyXdHyqMqOwBD0HbBnKwwLSCbc4rf4wCn3qn8gFNK59+XK1ln6RuXdl3z6JFNJlpSGMW7TYD
KW3P1p1PW3RWS4/e/M6JnPFFc4TL7JFN6qQMqNi6IBJ7Rh2+vVyf7d13rygOoA99P7JnlJ2X9jGclDrX
lUO4VfS6NL6JCDeIXi/Em64Ql1F37K59zCTjORF6Hz3u2sYFppB8MIfEI4o7YYwpFcXak0HmFaUJUxKK
sbsg/4Ki1HeRnfeEaK5EAwsq4kRSSDsrmpKCbL7Qh1048YIli1+xfDx9U71IVGuQVmw6tIXrpu8NXxw+
fYSZW3M7bBj+NCHNdfhgV2EU2xe6lSrSZRJkWjXviIBCAUwFvc+xwo34F/ZZmGOk/55/BQAA//+OvKb6
QhcAAA==
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
