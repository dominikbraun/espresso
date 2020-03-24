package build

import (
	"fmt"
	"github.com/dominikbraun/espresso/model"
	"github.com/dominikbraun/espresso/settings"
)

type Site struct {
	Nav    *model.Nav
	Pages  map[string][]*model.ArticlePage
	Footer *model.Footer
}

type Builder struct {
	model    *Site
	settings *settings.Site
}

type Subject struct {
	Route   string
	Article *model.Article
}

func NewBuilder(settings *settings.Site) *Builder {
	b := Builder{
		model:    &Site{},
		settings: settings,
	}
	return &b
}

func (b *Builder) Add(subject *Subject) {
	fmt.Println("Adding build subject")
}
