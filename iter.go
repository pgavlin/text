// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import (
	"iter"
	"unicode"

	"github.com/pgavlin/text/utf8"
)

// Lines returns an iterator over the newline-terminated lines in the string s.
// The lines yielded by the iterator include their terminating newlines.
// If s is empty, the iterator yields no lines at all.
// If s does not end in a newline, the final yielded line will not end in a newline.
// It returns a single-use iterator.
func Lines[S String](s S) iter.Seq[S] {
	return func(yield func(S) bool) {
		for len(s) > 0 {
			var line S
			if i := IndexByte(s, '\n'); i >= 0 {
				line, s = s[:i+1], s[i+1:]
			} else {
				line, s = s, s[len(s):]
			}
			if !yield(line) {
				return
			}
		}
	}
}

// splitSeq is SplitSeq or SplitAfterSeq, configured by how many
// bytes of sep to include in the results (none or all).
func splitSeq[S1, S2 String](s S1, sep S2, sepSave int) iter.Seq[S1] {
	return func(yield func(S1) bool) {
		if len(sep) == 0 {
			for len(s) > 0 {
				_, size := utf8.DecodeRune(s)
				if !yield(s[:size]) {
					return
				}
				s = s[size:]
			}
			return
		}
		for {
			i := Index(s, sep)
			if i < 0 {
				break
			}
			frag := s[:i+sepSave]
			if !yield(frag) {
				return
			}
			s = s[i+len(sep):]
		}
		yield(s)
	}
}

// SplitSeq returns an iterator over all substrings of s separated by sep.
// The iterator yields the same strings that would be returned by [Split](s, sep),
// but without constructing the slice.
// It returns a single-use iterator.
func SplitSeq[S1, S2 String](s S1, sep S2) iter.Seq[S1] {
	return splitSeq(s, sep, 0)
}

// SplitAfterSeq returns an iterator over substrings of s split after each instance of sep.
// The iterator yields the same strings that would be returned by [SplitAfter](s, sep),
// but without constructing the slice.
// It returns a single-use iterator.
func SplitAfterSeq[S1, S2 String](s S1, sep S2) iter.Seq[S1] {
	return splitSeq(s, sep, len(sep))
}

// FieldsSeq returns an iterator over substrings of s split around runs of
// whitespace characters, as defined by [unicode.IsSpace].
// The iterator yields the same strings that would be returned by [Fields](s),
// but without constructing the slice.
func FieldsSeq[S String](s S) iter.Seq[S] {
	return func(yield func(S) bool) {
		start := -1
		for i := 0; i < len(s); {
			size := 1
			r := rune(s[i])
			isSpace := asciiSpace[s[i]] != 0
			if r >= utf8.RuneSelf {
				r, size = utf8.DecodeRune(s[i:])
				isSpace = unicode.IsSpace(r)
			}
			if isSpace {
				if start >= 0 {
					if !yield(s[start:i]) {
						return
					}
					start = -1
				}
			} else if start < 0 {
				start = i
			}
			i += size
		}
		if start >= 0 {
			yield(s[start:])
		}
	}
}

// FieldsFuncSeq returns an iterator over substrings of s split around runs of
// Unicode code points satisfying f(c).
// The iterator yields the same strings that would be returned by [FieldsFunc](s),
// but without constructing the slice.
func FieldsFuncSeq[S String](s S, f func(rune) bool) iter.Seq[S] {
	return func(yield func(S) bool) {
		start := -1
		for i := 0; i < len(s); {
			size := 1
			r := rune(s[i])
			if r >= utf8.RuneSelf {
				r, size = utf8.DecodeRune(s[i:])
			}
			if f(r) {
				if start >= 0 {
					if !yield(s[start:i]) {
						return
					}
					start = -1
				}
			} else if start < 0 {
				start = i
			}
			i += size
		}
		if start >= 0 {
			yield(s[start:])
		}
	}
}
