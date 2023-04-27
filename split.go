// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

// Generic split: splits after each instance of sep,
// including sepSave bytes of sep in the subarrays.
func genSplit[T Text, A Algorithms[T]](s, sep T, sepSave, n int) []T {
	var alg A

	if n == 0 {
		return nil
	}
	if len(sep) == 0 {
		return explode[T, A](s, n)
	}
	if n < 0 {
		n = alg.Count(s, sep) + 1
	}

	if n > len(s)+1 {
		n = len(s) + 1
	}
	a := make([]T, n)
	n--
	i := 0
	for i < n {
		m := alg.Index(s, sep)
		if m < 0 {
			break
		}
		a[i] = s[:m+sepSave]
		s = s[m+len(sep):]
		i++
	}
	a[i] = s
	return a[:i+1]
}

// explode splits s into a slice of UTF-8 strings,
// one string per Unicode character up to a maximum of n (n < 0 means no limit).
// Invalid UTF-8 bytes are sliced individually.
func explode[T Text, A Algorithms[T]](s T, n int) []T {
	var alg A

	l := alg.UTF8RuneCount(s)
	if n < 0 || n > l {
		n = l
	}
	a := make([]T, n)
	for i := 0; i < n-1; i++ {
		_, size := alg.UTF8DecodeRune(s)
		a[i] = s[:size]
		s = s[size:]
	}
	if n > 0 {
		a[n-1] = s
	}
	return a
}
