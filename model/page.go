package model

type Page struct {
	Route string
}

type ArticlePage struct {
	Page
	Article Article
}

func NewArticlePage(route string, article Article) *ArticlePage {
	a := ArticlePage{
		Page:    Page{Route: route},
		Article: article,
	}

	return &a
}
