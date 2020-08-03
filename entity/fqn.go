// Package entity provides the Espresso domain models. They are
// used to represent the entire website and its components.
package entity

import "path/filepath"

// FQN is a Fully Qualified Name for an Espresso page similar
// to an absolute path consisting of a page route and a page ID.
// Thus, a valid FQN has the form `/blog/coffee/my-espresso`.
//
// FQNs are the common way to uniquely identify an Espresso page.
type FQN string

// Route returns the FQN's route in the form `/blog/coffee`.
func (f FQN) Route() string {
	return filepath.Dir(string(f))
}

// PageID returns the FQN's page ID in the form `my-espresso`.
func (f FQN) PageID() string {
	return filepath.Base(string(f))
}

// String() returns the FQN as a string.
func (f FQN) String() string {
	return string(f)
}
