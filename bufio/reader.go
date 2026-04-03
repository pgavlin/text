// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

import (
	stdbufio "bufio"
	"io"

	"github.com/pgavlin/text"
)

// Reader[S] implements buffered I/O for an io.Reader. It is a generic wrapper
// around [bufio.Reader]. Methods that return text — Peek, ReadSlice, ReadLine,
// and ReadText — return type S rather than []byte or string, allowing callers
// to receive data in their preferred representation without a copy.
//
// When S is string, methods that return a buffer-backed slice (Peek, ReadSlice,
// ReadLine) return a string whose data is only valid until the next read
// operation, matching the same contract as the underlying []byte methods.
// ReadText always returns independently allocated data that is safe to retain.
type Reader[S text.String] struct {
	r *stdbufio.Reader
}

// NewReader returns a new Reader whose buffer has the default size.
func NewReader[S text.String](rd io.Reader) *Reader[S] {
	return &Reader[S]{r: stdbufio.NewReader(rd)}
}

// NewReaderSize returns a new Reader whose buffer has at least the specified
// size. If the argument io.Reader is already a Reader with large enough size,
// it returns the underlying Reader.
func NewReaderSize[S text.String](rd io.Reader, size int) *Reader[S] {
	return &Reader[S]{r: stdbufio.NewReaderSize(rd, size)}
}

// Size returns the size of the underlying buffer in bytes.
func (r *Reader[S]) Size() int { return r.r.Size() }

// Reset discards any buffered data, resets all state, and switches the
// buffered reader to read from rd.
func (r *Reader[S]) Reset(rd io.Reader) { r.r.Reset(rd) }

// Buffered returns the number of bytes that can be read from the current buffer.
func (r *Reader[S]) Buffered() int { return r.r.Buffered() }

// Discard skips the next n bytes, returning the number of bytes discarded.
func (r *Reader[S]) Discard(n int) (discarded int, err error) { return r.r.Discard(n) }

// Read reads data into p. It returns the number of bytes read into p.
func (r *Reader[S]) Read(p []byte) (n int, err error) { return r.r.Read(p) }

// ReadByte reads and returns a single byte. If no byte is available,
// returns an error.
func (r *Reader[S]) ReadByte() (byte, error) { return r.r.ReadByte() }

// UnreadByte unreads the last byte. Only the most recently read byte can be unread.
func (r *Reader[S]) UnreadByte() error { return r.r.UnreadByte() }

// ReadRune reads a single UTF-8 encoded Unicode character and returns the
// rune and its size in bytes.
func (r *Reader[S]) ReadRune() (ch rune, size int, err error) { return r.r.ReadRune() }

// UnreadRune unreads the last rune. If the most recent method called on the
// buffer was not a successful ReadRune, UnreadRune returns an error.
func (r *Reader[S]) UnreadRune() error { return r.r.UnreadRune() }

// WriteTo implements io.WriterTo.
func (r *Reader[S]) WriteTo(w io.Writer) (n int64, err error) { return r.r.WriteTo(w) }

// Peek returns the next n bytes without advancing the reader. The returned
// value is only valid until the next read call. If Peek returns fewer than n
// bytes, it also returns an error explaining why the read is short.
func (r *Reader[S]) Peek(n int) (S, error) {
	b, err := r.r.Peek(n)
	return asS[S](b), err
}

// ReadSlice reads until the first occurrence of delim in the input, returning
// a slice pointing to bytes in the buffer. The returned value is only valid
// until the next read. If ReadSlice encounters the end of the buffer without
// finding delim, it returns ErrBufferFull. For most uses, ReadText or Scanner
// are more convenient.
func (r *Reader[S]) ReadSlice(delim byte) (S, error) {
	b, err := r.r.ReadSlice(delim)
	return asS[S](b), err
}

// ReadLine is a low-level line-reading primitive. Most callers should use
// ReadText or Scanner instead.
//
// ReadLine tries to return a single line, not including the end-of-line bytes.
// The returned value is only valid until the next call to ReadLine. ReadLine
// either returns a non-empty line or it returns an error, never both.
//
// The text returned from ReadLine does not include the line end ("\r\n" or
// "\n"). No indication or error is given if the input ends without a final
// line end. If the line was too long for the buffer then isPrefix is set and
// the beginning of the line is returned. The rest of the line will be
// returned from future calls. isPrefix will be false when returning the last
// fragment of the line.
func (r *Reader[S]) ReadLine() (line S, isPrefix bool, err error) {
	b, isPrefix, err := r.r.ReadLine()
	return asS[S](b), isPrefix, err
}

// ReadText reads until the first occurrence of delim in the input, returning
// a string containing the data up to and including the delimiter. If ReadText
// encounters an error before finding a delimiter, it returns the data read
// before the error and the error itself (often io.EOF). ReadText returns
// err != nil if and only if the returned data does not end in delim. For
// simple uses, a Scanner may be more convenient.
//
// Unlike ReadSlice and ReadLine, the returned value is safe to retain
// indefinitely; it does not alias the reader's internal buffer.
func (r *Reader[S]) ReadText(delim byte) (S, error) {
	b, err := r.r.ReadBytes(delim)
	return asS[S](b), err
}
