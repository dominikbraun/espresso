package build

import (
	"github.com/dominikbraun/espresso/model"
	"github.com/dominikbraun/espresso/parser"
	"github.com/dominikbraun/espresso/settings"
	"io/ioutil"
	"path/filepath"
	"sync"
)

const (
	numWorkers int = 5
)

type Site struct {
	Nav    *model.Nav
	Routes map[string][]*model.ArticlePage
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
			Routes: make(map[string][]*model.ArticlePage),
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

	b.model.Routes[page.Route] = append(b.model.Routes[page.Route], page)
}
