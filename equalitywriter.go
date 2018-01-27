package equality

import (
	"io"
	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewEqualityWriter())
	if err != nil {
		panic(err)
	}
}

type EqualityWriter struct{}

func NewEqualityWriter() *EqualityWriter {
	return &EqualityWriter{}
}

func (this *EqualityWriter) Name() string {
	return "equality"
}

func (this *EqualityWriter) Imports(t typewriter.Type) (result []typewriter.ImportSpec) {
	return result
}

func (this *EqualityWriter) Write(w io.Writer, t typewriter.Type) error {
	tag, found := t.FindTag(this)

	if !found {
		return nil
	}

	tmpl, err := templates.ByTag(t, tag)

	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, t); err != nil {
		return err
	}

	return nil
}
