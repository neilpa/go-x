// +build windows

// Package win bridges some windows primitives for easier use in go.
package win // import "neilpa.me/go-x/win"

import "fmt"

// HRESULT provides some nominal result values for Windows error codes.
// Unlike the Windows header files, the HRESULT type is declared as uint32
// rather than an int32. This side steps compile-time overflow issues in go
// when using the standard hex constants to define the named values.
//
// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-erref/0642cb2f-2075-4469-918c-4441e69c548a
type HRESULT uint32

// https://docs.microsoft.com/en-us/windows/win32/seccrypto/common-hresult-values
const (
	S_OK           HRESULT = 0x00000000
	E_ABORT                = 0x80004004
	E_ACCESSDENIED         = 0x80070005
	E_FAIL                 = 0x80004005
	E_HANDLE               = 0x80070006
	E_INVALIDARG           = 0x80070057
	E_NOINTERFACE          = 0x80004002
	E_NOTIMPL              = 0x80004001
	E_OUTOFMEMORY          = 0x8007000E
	E_POINTER              = 0x80004003
	E_UNEXPECTED           = 0x8000FFFF
)

// Failed reports if the HRESULT corresponds to an actual error.
// https://docs.microsoft.com/en-us/windows/win32/api/winerror/nf-winerror-failed
func (hr HRESULT) Failed() bool {
	return hr&0x80000000 != 0
}

// Error implements the error interface.
func (hr HRESULT) Error() string {
	msg := "<unknown hresult>"
	switch hr {
	case S_OK:
		msg = "Operation successful"
	case E_ABORT:
		msg = "Operation aborted"
	case E_ACCESSDENIED:
		msg = "General access denied error"
	case E_FAIL:
		msg = "Unspecified failure"
	case E_HANDLE:
		msg = "Handle that is not valid"
	case E_INVALIDARG:
		msg = "One or more arguments are not valid"
	case E_NOINTERFACE:
		msg = "No such interface supported"
	case E_NOTIMPL:
		msg = "Not implemented"
	case E_OUTOFMEMORY:
		msg = "Failed to allocate necessary memory"
	case E_POINTER:
		msg = "Pointer that is not valid"
	case E_UNEXPECTED:
		msg = "Unexpected failure"
	}
	return fmt.Sprintf("%s (hr = 0x%X)", msg, hr)
}
