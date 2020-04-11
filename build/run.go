// Package build provides all functionality required for performing a
// build from reading content files to modelling the static site.
package build

import (
	"github.com/dominikbraun/espresso/config"
	"github.com/dominikbraun/espresso/parser"
	"io/ioutil"
	"sync"
)

const (
	numWorkers int = 8
)

// Context represents the build context. This context provides information
// for a particular build and may be entirely different for another one.
type Context struct {
	BuildPath string
	Settings  *config.Site
	Parser    parser.Parser
}

// Run starts new build with a given build context. It accepts a read-only
// channel for all content files that have to be included in the build and
// returns a build.Site model that can be used for rendering the website.
func Run(ctx Context, files <-chan string) *Site {
	builder := newBuilder(ctx)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go processQueue(builder, files, &wg)
	}

	wg.Wait()

	_ = builder.buildNav()
	_ = builder.buildListPages()
	_ = builder.buildRelated()
	_ = builder.buildFooter()

	return builder.model
}

// processQueue is a worker function that reads from the file channel and
// forwards these files to the builder for further processing. Reduces the
// WaitGroup counter by one after all received files have been built.
func processQueue(builder *builder, files <-chan string, wg *sync.WaitGroup) {
	for file := range files {
		source, _ := ioutil.ReadFile(file)
		_, _ = builder.buildPage(source, file, DirectRegister)
	}
	wg.Done()
}
