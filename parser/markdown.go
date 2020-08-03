// Package parser provides a public interface for parsing files
// and converting them to Espresso entities.
package parser

import (
	"bytes"
	"github.com/dominikbraun/espresso/entity"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// Markdown represents a Markdown parser.
type Markdown struct {
	gm goldmark.Markdown
}

// NewMarkdown creates a fully initialized Markdown parser.
func NewMarkdown() Markdown {
	m := Markdown{
		gm: goldmark.New(
			goldmark.WithExtensions(meta.Meta, highlighting.Highlighting),
		),
	}
	return m
}

// ParsePage implements Parser.ParsePage. It converts a given
// byte slice into a Page instance with populated content and
// metadata. Note that it leaves out the ID, thus ID is empty.
func (m Markdown) ParsePage(source []byte) (entity.Page, error) {
	var (
		page entity.Page
		buf  bytes.Buffer
		ctx  = parser.NewContext()
	)

	if err := m.gm.Convert(source, &buf, parser.WithContext(ctx)); err != nil {
		return page, err
	}

	page.Content = buf.String()
	metadata := meta.Get(ctx)

	readMetadata(metadata, &page)

	return page, nil
}
