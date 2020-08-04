package builder

import (
	"github.com/dominikbraun/espresso/entity"
)

type Builder interface {
	RegisterPage(route string, page entity.Page) error
	Dispatch() (entity.Site, error)
}

func New() Builder {
	b := builder{
		site: entity.Site{},
	}
	return &b
}

type builder struct {
	site entity.Site
}

func (b *builder) RegisterPage(route string, page entity.Page) error {
	r := b.site.CreateRoute(route)

	r.Pages = append(r.Pages, page)
	r.IndexPage.Pages = append(r.IndexPage.Pages, page)

	return nil
}

func (b *builder) Dispatch() (entity.Site, error) {
	return b.site, nil
}
