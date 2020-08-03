// Package parser provides a public interface for parsing files
// and converting them to Espresso entities.
package parser

import "github.com/dominikbraun/espresso/entity"

// Parser prescribes methods for parsing and converting files.
type Parser interface {
	ParsePage(source []byte) (entity.Page, error) // ParsePage converts a byte slice into a Page.
}
