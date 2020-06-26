// Package filesystem provides utility functions serving as unified
// interfaces for creating and accessing filesystem objects.
package filesystem

import (
	"os"
	"path/filepath"
	"strings"
)

// MarkdownOnly is a predefined filter that only allows .md files
// to be processed by the Stream function.
var MarkdownOnly = func(path string) bool {
	return filepath.Ext(path) == ".md"
}

// NoUnderscores is a predefined filter that doesn't let pass
// files starting with an underscore.
var NoUnderscores = func(path string) bool {
	return !strings.HasPrefix(filepath.Base(path), "_")
}

// Stream streams all files in a given path that pass a given filter
// by sending them through the files channel. Stream.
func Stream(path string, files chan<- string, filters ...func(path string) bool) error {
	err := filepath.
		Walk(path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			for _, filter := range filters {
				if !filter(path) {
					return nil
				}
			}

			files <- path

			return nil
		})

	close(files)
	return err
}
