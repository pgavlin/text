package text

import (
	"io"
	"unicode"
)

// Text represents a piece of text (either a string or a byte slice).
type Text interface {
	~string | ~[]byte
}

type Builder[T Text] interface {
	Cap() int
	Grow(n int)
	Len() int
	Reset()
	String() string
	Text() T
	Write(p []byte) (int, error)
	WriteByte(c byte) error
	WriteRune(r rune) (int, error)
	WriteString(s string) (int, error)
	WriteText(t T) (int, error)
}

type Reader[T Text] interface {
	Len() int
	Read(b []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
	ReadByte() (byte, error)
	ReadRune() (ch rune, size int, err error)
	Reset(t T)
	Seek(offset int64, whence int) (int64, error)
	Size() int64
	UnreadByte() error
	UnreadRune() error
	WriteTo(w io.Writer) (n int64, err error)
}

type Algorithms[T Text] interface {
	Clone(s T) T
	Compare(a, b T) int
	Concat(a, b T) T
	Contains(s, substr T) bool
	ContainsAny(s T, chars string) bool
	ContainsRune(s T, r rune) bool
	Count(s, substr T) int
	Cut(s, sep T) (before, after T, found bool)
	CutPrefix(s, prefix T) (after T, found bool)
	CutSuffix(s, suffix T) (before T, found bool)
	EqualFold(s, t T) bool
	Fields(s T) []T
	FieldsFunc(s T, f func(rune) bool) []T
	HasPrefix(s, prefix T) bool
	HasSuffix(s, suffix T) bool
	Index(s, substr T) int
	IndexAny(s T, chars string) int
	IndexByte(s T, c byte) int
	IndexFunc(s T, f func(rune) bool) int
	IndexRune(s T, r rune) int
	Join(elems []T, sep T) T
	LastIndex(s, substr T) int
	LastIndexAny(s T, chars string) int
	LastIndexByte(s T, c byte) int
	LastIndexFunc(s T, f func(rune) bool) int
	Map(mapping func(rune) rune, s T) T
	Repeat(s T, count int) T
	Replace(s, old, new T, n int) T
	ReplaceAll(s, old, new T) T
	Split(s, sep T) []T
	SplitAfter(s, sep T) []T
	SplitAfterN(s, sep T, n int) []T
	SplitN(s, sep T, n int) []T
	Title(s T) T
	ToLower(s T) T
	ToLowerSpecial(c unicode.SpecialCase, s T) T
	ToRunes(s T) []rune
	ToText(r []rune) T
	ToTitle(s T) T
	ToTitleSpecial(c unicode.SpecialCase, s T) T
	ToUpper(s T) T
	ToUpperSpecial(c unicode.SpecialCase, s T) T
	ToValidUTF8(s, replacement T) T
	Trim(s T, cutset string) T
	TrimFunc(s T, f func(rune) bool) T
	TrimLeft(s T, cutset string) T
	TrimLeftFunc(s T, f func(rune) bool) T
	TrimPrefix(s, prefix T) T
	TrimRight(s T, cutset string) T
	TrimRightFunc(s T, f func(rune) bool) T
	TrimSpace(s T) T
	TrimSuffix(s, suffix T) T

	UTF8DecodeRune(t T) (rune, int)
	UTF8RuneCount(t T) int
}

type Package[T Text, BuilderT Builder[T], ReaderT Reader[T]] interface {
	Algorithms[T]

	NewBuilder() BuilderT
	NewReader(t T) ReaderT
}
