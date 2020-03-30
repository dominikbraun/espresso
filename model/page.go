// Package model provides all domain models representing the individual
// site components.
package model

// Page represents a particular page of the website. This type provides
// data required for every kind of page and is useful for embedding.
type Page struct {
	Path string
	ID   string
}

// ArticlePage is a website page type that holds an article.
type ArticlePage struct {
	Page
	Article Article
}
