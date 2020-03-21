package flag_test

import (
	"flag"
	"strings"
	"testing"

	xflag "neilpa.me/go-x/flag"
)

func TestMultiString(t *testing.T) {
	tests := []struct{
		in string
		out []string
	} {
		{ "" , []string{} },
		{ "-f abc" , []string{"abc"} },
		{ "-f 1 -f 2 -f 3" , []string{"1", "2", "3"} },
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			fs := flag.NewFlagSet(tt.in, flag.ContinueOnError)
			var multi xflag.MultiString
			fs.Var(&multi, "f", "test flag")

			err := fs.Parse(strings.Split(tt.in, " "))
			if err != nil {
				t.Fatalf("parse: %s", err)
			}
			if len(multi) != len(tt.out) {
				t.Fatalf("length: got %d, want %d", len(multi), len(tt.out))
			}
			for i, s := range multi {
				if tt.out[i] != s {
					t.Errorf("mismatch @%d: got %s, want %s", i, s, tt.out[i])
				}
			}
		})
	}
}
