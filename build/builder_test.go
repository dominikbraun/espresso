// Package build provides all functionality required for performing a
// build from reading content files to modelling the static site.
package build

import (
	"github.com/dominikbraun/espresso/parser"
	"github.com/dominikbraun/espresso/settings"
	"testing"
)

// defaultBuilder creates a builder instance and initializes it with
// a generic build context that utilizes a Markdown parser.
func defaultBuilder() *builder {
	builder := newBuilder(Context{
		BuildPath: ".",
		Settings:  &settings.Site{},
		Parser:    parser.NewMarkdown(),
	})
	return builder
}

// fail lets a provided testing fail. The corresponding error message
// indicates the expected result and the actual value.
func fail(t *testing.T, run int, expected, got string) {
	t.Errorf("%v. test: Exptected %s, got %s", run, expected, got)
}

// TestBuildPages tests builder.buildPage. At the moment, the page path,
// page ID, article title and first article tag are tested.
func TestBuildPage(t *testing.T) {
	builder := defaultBuilder()

	testdata := []struct {
		source         string
		file           string
		expectedPath   string
		expectedID     string
		expectedTitle  string
		expected1stTag string
	}{
		{
			source: `---
Title: Making Barista-Quality Espresso
Tags:
    - Espresso
    - Coffee
---`,
			file:           "./content/coffee/making-barista-quality-espresso.md",
			expectedPath:   "content/coffee",
			expectedID:     "making-barista-quality-espresso",
			expectedTitle:  "Making Barista-Quality Espresso",
			expected1stTag: "Espresso",
		},
		{
			source: `---
Title: Coffee Roasting Basics
Date: 2020-03-30
Tags:
    - Coffee
    - Roasting
---`,
			file:           "./content/coffee/coffee-roasting-basics.md",
			expectedPath:   "content/coffee",
			expectedID:     "coffee-roasting-basics",
			expectedTitle:  "Coffee Roasting Basics",
			expected1stTag: "Coffee",
		},
		{
			source: `---
Title: About Me
Tags:
    - About
---`,
			file:           "./content/about-me.md",
			expectedPath:   "content",
			expectedID:     "about-me",
			expectedTitle:  "About Me",
			expected1stTag: "About",
		},
	}

	for i, test := range testdata {
		page, err := builder.buildPage([]byte(test.source), test.file)
		if err != nil {
			t.Error(err)
		}

		if page.Path != test.expectedPath {
			fail(t, i, test.expectedPath, page.Path)
		}

		if page.ID != test.expectedID {
			fail(t, i, test.expectedID, page.ID)
		}

		if page.Article.Title != test.expectedTitle {
			fail(t, i, test.expectedTitle, page.Article.Title)
		}

		if page.Article.Tags[0] != test.expected1stTag {
			fail(t, i, test.expected1stTag, page.Article.Tags[0])
		}
	}
}
