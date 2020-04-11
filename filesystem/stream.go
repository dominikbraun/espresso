// Package filesystem provides utility functions serving as unified
// interfaces for creating and accessing filesystem objects.
package filesystem

import (
	"os"
	"path/filepath"
)

// MarkdownOnly is a predefined filter that only allows .md files
// to be processed by the Stream function.
var MarkdownOnly = func(path string) bool {
	return filepath.Ext(path) == ".md"
}

// Stream streams all files in a given path that pass a given filter
// by sending them through the files channel. Stream.
func Stream(path string, filter func(path string) bool, files chan<- string) error {
	err := filepath.
		Walk(path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() || !filter(path) {
				return nil
			}

			files <- path

			return nil
		})

	close(files)
	return err
}
