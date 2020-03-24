package model

import "time"

type Article struct {
	Title       string
	Author      string
	Date        time.Time
	Tags        []string
	Description string
	Content     string
}
