// Package model provides all domain models representing the individual
// site components.
package model

// FooterItem represents an hyperlink in the site footer where Target
// is the link's target URL.
type FooterItem struct {
	Label  string
	Target string
}

// Footer represents the site footer which will be included in all pages.
type Footer struct {
	Text  string
	Items []FooterItem
}
