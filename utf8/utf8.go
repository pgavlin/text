// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package utf8 implements functions and constants to support text encoded in
// UTF-8. It includes functions to translate between runes and UTF-8 byte sequences.
// See https://en.wikipedia.org/wiki/UTF-8
package utf8

import (
	"unicode/utf8"

	"github.com/pgavlin/text/internal/bytealg"
)

// Numbers fundamental to the encoding.
const (
	RuneError = utf8.RuneError // the "error" Rune or "Unicode replacement character"
	RuneSelf  = utf8.RuneSelf  // characters below RuneSelf are represented as themselves in a single byte.
	MaxRune   = utf8.MaxRune   // Maximum valid Unicode code point.
	UTFMax    = utf8.UTFMax    // maximum number of bytes of a UTF-8 encoded Unicode character.
)

// FullRune reports whether the bytes in p begin with a full UTF-8 encoding of a rune.
// An invalid encoding is considered a full Rune since it will convert as a width-1 error rune.
func FullRune[S ~string | ~[]byte](p S) bool {
	return utf8.FullRuneInString(bytealg.AsString(p))
}

// DecodeRune unpacks the first UTF-8 encoding in p and returns the rune and
// its width in bytes. If p is empty it returns (RuneError, 0). Otherwise, if
// the encoding is invalid, it returns (RuneError, 1). Both are impossible
// results for correct, non-empty UTF-8.
//
// An encoding is invalid if it is incorrect UTF-8, encodes a rune that is
// out of range, or is not the shortest possible UTF-8 encoding for the
// value. No other validation is performed.
func DecodeRune[S ~string | ~[]byte](s S) (r rune, size int) {
	return utf8.DecodeRuneInString(bytealg.AsString(s))
}

// DecodeLastRune unpacks the last UTF-8 encoding in p and returns the rune and
// its width in bytes. If p is empty it returns (RuneError, 0). Otherwise, if
// the encoding is invalid, it returns (RuneError, 1). Both are impossible
// results for correct, non-empty UTF-8.
//
// An encoding is invalid if it is incorrect UTF-8, encodes a rune that is
// out of range, or is not the shortest possible UTF-8 encoding for the
// value. No other validation is performed.
func DecodeLastRune[S ~string | ~[]byte](s S) (r rune, size int) {
	return utf8.DecodeLastRuneInString(bytealg.AsString(s))
}

// RuneLen returns the number of bytes required to encode the rune.
// It returns -1 if the rune is not a valid value to encode in UTF-8.
func RuneLen(r rune) int {
	return utf8.RuneLen(r)
}

// EncodeRune writes into p (which must be large enough) the UTF-8 encoding of the rune.
// If the rune is out of range, it writes the encoding of RuneError.
// It returns the number of bytes written.
func EncodeRune(p []byte, r rune) int {
	return utf8.EncodeRune(p, r)
}

// AppendRune appends the UTF-8 encoding of r to the end of p and
// returns the extended buffer. If the rune is out of range,
// it appends the encoding of RuneError.
func AppendRune(p []byte, r rune) []byte {
	return utf8.AppendRune(p, r)
}

// RuneCount returns the number of runes in p. Erroneous and short
// encodings are treated as single runes of width 1 byte.
func RuneCount[S ~string | ~[]byte](s S) int {
	return utf8.RuneCountInString(bytealg.AsString(s))
}

// RuneStart reports whether the byte could be the first byte of an encoded,
// possibly invalid rune. Second and subsequent bytes always have the top two
// bits set to 10.
func RuneStart(b byte) bool {
	return utf8.RuneStart(b)
}

// Valid reports whether p consists entirely of valid UTF-8-encoded runes.
func Valid[S ~string | ~[]byte](s S) bool {
	return utf8.ValidString(bytealg.AsString(s))
}

// ValidRune reports whether r can be legally encoded as UTF-8.
// Code points that are out of range or a surrogate half are illegal.
func ValidRune(r rune) bool {
	return utf8.ValidRune(r)
}
