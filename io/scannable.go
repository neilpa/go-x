// Package xio provides extensions and hacks on top of io from the stdlib
package xio

import "io"

// ScannableReader wraps poorly behaved io.Reader implementations to make
// them safe to use with bufio.Scanner. In particular, when the underlying
// reader returns an error, it ensures that negative values for bytes read
// don't bubble out. Instead, this will return 0 along with the original
// error. One source of such bugs returning values directly from
// golang.org/x/sys/unix.Read.
type ScannableReader struct {
	r io.Reader
}

// NewScanableReader wraps an io.Reader in a ScannableReader
func NewScannableReader(r io.Reader) *ScannableReader {
	return &ScannableReader{r}
}

// Read implements io.Reader and "fixes" poorly behaved implementations
// during error conditions
func (sr ScannableReader) Read(p []byte) (n int, err error) {
	n, err = sr.r.Read(p)
	if err != nil && n < 0 {
		n = 0
	}
	return
}
