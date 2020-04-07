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
	indexFile  string = "index.html"
	numWorkers int    = 8
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
	pages := make(chan *model.ArticlePage)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go processQueue(&ctx, pages, &wg)
	}

	go streamPages(&ctx, site, pages)
	wg.Wait()

	_ = copyAssetDir(&ctx)

	return nil
}

// streamPages walks down the route tree and sends all pages through
// the pages channel, which is used to receive and build these pages.
func streamPages(ctx *Context, site *build.Site, pages chan<- *model.ArticlePage) {
	site.WalkRoutes(func(r *build.Route) {
		for _, page := range r.Pages {
			pages <- page
		}
		_ = renderArticleListPage(ctx, r.ListPage)
	}, -1)

	close(pages)
}

// processQueue receives pages from the pages channel and processes
// them by invoking the renderArticlePage function.
func processQueue(ctx *Context, pages <-chan *model.ArticlePage, wg *sync.WaitGroup) {
	for page := range pages {
		_ = renderArticlePage(ctx, page)
	}
	wg.Done()
}

// renderArticlePage renders a given ArticlePage as an HTML file.
func renderArticlePage(ctx *Context, page *model.ArticlePage) error {
	pagePath := filepath.Join(ctx.TargetDir, page.Path, page.Article.ID)

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

// copyAssetDir copies the asset directory from the build path into the
// build target directory recursively.
func copyAssetDir(ctx *Context) error {
	assetTarget := filepath.Join(ctx.TargetDir, config.AssetDir)

	if err := filesystem.CopyDir(ctx.AssetDir, assetTarget); err != nil {
		return err
	}

	return nil
}
