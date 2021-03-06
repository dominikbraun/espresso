// Package render provides functions for rendering the actual website.
package render

import "github.com/dominikbraun/espresso/model"

// Plugin represents a rendering plugin that will be invoked when
// rendering the site model.
type Plugin interface {
	ProcessArticlePage(ctx *Context, page *model.ArticlePage) error
	Finalize(ctx *Context) error
}
