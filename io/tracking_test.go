package xio

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

type reader struct {
	r io.Reader
}

func (r reader) Read(p []byte) (int, error) {
	return r.r.Read(p)
}

func ExampleTrackingReader() {
	text := "some text"
	r := NewTrackingReader(reader{strings.NewReader(text)})

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
	if r.Offset() != int64(len(text)) {
		log.Fatalf("offset got %d want %d", r.Offset(), len(text))
	}

	// Output:
	// some text
}

func TestTrackingReader_Seek(t *testing.T) {
	text := "some text"

	tests := []struct{
		skip int64

		seek int64
		whence int

		offset int64
		fail bool
	} {
		{ 0, 0, io.SeekCurrent, 0, false },
		{ 0, 0, io.SeekStart, 0, false },

		{ 1, 0, io.SeekCurrent, 1, false },
		{ 1, 1, io.SeekCurrent, 2, false },

		{ 1, 2, io.SeekStart, 2, false },
		{ 1, 1, io.SeekStart, 1, false },
		{ 1, 0, io.SeekStart, 1, true },

		{ 2, -1, io.SeekCurrent, 2, true},

		{ 0, 0, io.SeekEnd, 0, true },
		{ 5, 0, io.SeekEnd, 5, true },
		{ int64(len(text)), -2, io.SeekEnd, int64(len(text)), true },
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := NewTrackingReader(reader{strings.NewReader(text)})
			r.pos = tt.skip

			off, err := r.Seek(tt.seek, tt.whence)
			if err != nil && !tt.fail {
				t.Errorf("seek: unexpected failure %s", err)
			}
			if off != tt.offset {
				t.Errorf("offset: got %d want %d", off, tt.offset)
			}
			if r.Offset() != tt.offset {
				t.Errorf("reader: got %d want %d", r.Offset(), tt.offset)
			}
		})
	}
}
