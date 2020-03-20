package fileuri

import (
	"fmt"
	"path/filepath"
	"syscall"
	"unicode/utf16"
	"unsafe"

	"neilpa.me/go-x/win"
)

var (
	shlwapi  = syscall.NewLazyDLL("shlwapi.dll")
	fromPath = shlwapi.NewProc("UrlCreateFromPathW")
)

// FromPath converts a local filesystem path to URI with the file: scheme.
// Relative paths are resolved using filepath.Abs before conversion. The
// Windows implementation uses UrlCreateFromPath in shlwapi.dll to perform
// the conversion.
func FromPath(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	// TODO Test this with "international" characters
	buffer := make([]uint16, 1024)
	size := len(buffer)

	res, _, err := fromPath.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(abs))),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(unsafe.Pointer(&size)),
		0)
	// syscalls always return an error, even if it's Errno = 0 meaning
	// "The operation completed succssefully"
	if errno, ok := err.(syscall.Errno); ok && errno != 0 {
		return "", err
	}
	if hr := win.HRESULT(res); hr.Failed() {
		return "", hr
	}
	return string(utf16.Decode(buffer[:size])), nil
}

// FromWinPath converts a Windows filesystem path to URI with file: scheme.
// Unlike FromPath, this is not platform dependent. For example, C:/foo/bar.txt
// is a valid Unix-path and would be converted differently on POSIX systems vs
// Windows. That is not the case here.
func FromWinPath(path string) (string, error) {
	return FromPath(path)
}
