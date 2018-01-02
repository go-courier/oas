package oas

import (
	"testing"
)

func TestSpecExtensions(t *testing.T) {
	g := NewCaseGroup("SpecExtensions")

	g.It("empty", `{}`, SpecExtensions{})

	g.It("with extensions", `{"x-a":"xxx"}`, func() *SpecExtensions {
		e := &SpecExtensions{}
		e.AddExtension("x-b", nil)
		e.AddExtension("x-a", "xxx")
		return e
	}())

	g.Run(t)
}
