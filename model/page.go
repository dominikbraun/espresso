package model

type Page struct {
	Path string
	ID   string
}

type ArticlePage struct {
	Page
	Article Article
}
