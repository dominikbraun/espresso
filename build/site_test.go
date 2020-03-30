// Package build provides all functionality required for performing a
// build from reading content files to modelling the static site.
package build

import (
	"github.com/dominikbraun/espresso/model"
	"testing"
)

func TestRegisterPage(t *testing.T) {
	site := newSite()

	site.registerPage(&model.ArticlePage{
		Page: model.Page{
			Path: "content/coffee",
			ID:   "making-barista-quality-espresso",
		},
		Article: model.Article{},
	})

	site.registerPage(&model.ArticlePage{
		Page: model.Page{
			Path: "content",
			ID:   "about-me",
		},
		Article: model.Article{},
	})

	if _, ok := site.root.children["content"]; !ok {
		t.Errorf("Could not find %s segment", "content")
	}

	if _, ok := site.root.children["content"].children["coffee"]; !ok {
		t.Errorf("Could not find %s segment", "coffee")
	}

	if len(site.root.children["content"].children["coffee"].pages) < 1 {
		t.Errorf("Could not find page under %s", "content/coffee")
	}
}
