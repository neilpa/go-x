package flag // import "neilpa.me/go-x/flag" 

import "fmt"

// MultiString implements the standard flag.Value interface for accumulating string values
// that may be set multiple times for the same argument name. For example,
//
//   $ command -i one -i two
//
// Result in the slice { "one", "two" } after flag parsing
type MultiString []string

// String serializes to a string for the flag.Value interface.
func (ms *MultiString) String() string {
	if ms == nil {
		return "<nil-multi-string>"
	}
	return fmt.Sprintf("%s", []string(*ms))
}

// Set accumulates values for the flag.Value interface.
func (ms *MultiString) Set(value string) error {
	*ms = append(*ms, value)
	return nil
}
