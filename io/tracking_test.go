package xio

import (
	"testing"
)

func TestTrackingReader(t *testing.T) { // TODO
	// Read basics
	// Seek wrapping io.ReadSeeker works
	// Seek wrapping io.Reader
	//	- works from current
	//	- may work from start
	//	- always fails from end
}
