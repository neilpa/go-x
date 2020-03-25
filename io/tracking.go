package xio

import (
	"errors"
	"io"
	"io/ioutil"
)

var (
	// ErrUnseekableReader means a Seek was attempted from the start or end
	// of an io.Reader that only supports streaming.
	ErrUnseekableReader = errors.New("Unseekable reader")
)

// TrackingReader wraps an io.Reader to that tracks the offset as bytes
// are read. Additionally, it adds a best-effort io.Seeker implementation.
// For a pure io.Reader that is limited to usage of io.SeeekCurrent and
// otherwise fails for seeks relative to the start or end of the stream.
type TrackingReader struct {
	reader io.Reader
	pos  int64
}

// NewTrackingReader wraps an io.Reader in a TrackingReader
func NewTrackingReader(r io.Reader) *TrackingReader {
	return &TrackingReader{r, 0}
}

// Offset returns the current position of the reader.
func (tr TrackingReader) Offset() int64 {
	return tr.pos
}

// Read is a pass-thru to the underlying Read.
func (tr *TrackingReader) Read(p []byte) (n int, err error) {
	n, err = tr.reader.Read(p)
	tr.pos += int64(n)
	return
}

// Seek implements io.Seeker. If the wrapped io.Reader also implements
// io.Seeker this is a pass-thru. Otherwise, only io.SeekCurrent is
// supported and ErrUnseekableReader is returned for seeks from start/end.
func (tr *TrackingReader) Seek(offset int64, whence int) (int64, error) {
	var err error
	switch s := tr.reader.(type) {
		case io.Seeker:
			tr.pos, err = s.Seek(offset, whence)
		default:
			// TODO Could support io.SeekStart if tr.pos <= offset
			if whence != io.SeekCurrent {
				err = ErrUnseekableReader
			} else {
				var n int64
				n, err = io.CopyN(ioutil.Discard, tr.reader, offset)
				tr.pos += n
			}
	}
	return tr.pos, err
}
