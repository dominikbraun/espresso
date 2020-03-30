// Package model provides all domain models representing the individual
// site components.
package model

import "time"

// Article represents an editorial article like a magazine article or
// blog post. An article consists of metadata and its actual content.
type Article struct {
	Title       string
	Author      string
	Date        time.Time
	Tags        []string
	Description string
	Content     string
}
