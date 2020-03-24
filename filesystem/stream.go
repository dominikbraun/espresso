package filesystem

import (
	"os"
	"path/filepath"
)

var MarkdownOnly = func(path string) bool {
	return filepath.Ext(path) == ".md"
}

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
