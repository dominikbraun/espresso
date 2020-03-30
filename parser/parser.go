// Package parser provides implementations for parsing content files.
package parser

import "github.com/dominikbraun/espresso/model"

// Parser is a generic interface that prescribes methods that any
// kind of content parser needs to implement.
type Parser interface {
	ParseArticle(source []byte) (model.Article, error)
}
