// Package plugin provides build plugins for Espresso.
package plugin

import "github.com/dominikbraun/espresso/entity"

// Plugin represents a component that receives an Espresso
// page and processes it for its own purposes.
type Plugin interface {
	ProcessPage(page *entity.Page) error // ProcessPage processes a fully initialized page.
	Finalize() error                     // Finalize is called after all pages have been built.
}
