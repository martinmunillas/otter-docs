package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/martinmunillas/otter-docs/component"
	"github.com/martinmunillas/otter-docs/component/markdown"
	"github.com/martinmunillas/otter-docs/envs"
	"github.com/martinmunillas/otter/i18n"
	"github.com/martinmunillas/otter/server"
)

func GetIndex() server.Handler {
	return func(r *http.Request, t server.Tools) {
		t.Send.Ok.HTML(component.IndexPage(t))
	}
}

func GetDocs() server.Handler {
	return func(r *http.Request, t server.Tools) {
		if envs.Etag != "" {
			if match := r.Header.Get("If-None-Match"); match == envs.Etag {
				t.Send.NotModified()
				return
			}
		}
		path := r.URL.Path[5:]
		if path == "/" {
			path = "index"
		}
		f, err := os.ReadFile(fmt.Sprintf("component/content/%s/%s.md", i18n.FromCtx(r.Context()), path))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				t.Send.NotFound.HTML(component.DocsPage(t.Translation("notFound"), nil))
				return
			}
			t.Send.InternalError.HTML(err, component.DocsPage(t.Translation("thereWasAnError"), nil))
			return
		}
		if envs.Etag != "" {
			t.AddHeader("Cache-control", "no-cache")
			t.AddHeader("Etag", envs.Etag)
		}
		t.Send.Ok.HTML(component.DocsPage(markdown.Parse(f), markdown.GetHeadings(f)))
	}
}
