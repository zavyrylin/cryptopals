package cptest

import (
	"cryptopals/aux"
	"io/ioutil"
	"os"
)

// FixtureFile loads fixture file with given name.
func FixtureFile(name string) *os.File {
	path := fixturePath(name)
	fd, err := os.Open(path)
	if err != nil {
		panic(err.Error)
	}
	return fd
}

// FixtureContentsBase64 loads contents of the fixture file encoded using base64
func FixtureContentsBase64(name string) []byte {
	data := loadFixture(name)
	return aux.Base64DecodeBytes(data)
}

func fixturePath(name string) string {
	var path string
	path = "testdata/" + name
	if _, err := os.Stat(path); err == nil {
		return path
	} else if os.IsNotExist(err) {
		path = "../" + path
		if _, err := os.Stat(path); err == nil {
			return path
		}
		panic(err.Error)
	} else {
		panic(err.Error)
	}
}

func loadFixture(name string) []byte {
	path := fixturePath(name)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error)
	}
	return data
}
