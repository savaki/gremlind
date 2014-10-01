package manifest

import (
	"github.com/hashicorp/hcl"
	"io/ioutil"
)

func ReadBytes(data []byte) (*Manifest, error) {
	return Read(string(data))
}

func ReadFile(filename string) (*Manifest, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return ReadBytes(data)
}

func Read(contents string) (*Manifest, error) {
	var m Manifest
	err := hcl.Decode(&m, contents)
	return &m, err
}
