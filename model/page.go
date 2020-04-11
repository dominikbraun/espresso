// Package model provides all domain models representing the individual
// site components.
package model

// Page represents a particular page of the website. This type provides
// data required for every kind of page and is useful for embedding.
type Page struct {
	// Path may only contain "/" characters as directory separators. This
	// constraint has to be taken into account on each instantiation.
	Path string
}

// ArticlePage is a website page type that holds an article.
type ArticlePage struct {
	Page
	Article Article
}

// ArticleListPage is a website page type that holds a list of articles.
type ArticleListPage struct {
	Page
	Articles []*Article
}
