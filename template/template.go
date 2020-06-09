// Package template provides functions for rendering Go templates.
package template

import (
	"io"
	"text/template"
)

const (
	Article   string = "article.html"
	ListPage  string = "list-page.html"
	IndexPage string = "index-page.html"
)

// Render attempts to render the specified template file, populate
// it with the provided data and write it to a io.Writer target.
func Render(file string, data interface{}, target io.Writer) error {
	tpl, err := template.ParseFiles(file)
	if err != nil {
		return err
	}

	return tpl.Execute(target, data)
}
