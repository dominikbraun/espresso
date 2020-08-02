// Package entity provides the Espresso domain models. They are
// used to represent the entire website and its components.
package entity

// Nav represents the global website navigation that is identical
// on every page. Espresso auto-generates the navigation, but nav
// items can be added manually by the user in espresso.yml. It is
// also possible to completely override the navigation from there.
type Nav struct {
	Items []NavItem // Items holds all navigation items.
}

// NavItem represents a single item in the global navigation.
type NavItem struct {
	Label  string // Label is the item label, like `Home`.
	Target string // Target is a URL in the form `https://example.com`.
}
