// Package build provides all functionality required for performing a
// build from reading content files to modelling the static site.
package build

import (
	"github.com/dominikbraun/espresso/model"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)

// builder is the type used for performing the actual build. It knows
// the current build context and generates the entire site model which
// can then be rendered to a static site.
type builder struct {
	ctx   Context
	model *Site
	mutex *sync.Mutex
}

// newBuilder creates a builder instance that utilizes the build context.
func newBuilder(ctx Context) *builder {
	b := builder{
		ctx:   ctx,
		model: newSite(),
		mutex: &sync.Mutex{},
	}
	return &b
}

// buildPage attempts to generate a model.Page from a file. This is done
// by reading the file, parsing its content and building a page model. The
// page is automatically added to the builder's site model.
//
// buildPage is safe for concurrent invocation. The file path must contain
// the build path.
func (b *builder) buildPage(file string) error {
	source, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	article, err := b.ctx.Parser.ParseArticle(source)
	if err != nil {
		return err
	}

	route := filepath.Dir(file)
	id := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))

	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.model.registerPage(&model.ArticlePage{
		Page: model.Page{
			Path: route,
			ID:   id,
		},
		Article: article,
	})

	return nil
}

// buildNav attempts to create a model.Nav from the existing pages that
// have to be built and registered first, meaning that buildNav must be
// called after all buildPage calls have finished.
//
// buildNav takes the site settings into account and overrides the Nav if
// this is specified in the site settings.
func (b *builder) buildNav() error {
	nav := &model.Nav{
		Brand: b.ctx.Settings.Name,
		Items: make([]model.NavItem, 0),
	}

	for _, i := range b.ctx.Settings.Nav.Items {
		item := model.NavItem{
			Label:  i.Label,
			Target: i.Target,
		}
		nav.Items = append(nav.Items, item)
	}

	if !b.ctx.Settings.Nav.Override {
		b.model.walkRoutes(func(r *route) {
			for seg, _ := range r.children {
				item := model.NavItem{
					Label:  strings.Title(seg),
					Target: seg,
				}
				nav.Items = append(nav.Items, item)
			}
		}, 1)
	}

	b.model.Nav = nav
	return nil
}

// buildFooter attempts to create a model.Footer under consideration of
// user-defined site settings. It is independent from any site pages.
func (b *builder) buildFooter() error {
	footer := &model.Footer{
		Text:  b.ctx.Settings.Footer.Text,
		Items: make([]model.FooterItem, 0),
	}

	for _, i := range b.ctx.Settings.Footer.Items {
		item := model.FooterItem{
			Label:  i.Label,
			Target: i.Target,
		}
		footer.Items = append(footer.Items, item)
	}

	b.model.Footer = footer
	return nil
}
