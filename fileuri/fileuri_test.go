// +build !windows

package fileuri

import (
	"testing"
)

func TestFromPath(t *testing.T) {
	tests := []struct{
		in, out string
	} {
		{ "/", "file:///" },
		{ "/foo/bar.txt", "file:///foo/bar.txt" },
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			uri, err := FromPath(tt.in)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
				return
			}
			if uri != tt.out {
				t.Errorf("got: %s, want: %s", uri, tt.out)
			}
		})
	}
}
