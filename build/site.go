// Package build provides all functionality required for performing a
// build from reading content files to modelling the static site.
package build

import (
	"github.com/dominikbraun/espresso/model"
	"os"
	"strings"
)

// Site represents the actual website. It is a generic data model that
// holds all components and pages and can be rendered to a static site.
type Site struct {
	nav    *model.Nav
	root   route
	footer *model.Footer
}

// route represents a website route. Each route can have multiple pages
// associated with it, as well as multiple child routes. For example, a
// website route like /blog/my-category can be represented as:
//
//	"blog" {
//		children:
//			"my-category" {
//				pages: ...
//			}
//	}
//
// The root field of Site is considered as the root route that holds
// all sub-routes: "/blog" would be a child route of the site's root.
type route struct {
	pages    []*model.ArticlePage
	children map[string]*route
}

// newSite creates and initializes a new Site instance.
func newSite() *Site {
	s := Site{
		root: route{
			pages:    make([]*model.ArticlePage, 0),
			children: make(map[string]*route),
		},
	}
	return &s
}

// registerPage registers a given page under the route (path) that is
// stored in page.Path. This path must not end with a trailing slash.
func (s *Site) registerPage(page *model.ArticlePage) {
	node := &s.root
	segments := strings.Split(page.Path, string(os.PathSeparator))

	for i, seg := range segments {
		if _, exists := node.children[seg]; !exists {
			node.children[seg] = &route{
				pages:    make([]*model.ArticlePage, 0),
				children: make(map[string]*route),
			}
		}
		if i == len(segments)-1 {
			node.children[seg].pages = append(node.children[seg].pages, page)
			break
		}
		node = node.children[seg]
	}
}

// walkRoutes walks all site routes recursively and invokes a function
// for each route. depth specifies the maximal depth that the route tree
// will be walked down. Use -1 to walk down to the lowest level.
func (s *Site) walkRoutes(walkFn func(r *route), depth int) {
	s.walkRoute(&s.root, walkFn, depth, 0)
}

// walkRoute is used internally by walkRoutes and should not be called
// by other functions. It is the actual implementation of walkRoutes.
func (s *Site) walkRoute(route *route, walkFn func(r *route), depth int, currentDepth int) {
	if depth != -1 && currentDepth == depth {
		return
	}
	currentDepth++

	for _, route := range route.children {
		walkFn(route)
		s.walkRoute(route, walkFn, depth, currentDepth)
	}
}
