package bytealg

import (
	"unsafe"
)

// AsString returns its input as a string.
func AsString[S ~string | ~[]byte](s S) string {
	return *(*string)(unsafe.Pointer(&s))
}
