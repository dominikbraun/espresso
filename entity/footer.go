// Package entity provides the Espresso domain models. They are
// used to represent the entire website and its components.
package entity

// Footer represents the global website footer that is identical
// on every page. Espresso auto-generates the footer, but footer
// items can be added manually by the user in espresso.yml. It is
// also possible to completely override the footer from there.
type Footer struct {
	Items []FooterItem // Items holds all footer items.
}

// FooterItem represents a single item in the global footer.
type FooterItem struct {
	Label  string // Label is the item label, like `Home`.
	Target string // Target is a URL in the form `https://example.com`.
}
