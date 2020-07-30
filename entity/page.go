package entity

import "time"

type Page struct {
	Path        FQN
	Title       string
	Author      string
	Date        time.Time
	Tags        []string
	Img         Image
	Description string
	Content     string
	Related     []Page
	relatedFQNs []FQN
	template    string
	hide        bool
}

type FQN string

type Image struct {
	Src    string
	Credit string
}
