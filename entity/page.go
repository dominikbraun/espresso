// Package entity provides the Espresso domain models. They are
// used to represent the entire website and its components.
package entity

import "time"

// Page represents a particular website page that is available
// under a specific URL. Espresso will treat all Markdown files
// as individual pages associated with a Route.
type Page struct {
	ID          string    // ID is a page ID that is unique in its route.
	Title       string    // Title is the page title.
	Author      string    // Author is an author who can be distinct from Meta.Author.
	Date        time.Time // Date is the date the page has been created.
	Tags        []string  // Tags is a set of topics that the page covers.
	Img         Image     // Img typically is the image used as cover and for social media.
	Description string    // Description is a short, descriptive introduction text.
	Content     string    // Content is the actual page content.
	Related     []Page    // Related is a set of related Page instances.
	relatedFQNs []FQN     // relatedFQNs is a set of FQNs defined in the `Related` Markdown section.
	template    string    // template is an overriding template to be used for the page.
	hide        bool      // hide indicates if the page should be visible in lists and search engines.
}

// FQN is a Fully Qualified Name for an Espresso page similar
// to an absolute path consisting of a page route and a page ID.
// Thus, a valid FQN has the form `/blog/coffee/my-espresso`.
//
// FQNs are the common way to uniquely identify an Espresso page.
type FQN string

// Image represents a cover or social media image.
type Image struct {
	Src    string // Src is an image URL. A template may allow relative paths for Src.
	Credit string // Credit is the image credit or copyright owner.
}
