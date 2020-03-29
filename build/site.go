package build

import (
	"github.com/dominikbraun/espresso/model"
	"os"
	"strings"
)

type Site struct {
	nav    *model.Nav
	root   route
	footer *model.Footer
}

type route struct {
	pages    []*model.ArticlePage
	children map[string]*route
}

func newSite() *Site {
	s := Site{
		nav: &model.Nav{
			Items: make([]model.NavItem, 0),
		},
		root: route{
			pages:    make([]*model.ArticlePage, 0),
			children: make(map[string]*route),
		},
		footer: &model.Footer{
			Items: make([]model.FooterItem, 0),
		},
	}
	return &s
}

func (s *Site) registerPage(page *model.ArticlePage) {
	node := &s.root
	segments := strings.Split(page.Path, string(os.PathSeparator))

	for i, seg := range segments {
		if _, exists := node.children[seg]; !exists {
			node.children[seg] = &route{
				pages:    make([]*model.ArticlePage, 0),
				children: make(map[string]*route),
			}
		}
		if i == len(segments)-1 {
			node.children[seg].pages = append(node.children[seg].pages, page)
			break
		}
		node = node.children[seg]
	}
}
