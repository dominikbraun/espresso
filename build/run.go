package build

import (
	"github.com/dominikbraun/espresso/parser"
	"github.com/dominikbraun/espresso/settings"
	"sync"
)

const (
	numWorkers int = 5
)

type Context struct {
	BuildPath string
	Settings  *settings.Site
	Parser    parser.Parser
}

func Run(ctx *Context, files <-chan string) {
	builder := newBuilder(ctx)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go processQueue(builder, files, &wg)
	}

	wg.Wait()
}

func processQueue(builder *builder, files <-chan string, wg *sync.WaitGroup) {
	for file := range files {
		_ = builder.buildPage(file)
	}
	wg.Done()
}
