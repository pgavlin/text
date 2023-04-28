// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import (
	"unsafe"
)

// Clone returns a fresh copy of s.
// It guarantees to make a copy of s into a new allocation,
// which can be important when retaining only a small substring
// of a much larger string. Using Clone can help such programs
// use less memory. Of course, since using Clone makes a copy,
// overuse of Clone can make programs use more memory.
// Clone should typically be used only rarely, and only when
// profiling indicates that it is needed.
// For strings of length zero the string "" will be returned
// and no allocation is made.
func Clone[S String](s S) S {
	if IsEmpty(s) {
		return Empty[S]()
	}

	b := make([]byte, len(s))
	copy(b, s)

	if isString[S]() {
		return S(unsafe.String(&b[0], len(b)))
	}
	return S(b)
}
