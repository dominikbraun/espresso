package core

import (
	"fmt"
	"github.com/dominikbraun/espresso/build"
	"github.com/dominikbraun/espresso/filesystem"
	"github.com/dominikbraun/espresso/parser"
	"github.com/dominikbraun/espresso/settings"
	"io/ioutil"
	"path/filepath"
	"sync"
)

const (
	contentDir string = "content"
	numWorkers int    = 5
)

type Espresso struct {
	buildPath string
	builder   *build.Builder
	parser    parser.Parser
}

func NewEspresso(buildPath string, settings *settings.Site, parser parser.Parser) *Espresso {
	e := Espresso{
		buildPath: buildPath,
		builder:   build.NewBuilder(settings),
		parser:    parser,
	}

	return &e
}

func (e *Espresso) RunBuild() error {
	files := make(chan string)
	results := make(chan *build.Subject)

	contentPath := fmt.Sprintf("%s/%s", e.buildPath, contentDir)

	var wg sync.WaitGroup

	go func() {
		_ = filesystem.Stream(contentPath, filesystem.MarkdownOnly, files)
	}()

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go e.receiveFiles(files, results, &wg)
	}

	go e.forwardToBuilder(results)

	wg.Wait()
	close(results)

	return nil
}

func (e *Espresso) receiveFiles(files <-chan string, results chan<- *build.Subject, wg *sync.WaitGroup) {
	for file := range files {
		subject, _ := e.processFile(file)
		results <- subject
	}
	wg.Done()
}

func (e *Espresso) forwardToBuilder(results <-chan *build.Subject) {
	for result := range results {
		e.builder.Add(result)
	}
}

func (e *Espresso) processFile(file string) (*build.Subject, error) {
	source, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	article, err := e.parser.ParseArticle(source)
	if err != nil {
		return nil, err
	}

	subject := build.Subject{
		Route:   filepath.Dir(file),
		Article: &article,
	}

	return &subject, nil
}
