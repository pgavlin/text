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

// HashStr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func HashStr[S ~string | ~[]byte](sep S) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*PrimeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, PrimeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

// HashStrRev returns the hash of the reverse of sep and the
// appropriate multiplicative factor for use in Rabin-Karp algorithm.
func HashStrRev[S ~string | ~[]byte](sep S) (uint32, uint32) {
	hash := uint32(0)
	for i := len(sep) - 1; i >= 0; i-- {
		hash = hash*PrimeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, PrimeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

//go:linkname Cutover internal/bytealg.Cutover
func Cutover(n int) int

//go:linkname indexByteString internal/bytealg.IndexByteString
//go:noescape
func indexByteString(s string, c byte) int

func IndexByteString[S ~string | ~[]byte](s S, c byte) int {
	return indexByteString(AsString(s), c)
}

// IndexRabinKarp uses the Rabin-Karp search algorithm to return the index of the
// first occurrence of sep in s, or -1 if not present.
func IndexRabinKarp[S1 ~string | ~[]byte, S2 ~string | ~[]byte](s S1, sep S2) int {
	// Rabin-Karp search
	hashss, pow := HashStr(sep)
	n := len(sep)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*PrimeRK + uint32(s[i])
	}
	if h == hashss && string(s[:n]) == string(sep) {
		return 0
	}
	for i := n; i < len(s); {
		h *= PrimeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashss && string(s[i-n:i]) == string(sep) {
			return i - n
		}
	}
	return -1
}

//go:linkname indexString internal/bytealg.IndexString
//go:noescape
func indexString(a, b string) int

// IndexString returns the index of the first instance of b in a, or -1 if b is not present in a.
// Requires 2 <= len(b) <= MaxLen.
func IndexString[S1, S2 ~string | ~[]byte](a S1, b S2) int {
	return indexString(AsString(a), AsString(b))
}
