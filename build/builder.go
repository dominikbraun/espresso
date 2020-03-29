package build

import (
	"github.com/dominikbraun/espresso/model"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)

type builder struct {
	ctx   *Context
	model *Site
	mutex *sync.Mutex
}

func newBuilder(ctx *Context) *builder {
	b := builder{
		ctx:   ctx,
		model: newSite(),
		mutex: &sync.Mutex{},
	}
	return &b
}

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
