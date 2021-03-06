// Package model provides all domain models representing the individual
// site components.
package model

import "strings"

// Page represents a particular page of the website. This type provides
// data required for every kind of page and is useful for embedding.
type Page struct {
	// Path may only contain "/" characters as directory separators. This
	// constraint has to be taken into account on each instantiation.
	Path   string
	Nav    *Nav
	Footer *Footer
}

// RelativePath returns the page's relative path. By default, Espresso
// stores all paths inside the content directory as absolute paths.
func (p *Page) RelativePath() string {
	// According to Espresso's path model, this actually shouldn't happen.
	if !strings.HasPrefix(p.Path, "/") {
		return p.Path
	}
	return p.Path[1:]
}

// ArticlePage is a website page type that holds an article.
type ArticlePage struct {
	Page
	Article Article
}

// ListPage is a website page type that holds a list of articles.
type ListPage struct {
	Page
	ArticlePages []*ArticlePage
}

// IndexPage is an index page provided by the user. If a content directory
// contains an `index.md` file, this file will be parsed and stored as the
// IndexPage for a route.
//
// This IndexPage will be rendered instead of the  route's auto-generated
// ListPage.
type IndexPage struct {
	Page
	Article      Article
	ArticlePages []*ArticlePage
}
