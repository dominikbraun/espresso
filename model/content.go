// Package model provides all domain models representing the individual
// site components.
package model

import "time"

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
