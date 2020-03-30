// Package model provides all domain models representing the individual
// site components.
package model

// NavItem represents an hyperlink in the site navigation where Target
// is the link's target URL.
type NavItem struct {
	Label  string
	Target string
}

// Nav represents the site navigation which will be included in all
// pages.
type Nav struct {
	Brand string
	Items []NavItem
}
