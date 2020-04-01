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
			_ = renderArticlePage(&ctx, page)
		}
	}, -1)

	return nil
}

// renderArticlePage renders a given ArticlePage as an HTML file.
func renderArticlePage(ctx *Context, page *model.ArticlePage) error {
	pagePath := filepath.Join(ctx.TargetDir, page.Path, page.ID)

	if err := renderPage(ctx, pagePath, template.Article, page); err != nil {
		return err
	}

	return nil
}

// renderArticleListPage renders a given ArticleListPage as an HTML
// file.
func renderArticleListPage(ctx *Context, page *model.ArticleListPage) error {
	pagePath := filepath.Join(ctx.TargetDir, page.Path)

	if err := renderPage(ctx, pagePath, template.ArticleList, page); err != nil {
		return err
	}

	return nil
}

// renderPage is the common function for rendering any page model. It
// creates a directory structure corresponding to the page path and
// renders the specified template file into this path as `index.html`.
func renderPage(ctx *Context, path, tpl string, data interface{}) error {
	if err := filesystem.CreateDir(path, true); err != nil {
		return err
	}

	handle, err := filesystem.CreateFile(filepath.Join(path, indexFile))
	if err != nil {
		return err
	}

	tplPath := filepath.Join(ctx.TemplateDir, tpl)

	return template.Render(tplPath, data, handle)
}
