// Package build provides all functionality required for performing a
// build from reading content files to modelling the static site.
package build

import (
	"fmt"
	"github.com/dominikbraun/espresso/model"
)

// Site represents the actual website. It is a generic data model that
// holds all components and pages and can be rendered to a static site.
type Site struct {
	Nav    *model.Nav
	routes map[string]*RouteInfo
	Footer *model.Footer
}

// RouteInfo holds the contents of a route. While the route itself is
// a map key, the value of such a map is a RouteInfo instance. Each
// route contains multiple pages as well as a summarizing page.
//
// If the user provides their own `index.md` file, this page will be
// built as IndexPage and ListPage is nil. Otherwise, ListPage will be
// generated automatically and IndexPage is nil.
type RouteInfo struct {
	Pages     []*model.ArticlePage
	ListPage  *model.ArticleListPage
	IndexPage *model.IndexPage
}

// newSite creates and initializes a new Site instance.
func newSite() *Site {
	s := Site{
		routes: make(map[string]*RouteInfo),
	}
	return &s
}

// registerPage registers a given page under the route (path) that is
// stored in page.Path. This path must not end with a trailing slash.
func (s *Site) registerPage(page *model.ArticlePage) {
	if _, exists := s.routes[page.Path]; !exists {
		s.routes[page.Path] = &RouteInfo{
			Pages: make([]*model.ArticlePage, 0),
		}
	}

	s.routes[page.Path].Pages = append(s.routes[page.Path].Pages, page)
}

// registerIndexPage registers an index page when it has been provided
// by the user and built by Espresso.
func (s *Site) registerIndexPage(indexPage *model.IndexPage) {
	if _, exists := s.routes[indexPage.Path]; !exists {
		s.routes[indexPage.Path] = &RouteInfo{
			Pages: make([]*model.ArticlePage, 0),
		}
	}

	s.routes[indexPage.Path].IndexPage = indexPage
}

// WalkRoutes iterates over all routes of the site model and invokes a
// function walkFn on each route, with r being the route string and i
// being the RouteInfo instance holding all pages and the list page.
//
// For filtering the routes by their depth, split r by a `/` and count
// the slice length. To only process routes with a maximum depth of 2
// for example, perform a check like this:
//
//	site.WalkRoutes(func(r string, i *RouteInfo) {
//		if len(strings.Split(r, "/")) > 2 {
//			return
//		}
//		// Process route
//	})
//
// This code will only process routes with a depth of 2 or less.
func (s *Site) WalkRoutes(walkFn func(r string, i *RouteInfo)) {
	for route, info := range s.routes {
		walkFn(route, info)
	}
}

// resolvePath resolves a given path in the route tree and returns the
// page with the given ID in that path. Returns an error if the either
// the path or the article ID does not exist.
func (s *Site) resolvePath(path string, id string) (*model.ArticlePage, error) {
	if info, exists := s.routes[path]; exists {
		for _, p := range info.Pages {
			if p.Article.ID == id {
				return p, nil
			}
		}
	}

	return nil, fmt.Errorf("article `%s` not found in `%s`", id, path)
}
