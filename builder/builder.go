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
	if _, exists := b.site.Routes[route]; !exists {
		r := entity.Route{
			Pages:     make([]entity.Page, 0),
			IndexPage: entity.IndexPage{},
		}
		b.site.Routes[route] = &r
	}

	b.site.Routes[route].Pages = append(b.site.Routes[route].Pages, page)
	b.site.Routes[route].IndexPage.Pages = append(b.site.Routes[route].IndexPage.Pages, page)

	return nil
}

func (b *builder) Dispatch() (entity.Site, error) {
	var (
		site entity.Site
	)

	if err := b.resolveRelated(); err != nil {
		return site, err
	}

	return site, nil
}

func (b *builder) resolveRelated() error {
	err := b.site.
		WalkRoutes(func(route *entity.Route) error {
			for _, page := range route.Pages {
				for _, fqn := range page.RelatedFQNs {

					p, err := b.site.ResolvePage(fqn)
					if err != nil {
						return err
					}
					page.Related = append(page.Related, p)
				}
			}
			return nil
		})

	return err
}
