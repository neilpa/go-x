package xio

import (
	"io"
	"io/ioutil"
	"os"
)

// TempFileCopy copies path to a new temp file
func TempFileCopy(path, prefix string) (string, error) {
	temp, err := ioutil.TempFile(os.TempDir(), prefix)
	if err != nil {
		return "", err
	}
	defer temp.Close()

	src, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer src.Close()

	_, err = io.Copy(temp, src)
	if err != nil {
		return "", err
	}

	return temp.Name(), nil
}
