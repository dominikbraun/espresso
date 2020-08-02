// Package entity provides the Espresso domain models. They are
// used to represent the entire website and its components.
package entity

// Meta represents global metadata that applies to the website
// in general instead of a particular page.
//
// For example, the title shouldn't be 'How to make Espresso -
// MyCoffeeBlog' but only 'MyCoffeeBlog' instead. This generic
// metadata should be combined in with page-specific metadata
// in templates to obtain page titles like in the example.
//
// All this metadata can be defined by the user in espresso.yml.
type Meta struct {
	Title    string // Title is the global website title.
	Subtitle string // Subtitle is the global website subtitle.
	Author   string // Author is the website's author.
	Base     string // Base is an URL in the form `https://example.com`.
}
