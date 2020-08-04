// Package entity provides the Espresso domain models. They are
// used to represent the entire website and its components.
package entity

import (
	"fmt"
	"strings"
)

// Site represents the actual website and is thus the root entity
// which contains all website components.
//
// Each Espresso build aims to produce a completely populated Site
// entity which can be modified by plugins and used for rendering
// the website files.
type Site struct {
	Meta   Meta   // Meta contains page-agnostic, global metadata for the website itself.
	Nav    Nav    // Nav contains the global website navigation.
	Root   Route  // Root is the root element of the route tree.
	Footer Footer // Footer contains the global website footer.
}

// WalkRoutes walks down the route tree recursively and invokes a
// function on each route. If this function returns an error, the
// error will be propagated back to the caller.
//
// WalkRoutes will walk down the route tree to the desired depth.
// Set `maxDepth` to `-1` to disable this limitation.
func (s *Site) WalkRoutes(walkFn func(route *Route) error, maxDepth int) error {
	return s.walkRoute(&s.Root, walkFn, maxDepth, 0)
}

// walkRoute is the internal implementation of WalkRoutes which
// calls the walk function on a given route and then calls itself
// for each of the route's children.
func (s *Site) walkRoute(route *Route, walkFn func(route *Route) error, maxDepth, curDepth int) error {
	if maxDepth != -1 && curDepth == maxDepth {
		return nil
	}
	curDepth++

	if err := walkFn(route); err != nil {
		return err
	}

	for _, route := range route.Children {
		if err := s.walkRoute(route, walkFn, maxDepth, curDepth); err != nil {
			return err
		}
	}

	return nil
}

// ResolveRoute resolves a route in the route tree and returns
// it. Returns an error if the route cannot be found.
func (s *Site) ResolveRoute(route string) (*Route, error) {
	var (
		node     = &s.Root
		segments = strings.Split(route, "/")
	)

	for i, s := range segments {
		if i == len(segments)-1 {
			return node, nil
		}
		if _, exists := node.Children[s]; !exists {
			return nil, fmt.Errorf("child %s does not exist", s)
		}
		node = node.Children[s]
	}

	return nil, fmt.Errorf("route %s does not exist", route)
}

// ResolvePage resolves a FQN to a page in the route list.
func (s *Site) ResolvePage(fqn FQN) (Page, error) {
	route, err := s.
		ResolveRoute(fqn.Route())

	if err != nil {
		return Page{}, err
	}

	for _, p := range route.Pages {
		if p.ID == fqn.PageID() {
			return p, nil
		}
	}

	return Page{}, fmt.Errorf("related page %s not found", fqn.String())
}

// CreateRoute creates a route in the route tree along with its
// parents and returns that route.
func (s *Site) CreateRoute(route string) *Route {
	var (
		node     = &s.Root
		segments = strings.Split(route, "/")
	)

	for i, s := range segments {
		if _, exists := node.Children[s]; !exists {
			node.Children[s] = &Route{
				Children:  make(map[string]*Route),
				Pages:     make([]Page, 0),
				IndexPage: IndexPage{},
			}
		}
		if i == len(segments)-1 {
			return node
		}
		node = node.Children[s]
	}

	return nil
}
