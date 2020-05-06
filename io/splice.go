package xio

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Splice safely embedds new data in the middle of the file at path.
//
// Notes:
//	- requires creating a copy of the original file that is len(data) larger
//	- allocates a temp-file in os.TempDir
//
// TODO: potential tweaks
//	- optional overwrite N bytes at offset
//	- optional temp file path to use
//	- context for read/write deadlines or cancellation
//	- fail early if path is a directory
func Splice(path string, data []byte, off int64) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	return SpliceF(f, data, off)
}

// SpliceF safely embedds new data in the middle of the file.
//
// Notes:
//	- The provided file descriptor is closed on return via a defer
func SpliceF(f *os.File, data []byte, off int64) error {
	defer f.Close()
	f.Seek(0, io.SeekStart)
	path := f.Name()

	info, err := f.Stat()
	if err != nil {
		return err
	}
	if info.Size() < off {
		return fmt.Errorf("splice: %s: offset past end of file", path)
	}

	// Prep a temp file of the write size
	tmp, err := ioutil.TempFile(os.TempDir(), "splice-" + filepath.Base(path) + "-")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)

	//debug("%s: truncate tmp=%s", path, tmpPath)
	err = tmp.Truncate(info.Size() + int64(len(data))) // 
	if err != nil {
		return err
	}

	// Copy the header bytes
	//debug("%s: copy header tmp=%s off=0x%x", path, tmpPath, off)
	if _, err = io.Copy(tmp, io.LimitReader(f, off)); err != nil {
		return err
	}
	// Insert the spliced bytes
	//debug("%s: insert splice tmp=%s", path, tmpPath)
	if _, err = tmp.Write(data); err != nil {
		return err
	}
	// Copy the trailer bytes
	//debug("%s: copy trailer tmp=%s", path, tmpPath)
	if _, err = io.Copy(tmp, f); err != nil {
		return err
	}

	// "Atomically" move the temp file back to path
	//debug("%s: rename tmp=%s", path, tmpPath)
	err = f.Close()
	if err != nil {
		return err
	}
	err = tmp.Close()
	if err != nil {
		return err
	}
	return os.Rename(tmp.Name(), path)
}

func debug(format string, head interface{}, tail ...interface{}) {
	format = os.Args[0] + ": " + format + "\n"
	args := make([]interface{}, 1, len(tail) + 1)
	args[0] = head
	args = append(args, tail...)
	fmt.Fprintf(os.Stderr, format, args...)
}
