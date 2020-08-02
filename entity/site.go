// Package entity provides the Espresso domain models. They are
// used to represent the entire website and its components.
package entity

// Site represents the actual website and is thus the root entity
// which contains all website components.
//
// Each Espresso build aims to produce a completely populated Site
// entity which can be modified by plugins and used for rendering
// the website files.
type Site struct {
	Meta   Meta             // Meta contains page-agnostic, global metadata for the website itself.
	Nav    Nav              // Nav contains the global website navigation.
	Routes map[string]Route // Routes contains all routes and their pages.
	Footer Footer           // Footer contains the global website footer.
}
