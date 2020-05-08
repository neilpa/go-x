package xio

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var spliceTests = []struct {
	name   string
	data   string
	offset int64
	golden string
}{
	{"abc.txt", "123\n", 0, "abc.head.txt"},
	{"abc.txt", "456\n", 4, "abc.mid.txt"},
	{"abc.txt", "789\n", 8, "abc.tail.txt"},
}

func TestSplice(t *testing.T) {
	for _, tt := range spliceTests {
		t.Run(tt.golden, func(t *testing.T) {
			path, err := TempFileCopy(filepath.Join("testdata", tt.name), "jfif-test-update-"+tt.golden)
			if err != nil {
				t.Fatal(err)
			}
			//defer os.Remove(path)

			err = Splice(path, []byte(tt.data), tt.offset)
			if err != nil {
				t.Fatal(err)
			}

			compareFiles(t, path, filepath.Join("testdata", tt.golden))
		})
	}
}

func TestSpliceF(t *testing.T) {
	for _, tt := range spliceTests {
		t.Run(tt.golden, func(t *testing.T) {
			path, err := TempFileCopy(filepath.Join("testdata", tt.name), "jfif-test-update-"+tt.golden)
			if err != nil {
				t.Fatal(err)
			}
			//defer os.Remove(path)

			f, err := os.Open(path)
			if err != nil {
				t.Fatal(err)
			}

			err = SpliceFile(f, []byte(tt.data), tt.offset)
			if err != nil {
				t.Fatal(err)
			}

			compareFiles(t, path, filepath.Join("testdata", tt.golden))
		})
	}
}

func compareFiles(t *testing.T, path, golden string) {
	want, err := ioutil.ReadFile(golden)
	if err != nil {
		t.Fatal(err)
	}
	got, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, want) {
		t.Errorf("bytes don't match\ngot:  % x\nwant: % x", got, want) // TODO Better diff
	}
}
