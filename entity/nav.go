package entity

type Nav struct {
	Items []NavItem
}

type NavItem struct {
	Label  string
	Target string
}
