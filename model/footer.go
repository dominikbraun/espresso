package model

type FooterItem struct {
	Label  string
	Target string
}

type Footer struct {
	Text  string
	Items []FooterItem
}
