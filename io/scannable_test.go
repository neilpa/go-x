package xio

import (
	"bufio"
	"errors"
	"testing"
)

func TestScannableReader(t *testing.T) {
	lines := []string{"one", "two", "three"}
	r := &badReader{lines, errors.New("bad")}
	scanner := bufio.NewScanner(NewScannableReader(r))
	for i := 0; scanner.Scan(); i++ {
		got := scanner.Text()
		if lines[i] != got {
			t.Errorf("scan @%d: got %q want %q", i, got, lines[i])
		}
	}
	if err := scanner.Err(); err != r.err {
		t.Errorf("scan err: got %q want %q", err, r.err)
	}
}

type badReader struct {
	lines []string
	err   error
}

func (r *badReader) Read(p []byte) (n int, err error) {
	if len(r.lines) > 0 {
		n = copy(p, []byte(r.lines[0]+"\n"))
		r.lines = r.lines[1:]
	} else {
		n = -1
		err = r.err
	}
	return
}
