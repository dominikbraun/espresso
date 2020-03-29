package core

import (
	"github.com/dominikbraun/espresso/filesystem"
	"github.com/dominikbraun/espresso/parser"
	"github.com/dominikbraun/espresso/settings"
)

const (
	contentDir string = "/content"
)

type Espresso struct {
	buildPath   string
	contentPath string
}

func NewEspresso(buildPath string, settings *settings.Site, parser parser.Parser) *Espresso {
	e := Espresso{
		buildPath:   buildPath,
		contentPath: buildPath + contentDir,
	}

	return &e
}

func (e *Espresso) RunBuild() error {
	files := make(chan string)

	go func() {
		_ = filesystem.Stream(e.contentPath, filesystem.MarkdownOnly, files)
	}()

	return nil
}
