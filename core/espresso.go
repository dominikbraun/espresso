// Package core provides the entry-point functions for all Espresso
// commands and ties together the individual Espresso components.
package core

import (
	"github.com/dominikbraun/espresso/build"
	"github.com/dominikbraun/espresso/config"
	"github.com/dominikbraun/espresso/filesystem"
	"github.com/dominikbraun/espresso/parser"
	"github.com/dominikbraun/espresso/plugins/atom"
	"github.com/dominikbraun/espresso/render"
	"path/filepath"
)

type Options struct {
	OutputDir string
	RenderRSS bool
}

// RunBuild performs a website build based on content files and settings
// stored in the build path, rendering a complete static website.
func RunBuild(buildPath string, settings *config.Site, options *Options) error {
	files := make(chan string)
	contentPath := filepath.Join(buildPath, config.ContentDir)
	targetDir := filepath.Join(buildPath, config.TargetDir)

	if options.OutputDir != "" {
		targetDir = options.OutputDir
	}

	go func() {
		_ = filesystem.Stream(contentPath, files, filesystem.MarkdownOnly, filesystem.NoUnderscores)
	}()

	site := build.Run(build.Context{
		BuildPath: buildPath,
		Settings:  settings,
		Parser:    parser.NewMarkdown(),
	}, files)

	plugins := loadBuiltInPlugins(settings, options)

	if err := render.AsWebsite(render.Context{
		TemplateDir: filepath.Join(buildPath, config.TemplateDir),
		AssetDir:    filepath.Join(buildPath, config.AssetDir),
		TargetDir:   targetDir,
	}, site, plugins...); err != nil {
		return err
	}

	return nil
}

// loadBuiltInPlugins reads the build options and initializes all required
// built-in plugins using the site configuration.
func loadBuiltInPlugins(settings *config.Site, options *Options) []render.Plugin {
	plugins := make([]render.Plugin, 0)

	if options.RenderRSS {
		meta := atom.Meta{
			Title:       settings.Title,
			BaseURL:     settings.BaseURL,
			Description: settings.Description,
			Author:      settings.Author,
			Subtitle:    settings.Subtitle,
			Copyright:   settings.Author,
		}
		plugins = append(plugins, atom.New(&meta))
	}

	return plugins
}
