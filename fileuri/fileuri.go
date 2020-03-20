// +build !windows

package fileuri // import "neilpa.me/go-x/fileuri"

import (
	"fmt"
	"net/url"
	"path/filepath"
)

// FromPath converts a local filesystem path to URI with the file: scheme.
// Relative paths are resolved using filepath.Abs before conversion.
func FromPath(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	// Ensure percent-encoding where needed
	perc, err := url.Parse(abs)
	if err != nil {
		return "", err
	}
	return "file://" + perc.EscapedPath(), nil
}

// FromWinPath converts a Windows filesystem path to URI with file: scheme.
// Unlike FromPath, this is not platform dependent. For example, C:/foo/bar.txt
// is a valid Unix-path and would be converted differently on POSIX systems vs
// Windows. That is not the case here.
func FromWinPath(path string) (string, error) {
	return "", fmt.Errorf("TODO: convert %s to URI", path)
}
