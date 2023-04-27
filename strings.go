package text

import (
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

type StringBuilder[T Text] struct {
	strings.Builder
}

func (s *StringBuilder[T]) Text() T {
	return T(s.String())
}

func (s *StringBuilder[T]) WriteText(t T) (int, error) {
	return s.WriteString(string(t))
}

type StringReader[T Text] struct {
	strings.Reader
}

func (b *StringReader[T]) Reset(t T) {
	b.Reader.Reset(string(t))
}

func StringsPackage[T Text]() Package[T, *StringBuilder[T], *StringReader[T]] {
	return Strings[T]{}
}

type Strings[T Text] struct{}

func UseStrings[T Text]() bool {
	var t T
	return reflect.TypeOf(t).Kind() == reflect.String
}

func (Strings[T]) Clone(s T) T {
	return T(strings.Clone(string(s)))
}

func (Strings[T]) Compare(a, b T) int {
	return strings.Compare(string(a), string(b))
}

func (Strings[T]) Concat(a, b T) T {
	return T(string(a) + string(b))
}

func (Strings[T]) Contains(s, substr T) bool {
	return strings.Contains(string(s), string(substr))
}

func (Strings[T]) ContainsAny(s T, chars string) bool {
	return strings.ContainsAny(string(s), chars)
}

func (Strings[T]) ContainsRune(s T, r rune) bool {
	return strings.ContainsRune(string(s), r)
}

func (Strings[T]) Count(s, substr T) int {
	return strings.Count(string(s), string(substr))
}

func (Strings[T]) Cut(s, sep T) (before, after T, found bool) {
	b, a, f := strings.Cut(string(s), string(sep))
	return T(b), T(a), f
}

func (Strings[T]) CutPrefix(s, prefix T) (after T, found bool) {
	a, f := strings.CutPrefix(string(s), string(prefix))
	return T(a), f
}

func (Strings[T]) CutSuffix(s, suffix T) (before T, found bool) {
	b, f := strings.CutSuffix(string(s), string(suffix))
	return T(b), f
}

func (Strings[T]) EqualFold(s, t T) bool {
	return strings.EqualFold(string(s), string(t))
}

func (Strings[T]) Fields(s T) []T {
	return fields[T, Strings[T]](s)
}

func (Strings[T]) FieldsFunc(s T, f func(rune) bool) []T {
	return fieldsFunc[T, Strings[T]](s, f)
}

func (Strings[T]) HasPrefix(s, prefix T) bool {
	return strings.HasPrefix(string(s), string(prefix))
}

func (Strings[T]) HasSuffix(s, suffix T) bool {
	return strings.HasSuffix(string(s), string(suffix))
}

func (Strings[T]) Index(s, substr T) int {
	return strings.Index(string(s), string(substr))
}

func (Strings[T]) IndexAny(s T, chars string) int {
	return strings.IndexAny(string(s), chars)
}

func (Strings[T]) IndexByte(s T, c byte) int {
	return strings.IndexByte(string(s), c)
}

func (Strings[T]) IndexFunc(s T, f func(rune) bool) int {
	return strings.IndexFunc(string(s), f)
}

func (Strings[T]) IndexRune(s T, r rune) int {
	return strings.IndexRune(string(s), r)
}

func (Strings[T]) Join(elems []T, sep T) T {
	return join[T, *StringBuilder[T], *StringReader[T], Strings[T]](elems, sep)
}

func (Strings[T]) LastIndex(s, substr T) int {
	return strings.LastIndex(string(s), string(substr))
}

func (Strings[T]) LastIndexAny(s T, chars string) int {
	return strings.LastIndexAny(string(s), chars)
}

func (Strings[T]) LastIndexByte(s T, c byte) int {
	return strings.LastIndexByte(string(s), c)
}

func (Strings[T]) LastIndexFunc(s T, f func(rune) bool) int {
	return strings.LastIndexFunc(string(s), f)
}

func (Strings[T]) Map(mapping func(rune) rune, s T) T {
	return T(strings.Map(mapping, string(s)))
}

func (Strings[T]) NewBuilder() *StringBuilder[T] {
	return &StringBuilder[T]{}
}

func (Strings[T]) NewReader(t T) *StringReader[T] {
	r := &StringReader[T]{}
	r.Reset(t)
	return r
}

func (Strings[T]) Repeat(s T, count int) T {
	return T(strings.Repeat(string(s), count))
}

func (Strings[T]) Replace(s, old, new T, n int) T {
	return T(strings.Replace(string(s), string(old), string(new), n))
}

func (Strings[T]) ReplaceAll(s, old, new T) T {
	return T(strings.ReplaceAll(string(s), string(old), string(new)))
}

func (Strings[T]) Split(s, sep T) []T {
	return genSplit[T, Strings[T]](s, sep, 0, -1)
}

func (Strings[T]) SplitAfter(s, sep T) []T {
	return genSplit[T, Strings[T]](s, sep, len(sep), -1)
}

func (Strings[T]) SplitAfterN(s, sep T, n int) []T {
	return genSplit[T, Strings[T]](s, sep, len(sep), n)
}

func (Strings[T]) SplitN(s, sep T, n int) []T {
	return genSplit[T, Strings[T]](s, sep, 0, n)
}

func (Strings[T]) Title(s T) T {
	return T(strings.Title(string(s)))
}

func (Strings[T]) ToLower(s T) T {
	return T(strings.ToLower(string(s)))
}

func (Strings[T]) ToLowerSpecial(c unicode.SpecialCase, s T) T {
	return T(strings.ToLowerSpecial(c, string(s)))
}

func (Strings[T]) ToRunes(s T) []rune {
	return []rune(string(s))
}

func (Strings[T]) ToText(r []rune) T {
	return T(string(r))
}

func (Strings[T]) ToTitle(s T) T {
	return T(strings.ToTitle(string(s)))
}

func (Strings[T]) ToTitleSpecial(c unicode.SpecialCase, s T) T {
	return T(strings.ToTitleSpecial(c, string(s)))
}

func (Strings[T]) ToUpper(s T) T {
	return T(strings.ToUpper(string(s)))
}

func (Strings[T]) ToUpperSpecial(c unicode.SpecialCase, s T) T {
	return T(strings.ToUpperSpecial(c, string(s)))
}

func (Strings[T]) ToValidUTF8(s, replacement T) T {
	return T(strings.ToValidUTF8(string(s), string(replacement)))
}

func (Strings[T]) Trim(s T, cutset string) T {
	return T(strings.Trim(string(s), cutset))
}

func (Strings[T]) TrimFunc(s T, f func(rune) bool) T {
	return T(strings.TrimFunc(string(s), f))
}

func (Strings[T]) TrimLeft(s T, cutset string) T {
	return T(strings.TrimLeft(string(s), cutset))
}

func (Strings[T]) TrimLeftFunc(s T, f func(rune) bool) T {
	return T(strings.TrimLeftFunc(string(s), f))
}

func (Strings[T]) TrimPrefix(s, prefix T) T {
	return T(strings.TrimPrefix(string(s), string(prefix)))
}

func (Strings[T]) TrimRight(s T, cutset string) T {
	return T(strings.TrimRight(string(s), cutset))
}

func (Strings[T]) TrimRightFunc(s T, f func(rune) bool) T {
	return T(strings.TrimRightFunc(string(s), f))
}

func (Strings[T]) TrimSpace(s T) T {
	return T(strings.TrimSpace(string(s)))
}

func (Strings[T]) TrimSuffix(s, suffix T) T {
	return T(strings.TrimSuffix(string(s), string(suffix)))
}

func (Strings[T]) UTF8DecodeRune(t T) (rune, int) {
	return utf8.DecodeRuneInString(string(t))
}

func (Strings[T]) UTF8RuneCount(t T) int {
	return utf8.RuneCountInString(string(t))
}
