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
		size:    13814,
		modtime: 1474147522,
		compressed: `
H4sIAAAJbogA/+xazXLbNhC++ykwbG7t0MmhFx/rHpoekk56yUzG0UDkykJKAAwAylI7efcOQYmCxAUE
2iItW+LBY5sf9w/fLrAA/rsihJBECvg4S27Il+bvzZO8UVD/O/npOocZE8wwKfS1AV4W1MCtXiQt/scv
7a+9hPz9vXi6kM/cFWJ/u2tEJQ48uXGkJq4XN/vqSiVLUIaB7rzb+Rh9axEgKl5HFH1rEZkbPve56/zX
icv2c1kJ49dvVmVtXcKEgXtQSVeChXEmGLeWvvUh6HKNePfWPpHWFRUXePR27KNK0ZXPOmaA+2UQPzEa
9Xh4EXPJOhTv1/reRflY0lUhae738QCLWtwcaA4qiCFuzKZSFkA97gVctEI0lFRRI3vo00YxcR9Qh8e5
G8SrgJGJgu8VU5CjWbPNuADZdt7cYVUlccvOs2S9dusdbi9it+PqJetHzPrTIWw9xT0LYZf8QtiXQ9gh
p6lSgTGrSalYYERb9FEmK1flxNBp2MKB9OqSZqAnYSZ39G8Y/RpmzTVPO/XH0ze070Os/6uhZQFad2OE
jM0jVNzKQqo/YDmcfKHNcMLL1WCyFeTM3FKVf6j4FCPpcfT8Tg18oByGlP9x9htTZj6Uij8lQ2rIcWR/
gvvhyPlJPrwXOSo/OuPdHO29+BD1wB+cWNe9hWeyK6kxoOrKk3z98vX67uc3UeHaSH/0oofmuQKt00wK
wwQI4zFwD1wJo3yrhF0oM1G4ci4FVrw7QG3wKo8AFUCUO/+yMgTL6uIaBNgiM8moylM7HiFspRSILBiT
DSbNZB4UllMDaU5XqZxNHgD+icVyKcz8IDgOtQIaDE6dHOmMKR0cii0qnQGnRdBtBxwFLWiE8hoUq9ti
46ysiiIKFO12jbXIJ7QqR1soL5FerQPCa/fahs++pouMtHC0xdtXWKPnj3YBdlaTh62M6Rybesn4ZNRz
qcZq1y7ET9zO4MxYjzZDZHzGL2hR+V1pYdG72ydHeM/IB4ndoproxI7TUJkUGP8eSVauzi3HsD0BMn6K
zZTklwzzWGOD8yoSbH+b6LySzelfhWebjIyfekG34t1rkZTDMtBftbic6UwuvEc0O1hOtQFVRy4GvWCa
+jOUoCHePKdXOk5wNdpuw55V/u5vA6WeYO5+YLd3GugpZPule3umfNkcK5xfysjZZIqfqJDx+U/ve3Ry
h49cL+xvnhD77YnXWdH+G3rGR3rccbncIomf0Qa4f3hyaf2k/nEboRNsInvfGmkOuc+qnij8XJ+MP30q
NPgd2DESrWCcPeJ6VLg/PHzxb4tsLwD+2r37F+PAC6wU/Yg2YpXovdXU3lbpFIoNYeT0G2T7B9avu4zI
hwnDr/CQy6F289eo62PHjw5NvcOZUGMUm1b4JUwogOPXi5Isp4biUzD3fYMc7jiuXDU/f/wfAAD//yBO
W3f2NQAA
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
