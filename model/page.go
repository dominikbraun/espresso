package model

type page struct {
	Route  string
	Nav    *Nav
	Footer *Footer
}

type ArticlePage struct {
	page
	Article Article
}
