package parser

import "github.com/dominikbraun/espresso/entity"

type Parser interface {
	ParsePage(source []byte) (entity.Page, error)
}
