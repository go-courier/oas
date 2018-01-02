package oas

import (
	"net/url"
	"testing"
)

func TestTag(t *testing.T) {
	g := NewCaseGroup("Tag")

	g.It("empty", `{"name":""}`, Tag{})

	g.It("with external docs", `{"name":"tag","externalDocs":{"description":"google","url":"//google.com"}}`, func() *Tag {
		t := NewTag("tag")
		t.ExternalDocs = NewExternalDoc((&url.URL{Host: "google.com"}).String(), "google")
		return t
	}())

	g.Run(t)
}
