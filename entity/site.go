// Package entity provides the Espresso domain models. They are
// used to represent the entire website and its components.
package entity

import "fmt"

// Site represents the actual website and is thus the root entity
// which contains all website components.
//
// Each Espresso build aims to produce a completely populated Site
// entity which can be modified by plugins and used for rendering
// the website files.
type Site struct {
	Meta   Meta              // Meta contains page-agnostic, global metadata for the website itself.
	Nav    Nav               // Nav contains the global website navigation.
	Routes map[string]*Route // Routes contains all routes and their pages.
	Footer Footer            // Footer contains the global website footer.
}

// WalkRoutes invokes a function on each registered route.
func (s Site) WalkRoutes(walkFn func(route *Route) error) error {
	for _, route := range s.Routes {
		if err := walkFn(route); err != nil {
			return err
		}
	}
	return nil
}

// ResolvePage resolves a FQN to a page in the route list.
func (s *Site) ResolvePage(fqn FQN) (Page, error) {
	var (
		page   Page
		route  = fqn.Route()
		pageID = fqn.PageID()
	)

	if _, exists := s.Routes[route]; !exists {
		return page, fmt.Errorf("related page %s: route %s does not exist", fqn.String(), route)
	}

	for _, p := range s.Routes[route].Pages {
		if p.ID == pageID {
			return page, nil
		}
	}

	return page, fmt.Errorf("related page %s not found", fqn.String())
}
