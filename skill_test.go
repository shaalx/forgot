package forgot

import (
	"testing"
	// "unsafe"
)

func TestB2S(t *testing.T) {
	buf := []byte("This is a string buf.")
	p := B2S(buf)
	t.Log(uintptr(p))
}
