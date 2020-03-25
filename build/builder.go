package build

import (
	"fmt"
	"github.com/dominikbraun/espresso/model"
	"github.com/dominikbraun/espresso/parser"
	"github.com/dominikbraun/espresso/settings"
	"io/ioutil"
	"sync"
)

const (
	numWorkers int = 5
)

type Site struct {
	Nav    *model.Nav
	Pages  map[string][]*model.ArticlePage
	Footer *model.Footer
}

type Builder struct {
	model    *Site
	settings *settings.Site
	parser   parser.Parser
}

func NewBuilder(settings *settings.Site, parser parser.Parser) *Builder {
	b := Builder{
		model:    &Site{},
		settings: settings,
		parser:   parser,
	}
	return &b
}

func (b *Builder) Receive(files <-chan string) {
	var wg sync.WaitGroup
	results := make(chan *model.ArticlePage)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go b.processFiles(files, results, &wg)
	}

	go b.processResults(results)

	wg.Wait()
	close(results)
}

func (b *Builder) processFiles(files <-chan string, results chan<- *model.ArticlePage, wg *sync.WaitGroup) {
	for file := range files {
		page, _ := b.buildPage(file)
		results <- page
	}
	wg.Done()
}

func (b *Builder) processResults(results <-chan *model.ArticlePage) {
	for page := range results {
		fmt.Println("Adding \"" + page.Article.Title + "\" to model")
	}
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

	page := model.ArticlePage{
		Article: article,
	}

	return &page, nil
}
