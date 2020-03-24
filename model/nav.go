package model

type NavItem struct {
	Label  string
	Target string
}

type Nav struct {
	Brand string
	Items []NavItem
}
