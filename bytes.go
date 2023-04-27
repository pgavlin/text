package text

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

type BytesBuilder[T Text] struct {
	bytes.Buffer
}

func (s *BytesBuilder[T]) Text() T {
	return T(s.Bytes())
}

func (s *BytesBuilder[T]) WriteText(t T) (int, error) {
	return s.Write([]byte(t))
}

type BytesReader[T Text] struct {
	bytes.Reader
}

func (b *BytesReader[T]) Reset(t T) {
	b.Reader.Reset([]byte(t))
}

func BytesPackage[T Text]() Package[T, *BytesBuilder[T], *BytesReader[T]] {
	return Bytes[T]{}
}

type Bytes[T Text] struct{}

func (Bytes[T]) Clone(s T) T {
	return T(bytes.Clone([]byte(s)))
}

func (Bytes[T]) Compare(a, b T) int {
	return bytes.Compare([]byte(a), []byte(b))
}

func (Bytes[T]) Concat(a, b T) T {
	c := make([]byte, len(a)+len(b))
	copy(c, []byte(a))
	copy(c[len(a):], []byte(b))
	return T(c)
}

func (Bytes[T]) Contains(s, substr T) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}

func (Bytes[T]) ContainsAny(s T, chars string) bool {
	return bytes.ContainsAny([]byte(s), chars)
}

func (Bytes[T]) ContainsRune(s T, r rune) bool {
	return bytes.ContainsRune([]byte(s), r)
}

func (Bytes[T]) Count(s, substr T) int {
	return bytes.Count([]byte(s), []byte(substr))
}

func (Bytes[T]) Cut(s, sep T) (before, after T, found bool) {
	b, a, f := bytes.Cut([]byte(s), []byte(sep))
	return T(b), T(a), f
}

func (Bytes[T]) CutPrefix(s, prefix T) (after T, found bool) {
	a, f := bytes.CutPrefix([]byte(s), []byte(prefix))
	return T(a), f
}

func (Bytes[T]) CutSuffix(s, suffix T) (before T, found bool) {
	b, f := bytes.CutSuffix([]byte(s), []byte(suffix))
	return T(b), f
}

func (Bytes[T]) EqualFold(s, t T) bool {
	return bytes.EqualFold([]byte(s), []byte(t))
}

func (Bytes[T]) Fields(s T) []T {
	return fields[T, Bytes[T]](s)
}

func (Bytes[T]) FieldsFunc(s T, f func(rune) bool) []T {
	return fieldsFunc[T, Bytes[T]](s, f)
}

func (Bytes[T]) HasPrefix(s, prefix T) bool {
	return bytes.HasPrefix([]byte(s), []byte(prefix))
}

func (Bytes[T]) HasSuffix(s, suffix T) bool {
	return bytes.HasSuffix([]byte(s), []byte(suffix))
}

func (Bytes[T]) Index(s, substr T) int {
	return bytes.Index([]byte(s), []byte(substr))
}

func (Bytes[T]) IndexAny(s T, chars string) int {
	return bytes.IndexAny([]byte(s), chars)
}

func (Bytes[T]) IndexByte(s T, c byte) int {
	return bytes.IndexByte([]byte(s), c)
}

func (Bytes[T]) IndexFunc(s T, f func(rune) bool) int {
	return bytes.IndexFunc([]byte(s), f)
}

func (Bytes[T]) IndexRune(s T, r rune) int {
	return bytes.IndexRune([]byte(s), r)
}

func (Bytes[T]) Join(elems []T, sep T) T {
	return join[T, *BytesBuilder[T], *BytesReader[T], Bytes[T]](elems, sep)
}

func (Bytes[T]) LastIndex(s, substr T) int {
	return bytes.LastIndex([]byte(s), []byte(substr))
}

func (Bytes[T]) LastIndexAny(s T, chars string) int {
	return bytes.LastIndexAny([]byte(s), chars)
}

func (Bytes[T]) LastIndexByte(s T, c byte) int {
	return bytes.LastIndexByte([]byte(s), c)
}

func (Bytes[T]) LastIndexFunc(s T, f func(rune) bool) int {
	return bytes.LastIndexFunc([]byte(s), f)
}

func (Bytes[T]) Map(mapping func(rune) rune, s T) T {
	return T(bytes.Map(mapping, []byte(s)))
}

func (Bytes[T]) NewBuilder() *BytesBuilder[T] {
	return &BytesBuilder[T]{}
}

func (Bytes[T]) NewReader(t T) *BytesReader[T] {
	r := &BytesReader[T]{}
	r.Reset(t)
	return r
}

func (Bytes[T]) Repeat(s T, count int) T {
	return T(bytes.Repeat([]byte(s), count))
}

func (Bytes[T]) Replace(s, old, new T, n int) T {
	return T(bytes.Replace([]byte(s), []byte(old), []byte(new), n))
}

func (Bytes[T]) ReplaceAll(s, old, new T) T {
	return T(bytes.ReplaceAll([]byte(s), []byte(old), []byte(new)))
}

func (Bytes[T]) Split(s, sep T) []T {
	return genSplit[T, Bytes[T]](s, sep, 0, -1)
}

func (Bytes[T]) SplitAfter(s, sep T) []T {
	return genSplit[T, Bytes[T]](s, sep, len(sep), -1)
}

func (Bytes[T]) SplitAfterN(s, sep T, n int) []T {
	return genSplit[T, Bytes[T]](s, sep, len(sep), n)
}

func (Bytes[T]) SplitN(s, sep T, n int) []T {
	return genSplit[T, Bytes[T]](s, sep, 0, n)
}

func (Bytes[T]) Title(s T) T {
	return T(bytes.Title([]byte(s)))
}

func (Bytes[T]) ToLower(s T) T {
	return T(bytes.ToLower([]byte(s)))
}

func (Bytes[T]) ToLowerSpecial(c unicode.SpecialCase, s T) T {
	return T(bytes.ToLowerSpecial(c, []byte(s)))
}

func (Bytes[T]) ToRunes(s T) []rune {
	bytes := []byte(s)
	n := utf8.RuneCount(bytes)
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		r, sz := utf8.DecodeRune(bytes)
		bytes = bytes[sz:]
		runes[i] = r
	}
	return runes
}

func (Bytes[T]) ToText(r []rune) T {
	len := 0
	for _, r := range r {
		len += utf8.RuneLen(r)
	}
	bytes := make([]byte, len)

	cursor := bytes
	for _, r := range r {
		c := utf8.EncodeRune(cursor, r)
		cursor = cursor[c:]
	}
	return T(bytes)
}

func (Bytes[T]) ToTitle(s T) T {
	return T(bytes.ToTitle([]byte(s)))
}

func (Bytes[T]) ToTitleSpecial(c unicode.SpecialCase, s T) T {
	return T(bytes.ToTitleSpecial(c, []byte(s)))
}

func (Bytes[T]) ToUpper(s T) T {
	return T(bytes.ToUpper([]byte(s)))
}

func (Bytes[T]) ToUpperSpecial(c unicode.SpecialCase, s T) T {
	return T(bytes.ToUpperSpecial(c, []byte(s)))
}

func (Bytes[T]) ToValidUTF8(s, replacement T) T {
	return T(bytes.ToValidUTF8([]byte(s), []byte(replacement)))
}

func (Bytes[T]) Trim(s T, cutset string) T {
	return T(bytes.Trim([]byte(s), cutset))
}

func (Bytes[T]) TrimFunc(s T, f func(rune) bool) T {
	return T(bytes.TrimFunc([]byte(s), f))
}

func (Bytes[T]) TrimLeft(s T, cutset string) T {
	return T(bytes.TrimLeft([]byte(s), cutset))
}

func (Bytes[T]) TrimLeftFunc(s T, f func(rune) bool) T {
	return T(bytes.TrimLeftFunc([]byte(s), f))
}

func (Bytes[T]) TrimPrefix(s, prefix T) T {
	return T(bytes.TrimPrefix([]byte(s), []byte(prefix)))
}

func (Bytes[T]) TrimRight(s T, cutset string) T {
	return T(bytes.TrimRight([]byte(s), cutset))
}

func (Bytes[T]) TrimRightFunc(s T, f func(rune) bool) T {
	return T(bytes.TrimRightFunc([]byte(s), f))
}

func (Bytes[T]) TrimSpace(s T) T {
	return T(bytes.TrimSpace([]byte(s)))
}

func (Bytes[T]) TrimSuffix(s, suffix T) T {
	return T(bytes.TrimSuffix([]byte(s), []byte(suffix)))
}

func (Bytes[T]) UTF8DecodeRune(t T) (rune, int) {
	return utf8.DecodeRune([]byte(t))
}

func (Bytes[T]) UTF8RuneCount(t T) int {
	return utf8.RuneCount([]byte(t))
}
