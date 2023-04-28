package bytealg

import (
	"unsafe"
)

// MaxLen is the maximum length of the string to be searched for (argument b) in Index.
// If MaxLen is not 0, make sure MaxLen >= 4.
//
//go:linkname MaxLen internal/bytealg.MaxLen
var MaxLen int

// PrimeRK is the prime base used in Rabin-Karp algorithm.
const PrimeRK = 16777619

// AsString returns its input as a string.
func AsString[S ~string | ~[]byte](s S) string {
	return *(*string)(unsafe.Pointer(&s))
}

//go:linkname cmpstring runtime.cmpstring
func cmpstring(a, b string) int

func Compare[S1, S2 ~string | ~[]byte](a S1, b S2) int {
	return cmpstring(AsString(a), AsString(b))
}

//go:linkname countString internal/bytealg.CountString
func countString(s string, c byte) int

func CountString[S ~string | ~[]byte](s S, c byte) int {
	return countString(AsString(s), c)
}

//go:linkname hashStrRev internal/bytealg.HashStrRev
func hashStrRev(sep string) (uint32, uint32)

func HashStrRev[S ~string | ~[]byte](sep S) (uint32, uint32) {
	return hashStrRev(AsString(sep))
}

//go:linkname Cutover internal/bytealg.Cutover
func Cutover(n int) int

//go:linkname indexByteString internal/bytealg.IndexByteString
//go:noescape
func indexByteString(s string, c byte) int

func IndexByteString[S ~string | ~[]byte](s S, c byte) int {
	return indexByteString(AsString(s), c)
}

//go:linkname indexRabinKarp internal/bytealg.IndexRabinKarp
func indexRabinKarp(s, substr string) int

// IndexRabinKarp uses the Rabin-Karp search algorithm to return the index of the
// first occurrence of substr in s, or -1 if not present.
func IndexRabinKarp[S1, S2 ~string | ~[]byte](s S1, substr S2) int {
	return indexRabinKarp(AsString(s), AsString(substr))
}

//go:linkname indexString internal/bytealg.IndexString
//go:noescape
func indexString(a, b string) int

// IndexString returns the index of the first instance of b in a, or -1 if b is not present in a.
// Requires 2 <= len(b) <= MaxLen.
func IndexString[S1, S2 ~string | ~[]byte](a S1, b S2) int {
	return indexString(AsString(a), AsString(b))
}
