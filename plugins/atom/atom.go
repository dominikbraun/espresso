package atom

import (
	"fmt"
	"github.com/dominikbraun/espresso/model"
	"github.com/gorilla/feeds"
	"time"
)

type Meta struct {
	Title       string
	Link        string
	Description string
	Author      string
	Subtitle    string
	Copyright   string
}

type atom struct {
	feed *feeds.Feed
}

func New(meta *Meta) *atom {
	a := atom{
		feed: &feeds.Feed{
			Title:       meta.Title,
			Link:        &feeds.Link{Href: meta.Link},
			Description: meta.Description,
			Author:      &feeds.Author{Name: meta.Author},
			Created:     time.Now(),
			Subtitle:    meta.Subtitle,
			Items:       make([]*feeds.Item, 0),
			Copyright:   meta.Copyright,
		},
	}
	return &a
}

func (a *atom) ProcessArticlePage(page *model.ArticlePage) error {
	item := &feeds.Item{
		Title:       page.Article.Title,
		Link:        &feeds.Link{Href: fmt.Sprintf("%s%s", page.Path, page.Article.ID)},
		Description: page.Article.Description,
		Created:     page.Article.Date,
	}
	a.feed.Items = append(a.feed.Items, item)

	return nil
}
