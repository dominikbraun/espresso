package parser

import (
	"bytes"
	"github.com/dominikbraun/espresso/model"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"time"
)

type Markdown struct {
	inner goldmark.Markdown
}

func NewMarkdown() *Markdown {
	m := Markdown{
		inner: goldmark.New(goldmark.WithExtensions(meta.Meta)),
	}
	return &m
}

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

	return nil
}
