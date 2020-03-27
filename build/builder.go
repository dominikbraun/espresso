package build

import (
	"fmt"
	"github.com/dominikbraun/espresso/model"
	"github.com/dominikbraun/espresso/parser"
	"github.com/dominikbraun/espresso/settings"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)

const (
	numWorkers    int    = 5
	pathDelimiter string = "\"
)

type Route struct {
	Pages    []*model.ArticlePage
	Children map[string]Route
}

type Site struct {
	Nav    *model.Nav
	Root   Route
	Footer *model.Footer
}

type Builder struct {
	model    *Site
	settings *settings.Site
	parser   parser.Parser
	mutex    *sync.Mutex
}

func NewBuilder(settings *settings.Site, parser parser.Parser) *Builder {
	b := Builder{
		model: &Site{
			Nav: &model.Nav{},
		},
		settings: settings,
		parser:   parser,
		mutex:    &sync.Mutex{},
	}
	return &b
}

func (b *Builder) Receive(files <-chan string) {
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go b.processFiles(files, &wg)
	}

	wg.Wait()
}

func (b *Builder) processFiles(files <-chan string, wg *sync.WaitGroup) {
	for file := range files {
		page, _ := b.buildPage(file)
		b.registerPage(page)
	}
	wg.Done()
}

func (b *Builder) buildPage(file string) (*model.ArticlePage, error) {
	source, err := ioutil.ReadFile(file)
	if err != nil {
		return &model.ArticlePage{}, nil
	}

	article, err := b.parser.ParseArticle(source)
	if err != nil {
		return &model.ArticlePage{}, nil
	}

	route := filepath.Dir(file)
	page := model.NewArticlePage(route, article)

	return page, nil
}

func (b *Builder) registerPage(page *model.ArticlePage) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	node := &b.model.Root
	directories := strings.Split(page.Route, pathDelimiter)

	for i, dir := range directories {
		if _, exists := node.Children[dir]; !exists {
			node.Children[dir] = Route{
				Pages:    make([]*model.ArticlePage, 0),
				Children: make(map[string]Route),
			}
		}
		if i == len(directories) - 1 {
			node.Children[dir].Pages = append(node.Children[dir].Pages, page)
		}
	}

	b.model.Routes[page.Route] = append(b.model.Routes[page.Route], page)
}

func (b *Builder) GenerateModel() {
	b.buildNav()
}

func (b *Builder) buildNav() {
	b.model.Nav.Brand = b.settings.Name

	if b.settings.Nav.Override {
		for _, item := range b.settings.Nav.Items {
			navItem := model.NavItem{
				Label:  item.Label,
				Target: item.Target,
			}
			b.model.Nav.Items = append(b.model.Nav.Items, navItem)
		}
		return
	}
	fmt.Println("Won't build nav automatically yet.")
}
