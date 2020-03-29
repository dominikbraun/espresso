// Package core provides the entry-point functions for all Espresso
// commands and ties together the individual Espresso components.
package core

import (
	"github.com/dominikbraun/espresso/build"
	"github.com/dominikbraun/espresso/filesystem"
	"github.com/dominikbraun/espresso/parser"
	"github.com/dominikbraun/espresso/settings"
)

const (
	contentDir string = "/content"
)

// RunBuild performs a website build based on content files and settings
// stored in the build path, rendering a complete static website.
func RunBuild(buildPath string, settings *settings.Site) error {
	files := make(chan string)
	contentPath := buildPath + contentDir

	go func() {
		_ = filesystem.Stream(contentPath, filesystem.MarkdownOnly, files)
	}()

	build.Run(&build.Context{
		BuildPath: buildPath,
		Settings:  settings,
		Parser:    parser.NewMarkdown(),
	}, files)

	return nil
}
