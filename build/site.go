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

// Route represents a website Route. Each Route can have multiple pages
// associated with it, as well as multiple child routes. For example, a
// website route like /blog/my-category can be represented as:
//
//	"blog" {
//		Children:
//			"my-category" {
//				Pages: ...
//			}
//	}
//
// The root field of Site is considered as the root route that holds all
// sub-routes: "/blog" would be a child route of the site's root.
type Route struct {
	// Path is a convenience field for storing the route's absolute path.
	Path     string
	Pages    []*model.ArticlePage
	ListPage *model.ArticleListPage
	Children map[string]*Route
}

type RouteInfo struct {
	Pages    []*model.ArticlePage
	ListPage *model.ArticleListPage
}

// newSite creates and initializes a new Site instance.
func newSite() *Site {
	s := Site{
		routes: make(map[string]*RouteInfo),
	}
	return &s
}

// newRoute creates and initializes a new Route instance.
func newRoute() *Route {
	r := Route{
		Pages: make([]*model.ArticlePage, 0),
		ListPage: &model.ArticleListPage{
			Page:     model.Page{},
			Articles: make([]*model.Article, 0),
		},
		Children: make(map[string]*Route),
	}
	return &r
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

// WalkRoutes walks all site routes recursively and invokes a function
// for each route. depth specifies the maximal depth that the route tree
// will be walked down. Use -1 to walk down to the lowest level.
func (s *Site) WalkRoutes(walkFn func(r string, i *RouteInfo)) {
	for route, info := range s.routes {
		walkFn(route, info)
	}
}

// walkRoute is used internally by WalkRoutes and should not be called
// by other functions. It is the actual implementation of WalkRoutes.
func (s *Site) walkRoute(route *Route, walkFn func(r *Route), depth int, currentDepth int) {
	if depth != -1 && currentDepth == depth {
		return
	}
	currentDepth++

	for _, route := range route.Children {
		walkFn(route)
		s.walkRoute(route, walkFn, depth, currentDepth)
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
