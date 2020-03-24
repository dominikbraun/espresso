package parser

import "github.com/dominikbraun/espresso/model"

type Parser interface {
	ParseArticle(source []byte) (model.Article, error)
}
