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

	"/data/json_schema.1.json": {
		local:   "data/json_schema.1.json",
		size:    1225,
		modtime: 1470946880,
		compressed: `
H4sIAAAJbogA/8xUy07rMBBd118R+d5lpbZbtrBhxQegLkwzKUaxHewJtKr87/iZhxMJIbEgUh8+c2bO
zPEoN7KhyLEFelfRB4ZMM1Sabj187QKqXt7ghAHqtOpAIwfjAi7VkUB0LUNI5zHLoOby7LMcCLIXDnz2
hw09mY+Ib6h5b6n/d3Rf1mP0pHqJi2pcIpxBp3I1NKxvPeuwj4jgkougkQF2ScBhH54oQKJE2wtpFiJM
a3ZNEtwNNjAmFsXc3P5/DY1H/+1cS64F5EqaXeJ4iiXh40crZFdU56Kj5n1KjJ19pxnlSHAph7NmYpWD
DzccUCXhqRluKxowc0AaTAasbERAJRMwOa9vxbxsSEmwzfGc95NCgTHieVeqsvVyLeNqBkY+H7eEqi45
eCPV5KGvvK7BWYm6h0nE5hHI5Mduq4WNWn0+yhouf9jJssXCzZXw0tCB9OuextdGWnhivwIAAP//FxHp
PskEAAA=
`,
	},

	"/data/json_schema.2.json": {
		local:   "data/json_schema.2.json",
		size:    1001,
		modtime: 1470987754,
		compressed: `
H4sIAAAJbogA/8RSPU/DMBCd01+BDkaklpUVFiYkVsRgmmtl5I/0fCmtUP87tmM7TgkIsVCpjv187967
O38sGmDJCuH2Au4FCxJsCa4DfOwial/fcM0R6sh2SCzR+QtP9UGoOyUY07kBNL32h+dwaGDt9oEYtm6n
IOxe/HIKGKxtb7gQs5w0jFukgQZaGqljxlUCxCEBN6v4q9KpXhv3JaEgEseUTnq/JaKBK8JNCLlctrjx
SiytccshUTR7WsR/EIAqJFefIrOiNfi4KcUPGj+L3PktR6WhiN+xnuz7g2nxkIjnXR0Tn7diHGUzM0yP
GaHLKCueY5JmO/GZ73LoZO7RRqksmhuJnTgqK9qKO+PEo3uh+krhGzup/vyJ66AEhLteErbjc4zl1f7n
Wle6+4/do8mEcwP/WJt/v4vTZwAAAP//rn3K5ukDAAA=
`,
	},

	"/data/json_schema.json": {
		local:   "data/json_schema.json",
		size:    1471,
		modtime: 1470988446,
		compressed: `
H4sIAAAJbogA/8RUTU8CMRA97/4KUz0awKtXT14w8Wo4VHaW1PRjabsIMfx3Z7rtfkA1Gg6QsLSv8968
mQ77VRbMHxpgjzfMvH/A2rN7hBprGrBegMMDjMEgUI3kHuK+YKBbhZs32hRs7XZEpKXbSkarFT6OhLG1
abXviSmd0B42YDsaU0ILFRQXEeD7CDwswmckJ1ul3Zkgt5YfopxAv31Ewe4s1BRyO6+gxkxeGO3mnVAw
eyzDlxIwC9tWWKhidUPpZbGi85FE6k5USo6Mhpe6b07n4XcTT7j0wUlX5N9YS67g36RX8/msK9hH4ulV
DW5O+zvMR5GZEMQ02Um7gee8FXoz8ZnOUuhkmIKNvh3B3EBs+EEaXo24GSeI7rhsRxl+sBPrTz/h2WU6
nYJU3th/rnXLUQ+u0TlSmtXCuphuis0Ul5A9qCFzJPmZDEE5lYBnRepWygyUtUJ4FJlc/uX30k/9Fe/G
Tv55F9aGL6vy+B0AAP//nyA7wb8FAAA=
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
