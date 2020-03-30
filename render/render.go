// Package render provides functions for rendering the actual website.
package render

import (
	"github.com/dominikbraun/espresso/build"
	"github.com/dominikbraun/espresso/filesystem"
	"github.com/dominikbraun/espresso/model"
	"github.com/dominikbraun/espresso/template"
	"path/filepath"
)

const (
	indexFile string = "index.html"
)

// Context represents the rendering context and holds data required
// for a particular rendering process such as the target directory.
type Context struct {
	TemplateDir string
	AssetDir    string
	TargetDir   string
}

// AsWebsite starts rendering the site model as an HTML-base site.
func AsWebsite(ctx Context, site *build.Site) error {
	site.WalkRoutes(func(r *build.Route) {
		for _, page := range r.Pages {
			_ = renderPage(&ctx, page)
		}
	}, -1)

	return nil
}

// renderPage renders a particular page model using the respective
// template for the page type. The directory where the template will
// be rendered to is determined by the page path.
func renderPage(ctx *Context, page *model.ArticlePage) error {
	pagePath := filepath.Join(ctx.TargetDir, page.Path, page.ID)

	if err := filesystem.CreateDir(pagePath, true); err != nil {
		return err
	}

	handle, err := filesystem.CreateFile(filepath.Join(pagePath, indexFile))
	if err != nil {
		return err
	}

	tplPath := filepath.Join(ctx.TemplateDir, template.Article)

	return template.Render(tplPath, page.Article, handle)
}
