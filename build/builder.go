// Package build provides all functionality required for performing a
// build from reading content files to modelling the static site.
package build

import (
	"github.com/dominikbraun/espresso/config"
	"github.com/dominikbraun/espresso/model"
	"path/filepath"
	"strings"
	"sync"
)

// registerMode specified whether a freshly built page should be
// registered directly in the site model directly or not.
type registerMode uint

const (
	NoRegister     registerMode = 0
	DirectRegister registerMode = 1
)

// builder is the type used for performing the actual build. It knows
// the current build context and generates the entire site model which
// can then be rendered to a static site.
type builder struct {
	ctx   Context
	model *Site
	mutex *sync.Mutex
}

// newBuilder creates a builder instance that utilizes the build context.
func newBuilder(ctx Context) *builder {
	b := builder{
		ctx:   ctx,
		model: newSite(),
		mutex: &sync.Mutex{},
	}
	return &b
}

// buildPage attempts to generate a model.Page from a []byte. This is
// done by parsing its contents and building a page model which can be
// registered afterwards. The file parameter indicates the original file
// path.
//
// buildPage is safe for concurrent invocation. The file path must contain
// the build path.
func (b *builder) buildPage(source []byte, file string, mode registerMode) (*model.ArticlePage, error) {
	article, err := b.ctx.Parser.ParseArticle(source)
	if err != nil {
		return nil, err
	}

	// If the file path is `my-site/content/blog/category-1/post.md`, its
	// relative path within the build path is `/blog/category-1/post.md`.
	// contentDirLen marks the starting point for the relative path.
	contentDirLen := len(b.ctx.BuildPath + "/" + config.ContentDir)

	// If the build path is `.`, the file path is probably simplified to
	// `content/blog/category-1/post.md`. This edge case is handled here.
	if b.ctx.BuildPath == "." && file[:len(config.ContentDir)] == config.ContentDir {
		contentDirLen = len(config.ContentDir)
	}

	// Remove the build path and content dir to get the relative path.
	relativePath := file[contentDirLen:]

	route := filepath.ToSlash(filepath.Dir(relativePath))
	article.ID = strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))

	page := &model.ArticlePage{
		Page: model.Page{
			Path: route,
		},
		Article: article,
	}

	// The article ID is `index` if an `index.md` file has been created
	// by the user. In this case, it will be registered as an IndexPage.
	if article.ID == "index" {
		b.registerIndexPage(page)
	} else {
		b.registerPage(page)
	}

	return page, nil
}

// registerPage registers a page model to the builder's site model.
// This page model can be retrieved using buildPage for instance.
//
// registerPage is safe for concurrent invocation.
func (b *builder) registerPage(page *model.ArticlePage) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.model.registerPage(page)
}

// registerIndexPage registers an index page to the builder's site
// model for the page's route.
//
// registerPage is safe for concurrent invocation.
func (b *builder) registerIndexPage(indexPage *model.ArticlePage) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.model.registerIndexPage(indexPage)
}

// buildNav attempts to create a model.Nav from the existing pages that
// have to be built and registered first, meaning that buildNav must be
// called after all buildPage calls have finished.
//
// buildNav takes the site settings into account and overrides the Nav
// if this is specified in the site settings.
func (b *builder) buildNav() error {
	nav := &model.Nav{
		Brand: b.ctx.Settings.Name,
		Items: make([]model.NavItem, 0),
	}

	for _, i := range b.ctx.Settings.Nav.Items {
		item := model.NavItem{
			Label:  i.Label,
			Target: i.Target,
		}
		nav.Items = append(nav.Items, item)
	}

	if !b.ctx.Settings.Nav.Override {
		b.model.WalkRoutes(func(r string, i *RouteInfo) {
			segments := strings.Split(r, "/")
			// Only process routes with a depth of 1.
			if len(segments) == 1 {
				item := model.NavItem{
					Label:  strings.Title(segments[0]),
					Target: segments[0],
				}
				nav.Items = append(nav.Items, item)
			}
		})
	}

	b.model.Nav = nav
	return nil
}

// buildListPages attempts to build overview pages for all categories.
// For each route in the route tree, all articles are added to the
// routes's list page model.
func (b *builder) buildListPages() error {
	b.model.
		WalkRoutes(func(r string, i *RouteInfo) {
			i.ListPage = &model.ArticleListPage{
				Page:     model.Page{Path: r},
				Articles: make([]*model.Article, len(i.Pages)),
			}

			for n, page := range i.Pages {
				i.ListPage.Articles[n] = &page.Article
			}
		})

	return nil
}

// buildRelated attempts to store all related articles for each article
// in the page tree. Storing a pointer to these related articles allows
// the user to access the fields of each article in his templates.
func (b *builder) buildRelated() error {
	b.model.
		WalkRoutes(func(r string, i *RouteInfo) {
			for _, p := range i.Pages {
				for _, link := range p.Article.RelatedLinks {
					// A `link` consists of an Espresso path like `/coffee`
					// and an article ID like `coffee-roasting-basics.md`.
					// These components are split here to resolve the path.
					path := link[:strings.LastIndex(link, "/")]
					id := link[len(path)+1:]

					// Load the page and its article by resolving its path.
					page, _ := b.model.resolvePath(path, id)
					page.Article.Related = append(page.Article.Related, &page.Article)
				}
			}
		})

	return nil
}

// buildFooter attempts to create a model.Footer under consideration of
// user-defined site settings. It is independent from any site pages.
func (b *builder) buildFooter() error {
	footer := &model.Footer{
		Text:  b.ctx.Settings.Footer.Text,
		Items: make([]model.FooterItem, 0),
	}

	for _, i := range b.ctx.Settings.Footer.Items {
		item := model.FooterItem{
			Label:  i.Label,
			Target: i.Target,
		}
		footer.Items = append(footer.Items, item)
	}

	b.model.Footer = footer
	return nil
}
