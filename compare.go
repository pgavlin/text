// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import "github.com/pgavlin/text/internal/bytealg"

// Compare returns an integer comparing two strings lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
//
// Compare is included only for symmetry with package bytes.
// It is usually clearer and always faster to use the built-in
// string comparison operators ==, <, >, and so on.
func Compare[S1, S2 String](a S1, b S2) int {
	return bytealg.Compare(a, b)
}

// Equal compares two strings for equality.
func Equal[S1, S2 String](a S1, b S2) bool {
	return bytealg.AsString(a) == bytealg.AsString(b)
}
