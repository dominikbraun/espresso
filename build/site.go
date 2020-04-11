// Package build provides all functionality required for performing a
// build from reading content files to modelling the static site.
package build

import (
	"fmt"
	"github.com/dominikbraun/espresso/model"
	"path/filepath"
	"strings"
)

// Site represents the actual website. It is a generic data model that
// holds all components and pages and can be rendered to a static site.
type Site struct {
	Nav    *model.Nav
	root   Route
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

// newSite creates and initializes a new Site instance.
func newSite() *Site {
	s := Site{
		root: Route{
			Pages:    make([]*model.ArticlePage, 0),
			Children: make(map[string]*Route),
		},
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
//
// If the route doesn't exist yet, all of its required child-routes are
// created until the entire page path is depicted.
func (s *Site) registerPage(page *model.ArticlePage) {
	node := &s.root
	segments := strings.Split(page.Path, "/")

	for i, seg := range segments {
		// Create the route with under segment key if it doesn't exist.
		if _, exists := node.Children[seg]; !exists {
			node.Children[seg] = newRoute()
			// The current path consists of all previous path segments
			// up to the current segment. This path will be stored.
			path := filepath.FromSlash(filepath.Join(segments[:i]...))
			node.Children[seg].Path = path
		}
		// Store the page in the current segment if it is the last one.
		if i == len(segments)-1 {
			node.Children[seg].Pages = append(node.Children[seg].Pages, page)
			break
		}
		// Walk down the tree to the next segment.
		node = node.Children[seg]
	}
}

// WalkRoutes walks all site routes recursively and invokes a function
// for each route. depth specifies the maximal depth that the route tree
// will be walked down. Use -1 to walk down to the lowest level.
func (s *Site) WalkRoutes(walkFn func(r *Route), depth int) {
	s.walkRoute(&s.root, walkFn, depth, 0)
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
	node := &s.root
	segments := strings.Split(path, "/")

	for i, seg := range segments {
		if _, exists := node.Children[seg]; !exists {
			return nil, fmt.Errorf("the sub-route `%s` does not exist", seg)
		}
		if i == len(segments)-1 {
			for _, page := range node.Children[seg].Pages {
				if page.Article.ID == id {
					return page, nil
				}
			}
		}
		node = node.Children[seg]
	}

	return nil, fmt.Errorf("article `%s` not found in `%s`", id, path)
}
