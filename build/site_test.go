// Package build provides all functionality required for performing a
// build from reading content files to modelling the static site.
package build

import (
	"github.com/dominikbraun/espresso/model"
	"testing"
)

// TestRegisterPage tests Site.registerPage. This test will register
// two custom pages and check if the routes have been created.
func TestRegisterPage(t *testing.T) {
	site := newSite()

	site.registerPage(&model.ArticlePage{
		Page:    model.Page{Path: "content/coffee"},
		Article: model.Article{ID: "making-barista-quality-espresso"},
	})

	site.registerPage(&model.ArticlePage{
		Page:    model.Page{Path: "content"},
		Article: model.Article{ID: "about-me"},
	})

	if _, ok := site.root.Children["content"]; !ok {
		t.Errorf("Could not find %s segment", "content")
	}

	if _, ok := site.root.Children["content"].Children["coffee"]; !ok {
		t.Errorf("Could not find %s segment", "coffee")
	}

	if len(site.root.Children["content"].Children["coffee"].Pages) < 1 {
		t.Errorf("Could not find page under %s", "content/coffee")
	}
}
