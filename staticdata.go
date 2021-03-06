package main

import (
	"fmt"
	"io/ioutil"
)

//go:generate go get github.com/mjibson/esc
//go:generate esc -o data.go -pkg main data

func readFile(fileName string) (f []byte, err error) {
	useExternalData := opts.Embed
	fullPath := fmt.Sprintf("/data/%s", fileName)
	file, err := FS(useExternalData).Open(fullPath)
	if err != nil {
		return nil, fmt.Errorf(errStaticDataNotFound, fullPath)
	}
	defer func() {
		// overwrites named return value
		err = file.Close()
	}()

	return ioutil.ReadAll(file)
}
