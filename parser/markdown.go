package parser

import (
	"fmt"
	"github.com/dominikbraun/espresso/model"
	"github.com/yuin/goldmark"
)

type Markdown struct {
	inner *goldmark.Markdown
}

func NewMarkdown() *Markdown {
	return nil
}

func (m *Markdown) ParseArticle(source []byte) (model.Article, error) {
	fmt.Println("Parsing article")
	return model.Article{}, nil
}
