package entity

type Route struct {
	Routes []Route
	Pages  []Page
	Index  Page
}

type Page struct {
	Path FQN
}

type FQN string
