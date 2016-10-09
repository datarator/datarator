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
		size:    42596,
		modtime: 1475997723,
		compressed: `
H4sIAAAJbogA/+ydXXPTuBrH7/spPIa7w0nKxbnp3TkBBs5QYOnOLgtTMqqtNCq2ZWS5bdjpd9+R/FKl
lmwFItetHl219b/So0f6/SMrfvn7IAiCIKQZfr8Kj4Iv1e9NCZ8yLP4cPpnHeEUywgnNijnHaZ4gjhfF
Zdjqb561P+5Uycn35Ncr+ZSqlcifTquqQkUeHim1hmovju42xwlPsGhvcfJH+Gz7WM5ojhknuOj831bF
2qNSgbMyFdnWHpWKSE2tWk6f6avkm1zGW3BGsvPQoKJ5NxEd0ZrEMc7Co4CzEmtlN52/3nQbDCNaZtyc
hCZiknF8jpkp5JRkJJXpOjQp0HWteH4oi2V0SZn2pKKNDzGGNqboCMfpQDr1M7dqXj/GmnCDOhVv6vae
W/UxR5uEotjcx4Gp3OrWGMWY9WoCNWdnlCYYGbrX00VZSYFzxBCnO7RXz3pzc7bT+KAnyDAh2bei65Ly
2Loe5TXn+dF8HiNe9WHGMIr5Gsc0KmaEznE2F95Q8PlFQbMlysmy8YvZmqfJk+a3f3cs4OYO+yHD30vC
cKy1klsb6pn8W0dOdTYcqj5ttsmT397eh00W6ofHVl/AJjsKsMlBwkclSixazER9Or4Xoq5TIKqSAVG3
Hb23hUfOMOebZc5Iz4i26r0sP9Qmlxyd9UfoqN0iRxEulv0zudN+M6Mdr4NGccl6nnYM0nCq2h7vm/X/
jWOGi2JBM04ynPFuojQD9CvtlBlnG9etEO66iQ9rmmHHbZxwMWdct8Ewdj3qn0nuqoUFTaiG7j1W/hpf
u6s/K5zlfkFzZxAsGI4JXyAWvyvTM5297rud34Wfu2qlZAxnkbts1fUvaOysDy8Qxy/Q5v3qT4y/jdLI
O5S6780xzfjaZSvjNOA6V39h5IxBUf/71f8Ic5en/1OiWSHup26R+leEufPZtoFXOEWJs2Fumzl23UiZ
aM5391v/CKkqk8R1pt4it7NK1O8+U6IVl5n6iM/draA+0qs3Wayt3/qEqnMKZN6BQpV0FrXaHfejMvFB
MLjN0b9xlCPOMRPngeHXL1/np/96apXdpvaf3gfrdl6rfUC7Ynvbk7nW7Ft2RPp5XMfwybTJaIi9rRWn
Od98wCzSzdzdYnipVmUOxjbhBz1dGNgwkZiYpvDuYFd7DjZYV0ovoTZszASAtLEA0mogYyJNuBXPQuYj
zNr9zwBINhYgWQ1kRJKrffRhlHOp85DlXP9FQwAwGwvArAYyIszVF1bDMBe8+2WoFzAX+m/0AoDZWABm
NZBRYZbfDNvQLIVe4qz/8jwAno0FeFYDGZHnzyS3gPkHyX0k+YfuCpUAMDYWwFgNZAyMqyudzARH8rhH
7Eb6S78CoNZYgFo1kNGofY2vh8CdrfG1d/DKTgPAta5YUzbWFf1gFtM0i6zoO+OO5HGvXEJ7gXTgqUNc
oqQ057OVWd8CDAZRBaMfW8PU6TWCVlUNlX7mjuc8PZNxB1PK+75rj8RhryxJd1tF4KkjrRhNwZDsY7g/
Q5Ij9Sj86O69Rz3eJKXLCLF4llVin5yq23vwre2+DaA8mONWidLO6btWF5MiopfGO+y3tCkqOGZi+GzU
l6RAZkcJtJw3BXy3DubhnClu3xhpZ4EaC/bGALvpbYqP9gdQTxLq5j7kHpwbiU8YG2/PDgBfYwF81UDG
xFfe5j+M8CwSOg85rjoOMNc6gHmKMG8/T8NMc4w4nsVoM6Or5ZWQegR0p+/AdK0DpifP9DsdeGauZ5oQ
vIK7SgAQXuvgYo1dY3i0blI9QsrOSlKp9dRGUv2ztgJPLQSwnirWVkT7iTIwvKUDhifNsM0iX85oP5f3
SteB6FoHC/tdY3iM7iEfqDpgHBuh8c0yNtonzQaemgUAPFWAmycWDzBMV8szqfON47bjwHKtQ+c73Fox
/FoR+NivgnkwriEfQm62iwtx2COXuNA+kz14WOYAL5ba5czHwUsmwQWrYHaaAnY3dNwOl57S8Sy2Bzxr
9719TYPZgkUQs5UUeWTESq8fuh3DedejXkHdfROKDcmzVaX1Euim88B1rQOuJ831sTXVHjMNRKs6IHqy
RJdJ38u2qxktNN5hrH0BWgAAGwsArAYyJsCWK+0ySbxdaCt9B6hrHUA9ZaitltliWvsLNOCs6gDnqeIs
XxE7gHKCPNzQTrTvzg0AYGMBgNVAxgTYbpEtZrSvi2y17wB1rQOopwy1zSJbTmt/gQacVR3gPEWcP+Lz
3pcrMHncI3yrDgO1tY5p50dHto8L2hKSkh3uHGouITbXKOUpyUgqB/1wSImua+V/Dn+qA+Bwv3ZFXh96
41lnD032rkqv3mRxv7HSqyWRmjuhN/Obnl3gyKuNjduUgP3WOrCUaS2alKR20DZO8BBxzshZ2Xlnc/Vv
CU5FXnRXAceII/3lwanpfzQvzjg1OIwWmDDGK1QmYsjb0PpS8bJvmrRN6Z8eb/xsVj6Lnx/efhhXI39z
cPNPAAAA///FaAX3ZKYAAA==
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
