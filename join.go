// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

func join[T Text, B Builder[T], R Reader[T], P Package[T, B, R]](elems []T, sep T) T {
	var pkg P

	switch len(elems) {
	case 0:
		var t T
		return t
	case 1:
		return elems[0]
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}

	b := pkg.NewBuilder()
	b.Grow(n)
	b.WriteText(elems[0])
	for _, s := range elems[1:] {
		b.WriteText(sep)
		b.WriteText(s)
	}
	return b.Text()
}
