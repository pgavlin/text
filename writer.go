package text

import (
	"io"

	"github.com/pgavlin/text/internal/bytealg"
)

// Writer is the interface that wraps the WriteText method.
type Writer[S String] interface {
	WriteText(s S) (int, error)
}

type asWriter[S String] struct {
	w io.StringWriter
}

func (w asWriter[S]) WriteText(s S) (int, error) {
	return w.w.WriteString(bytealg.AsString(s))
}

// AsWriter projects an io.StringWriter as a Writer[S].
func AsWriter[S String](w io.StringWriter) Writer[S] {
	return asWriter[S]{w: w}
}
