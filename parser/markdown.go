// Package parser provides implementations for parsing content files.
package parser

import (
	"bytes"
	"github.com/dominikbraun/espresso/model"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"time"
)

// Markdown is a Parser implementation for parsing Markdown files and
// converting them to domain models.
type Markdown struct {
	inner goldmark.Markdown
}

// NewMarkdown creates and initializes a new Markdown parser.
func NewMarkdown() *Markdown {
	m := Markdown{
		inner: goldmark.New(
			goldmark.WithExtensions(meta.Meta, highlighting.Highlighting),
		),
	}
	return &m
}

// ParseArticle implements Parser.ParseArticle. It takes the contents
// of a Markdown file as a []byte, parses the contents as HTML and
// populates the article metadata with metadata form the Markdown file.
func (m *Markdown) ParseArticle(source []byte) (model.Article, error) {
	var buf bytes.Buffer
	ctx := parser.NewContext()

	if err := m.inner.Convert(source, &buf, parser.WithContext(ctx)); err != nil {
		return model.Article{}, err
	}

	article := model.Article{}
	article.Content = buf.String()

	if err := m.readMetadata(&article, meta.Get(ctx)); err != nil {
		return model.Article{}, err
	}

	return article, nil
}

// readMetadata populates an article's metadata fields with values from
// a metadata map as it is returned by the meta.Get function. Note that
// all metadata keys are case-sensitive.
func (m *Markdown) readMetadata(article *model.Article, metadata map[string]interface{}) error {
	if metadata["Title"] != nil {
		article.Title = metadata["Title"].(string)
	}

	if metadata["Author"] != nil {
		article.Author = metadata["Author"].(string)
	}

	if metadata["Description"] != nil {
		article.Description = metadata["Description"].(string)
	}

	if metadata["Date"] != nil {
		date, err := time.Parse("2006-01-02", metadata["Date"].(string))
		if err != nil {
			return err
		}
		article.Date = date
	}

	if metadata["Tags"] != nil {
		tags, _ := metadata["Tags"].([]interface{})
		for i := 0; i < len(tags); i++ {
			article.Tags = append(article.Tags, tags[i].(string))
		}
	}

	if metadata["Img"] != nil {
		article.Img = metadata["Img"].(string)
	}

	if metadata["ImgCredit"] != nil {
		article.ImgCredit = metadata["ImgCredit"].(string)
	}

	if metadata["Related"] != nil {
		related, _ := metadata["Related"].([]interface{})
		for i := 0; i < len(related); i++ {
			article.Related = append(article.Related, related[i].(string))
		}
	}

	if metadata["Template"] != nil {
		article.Template = metadata["Template"].(string)
	}

	return nil
}
