// Package render provides functions for rendering the actual website.
package render

import (
	"github.com/dominikbraun/espresso/build"
	"github.com/dominikbraun/espresso/config"
	"github.com/dominikbraun/espresso/filesystem"
	"github.com/dominikbraun/espresso/model"
	"github.com/dominikbraun/espresso/template"
	"path/filepath"
	"sync"
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
		var wg sync.WaitGroup

		for _, page := range r.Pages {
			wg.Add(1)
			renderArticlePage(&ctx, page, &wg)
		}
		_ = renderArticleListPage(&ctx, r.ListPage)
	}, -1)

	assetTarget := filepath.Join(ctx.TargetDir, config.AssetDir)
	if err := filesystem.CopyDir(ctx.AssetDir, assetTarget); err != nil {
		return err
	}

	return nil
}

// renderArticlePage renders a given ArticlePage as an HTML file.
func renderArticlePage(ctx *Context, page *model.ArticlePage, wg *sync.WaitGroup) {
	defer wg.Done()
	pagePath := filepath.Join(ctx.TargetDir, page.Path, page.Article.ID)

	if err := renderPage(ctx, pagePath, template.Article, page); err != nil {
		// Handle error by sending it through a channel or so.
	}
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
