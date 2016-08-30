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
		size:    7359,
		modtime: 1472572816,
		compressed: `
H4sIAAAJbogA/+RYQXPrJhA+y7+iQ9+tHiXv0EuPTQ99PbzXSS+ZySQeLK1iUgEKIMdqJ/+9CEmAJeJa
chI7L55MNCzssrvfLgv77yxCnMG3DP3yw/UsivQ4itAnATUB/XiWQkYYUYQzeaaAFjlWcCHXSC97mu+5
/q+HfNT6K9qsn0U3mgV5SzRPLQD5mjSkCBWCFyAUgW6Vt85SIgSspK2tZpw0xujfjfkYLWs6L5ny+FRV
1GIQYQruQKB5S6daN2pEnlsS3rSkz+fm1xOcl5TJgGgsBK6sYKKV91Y957RGXGvCk6fVl5b/89buBa5y
jlNv94DfNHUFONVWOopTc8l5Drjb0u4ZIQkFFljxMJtUgrA7yzXzPuZ/IwcJeCiJgNRi5EDsOXDWYmYY
kR9sB4aEbOL1uw+JF/R6nbIHen1DP4bXRyZiIUCpalHo5FHj0tHnXCi8lAewywInIBc9LAJovFZ6t97t
YswvWW1R+R9g/mz8nYNsRHeW7sN7wXMufofNBEYm1QSuohrPpJ1K1AUW6deSLjskRgj4TWPwFVOYxPgt
+5UItRrN+wcnbDTTpY608Vhc8scvLLWMofDyY2TXacZqNw2PjbbCdalfYKVA1DGLbq9vz25++rSlc8f2
3HmI01RoReKEM0UY6LybD6d0OooqMEFUgFqsdNoMyVK57NsiC4DApv+QwhGTOjO8oQnCRaKjMDbmuZlS
CGBJNaRoK1JvYaqViVNcxTxbPAL8HZ6h2imr3lSIVgH21KtxizMipArRYopzCE5kEJjK8UBMTQpJMfSg
kKzM8wApqEpNb4WE6uS+BWXjSrUZh3OmlXZli/KYI93khx/loWSzh+rRM80EcbzqjoaJfpUrLkZW6NOB
oi5TJ4CDrZYTMVjjvIR9Xh8vBkFn+sDzVpste8Zh1bM+DJ2+K5wAct2NZSJwmeD0ZHAzyrwybP3b2vEh
9Eo3826QEwHd3m24X32hoPrEnbtxSmTC1+5FV9OorpwgaqV86ppI3AHa6neKx6q9UR8d3P6lKvb037o+
NRMftAz6D5nTgEzDtXTvqomI4LtwSew/2U8OD/M4PDoQ9/aJOmhMvYf+0aim7JuUV6fR4TV2V7OoaRMc
PX6E61ZMTGDhG7IXfDmhZHefzlbTQN808junP5+/aXAEvDUpMHZdvmwnqIuNzjl8eQ9J+6g/YsDwxwVx
nap3+8z3ZHeO3jIVaTcJsixd6xdyoK7RhRJdA7HLdGrm7B4z/ff0XwAAAP//jOY9Kb8cAAA=
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
