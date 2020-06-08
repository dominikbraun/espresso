// Package model provides all domain models representing the individual
// site components.
package model

import "time"

// Content represents general-purpose content like an 'About Me' page
// or any other types of content that can't be categorized precisely.
type Content struct {
	ID           string
	Title        string
	RelatedLinks []string
	Related      []Article
	Description  string
	Content      string
}

// Article represents an editorial article like a magazine article or
// blog post. An article consists of metadata and its actual content.
type Article struct {
	ID           string
	Title        string
	Author       string
	Date         time.Time
	Tags         []string
	RelatedLinks []string
	Related      []*Article
	Description  string
	Content      string
}
