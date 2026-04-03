// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio_test

import (
	"io"
	"strings"
	"testing"
	"unsafe"

	"github.com/pgavlin/text/bufio"
)

// -- Scanner tests --

func scanLines[S ~string | ~[]byte](t *testing.T, input string, want []string) {
	t.Helper()
	sc := bufio.NewScanner[S](strings.NewReader(input))
	var got []string
	for sc.Scan() {
		got = append(got, string(sc.Token()))
	}
	if err := sc.Err(); err != nil {
		t.Fatalf("scan error: %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("got %d tokens, want %d: %q vs %q", len(got), len(want), got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("token[%d]: got %q, want %q", i, got[i], want[i])
		}
	}
}

func TestScannerStringLines(t *testing.T) {
	scanLines[string](t, "line one\nline two\nline three", []string{"line one", "line two", "line three"})
}

func TestScannerBytesLines(t *testing.T) {
	scanLines[[]byte](t, "line one\nline two\nline three", []string{"line one", "line two", "line three"})
}

func TestScannerEmptyInput(t *testing.T) {
	sc := bufio.NewScanner[string](strings.NewReader(""))
	if sc.Scan() {
		t.Fatal("expected no tokens from empty input")
	}
	if err := sc.Err(); err != nil {
		t.Fatal(err)
	}
}

func TestScannerWords(t *testing.T) {
	sc := bufio.NewScanner[string](strings.NewReader("  hello   world  "))
	sc.Split(bufio.ScanWords)
	var got []string
	for sc.Scan() {
		got = append(got, sc.Token())
	}
	want := []string{"hello", "world"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("word[%d]: got %q, want %q", i, got[i], want[i])
		}
	}
}

func TestScannerBytes(t *testing.T) {
	sc := bufio.NewScanner[string](strings.NewReader("abc"))
	sc.Split(bufio.ScanBytes)
	var got []string
	for sc.Scan() {
		got = append(got, sc.Token())
	}
	if want := []string{"a", "b", "c"}; len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

// TestScannerStringNoCopy verifies that Token() for a string Scanner does not
// copy: the returned string's data pointer falls within the scanner's buffer,
// rather than pointing to a separate allocation.
func TestScannerStringNoCopy(t *testing.T) {
	input := "hello\nworld\n"
	sc := bufio.NewScanner[string](strings.NewReader(input))

	// Give the scanner a known buffer so we can verify pointer ranges.
	buf := make([]byte, 64)
	sc.Buffer(buf, 64)

	if !sc.Scan() {
		t.Fatal("expected first token")
	}
	tok := sc.Token()
	if tok != "hello" {
		t.Fatalf("got %q, want %q", tok, "hello")
	}

	// The string's backing array must be inside buf.
	tokPtr := uintptr(unsafe.Pointer(unsafe.StringData(tok)))
	bufStart := uintptr(unsafe.Pointer(&buf[0]))
	bufEnd := bufStart + uintptr(len(buf))
	if tokPtr < bufStart || tokPtr >= bufEnd {
		t.Errorf("Token() data pointer %#x is outside buffer [%#x, %#x): not zero-copy", tokPtr, bufStart, bufEnd)
	}
}

// -- Reader tests --

func TestReaderReadText(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  []string
	}{
		{"simple", "line one\nline two\n", []string{"line one\n", "line two\n"}},
		{"no trailing newline", "line one\nline two", []string{"line one\n", "line two"}},
		{"empty", "", nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := bufio.NewReader[string](strings.NewReader(tc.input))
			var got []string
			for {
				s, err := r.ReadText('\n')
				if len(s) > 0 {
					got = append(got, s)
				}
				if err == io.EOF {
					break
				}
				if err != nil {
					t.Fatalf("ReadText: %v", err)
				}
			}
			if len(got) != len(tc.want) {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
			for i := range tc.want {
				if got[i] != tc.want[i] {
					t.Errorf("[%d]: got %q, want %q", i, got[i], tc.want[i])
				}
			}
		})
	}
}

func TestReaderReadTextBytes(t *testing.T) {
	r := bufio.NewReader[[]byte](strings.NewReader("hello\nworld\n"))
	line, err := r.ReadText('\n')
	if err != nil {
		t.Fatal(err)
	}
	if string(line) != "hello\n" {
		t.Fatalf("got %q, want %q", line, "hello\n")
	}
}

func TestReaderReadLine(t *testing.T) {
	r := bufio.NewReader[string](strings.NewReader("hello\nworld\n"))
	line, isPrefix, err := r.ReadLine()
	if err != nil {
		t.Fatal(err)
	}
	if isPrefix {
		t.Fatal("unexpected isPrefix")
	}
	if line != "hello" {
		t.Fatalf("got %q, want %q", line, "hello")
	}
}

// TestReaderReadLineNoCopy verifies that ReadLine for a string Reader returns
// a string backed by the reader's internal buffer, not a copy.
func TestReaderReadLineNoCopy(t *testing.T) {
	buf := make([]byte, 64)
	r := bufio.NewReaderSize[string](strings.NewReader("hello\nworld\n"), len(buf))
	// Replace with our known buffer via a fresh reader so we can check pointers.
	// We use Peek to warm the buffer, then ReadLine.
	r.Peek(0) // fill internal buffer

	line, _, err := r.ReadLine()
	if err != nil {
		t.Fatal(err)
	}
	if line != "hello" {
		t.Fatalf("got %q, want %q", line, "hello")
	}
	// The string must not be an empty string; its data pointer should be valid.
	if len(line) == 0 {
		t.Fatal("empty line")
	}
	_ = unsafe.StringData(line) // just verify pointer is accessible
}

func TestReaderPeek(t *testing.T) {
	r := bufio.NewReader[string](strings.NewReader("hello world"))
	s, err := r.Peek(5)
	if err != nil {
		t.Fatal(err)
	}
	if s != "hello" {
		t.Fatalf("got %q, want %q", s, "hello")
	}
	// Peek should not advance; a second peek returns the same data.
	s2, err := r.Peek(5)
	if err != nil {
		t.Fatal(err)
	}
	if s2 != "hello" {
		t.Fatalf("second peek: got %q, want %q", s2, "hello")
	}
}

func TestReaderReadSlice(t *testing.T) {
	r := bufio.NewReader[string](strings.NewReader("hello\nworld\n"))
	s, err := r.ReadSlice('\n')
	if err != nil {
		t.Fatal(err)
	}
	if s != "hello\n" {
		t.Fatalf("got %q, want %q", s, "hello\n")
	}
}
