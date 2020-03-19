package fileuri

import (
	"testing"
)

func TestFromPath(t *testing.T) {
	tests := []struct{
		in, out string
	} {
		{ `C:`, "file:///C:" },
		{ `C:\foo\bar.txtt`, "file:///C:/foo/bar.txt" },
		{ `\\host\path\file.ext`, "file://host/path/file.ext" },
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
