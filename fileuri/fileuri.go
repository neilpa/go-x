package fileuri

import (
	"fmt"
	"path/filepath"
	"strings"
)

// FromPath converts a local filesystem path URI with the file: scheme.
// Relative paths are resolved using filepath.Abs before conversion.
func FromPath(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	// TODO: Need to percent-encode more than just spaces and figure
	// out if there are unix vs. windows differences.
	perc := strings.ReplaceAll(path, " ", "%20")

	vol := filepath.VolumeName(abs)
	if len(vol) > 0 {
		// TODO: Windows requires special URI logic
		// https://docs.microsoft.com/en-us/archive/blogs/ie/file-uris-in-windows
		if vol[0] == '\\' {
			return "", fmt.Errorf("fileuri: TODO: UNC path not implemented: %s", path)
		}
		perc = "/" + filepath.ToSlash(perc)
	}

	return "file://" + perc, nil
}
