// Package fs provides a variety of filesystem utilities. It
// is Espresso's interface for interacting with the filesystem.
package fs

import (
	"github.com/otiai10/copy"
	"os"
	"path/filepath"
	"strings"
)

var (
	// MarkdownOnly is a predefined filter that only allows
	// .md files.
	MarkdownOnly = func(file string) bool {
		return filepath.Ext(file) == ".md"
	}
	// IgnoreUnderscores is a filter that doesn't let pass
	// files starting with an underscore.
	IgnoreUnderscores = func(file string) bool {
		return !strings.HasPrefix(filepath.Base(file), "_")
	}
)

// StreamFiles streams all files in `path` that match all `filters`
// by sending them through the `files` channel.
func StreamFiles(path string, files chan<- string, filters ...func(file string) bool) error {
	err := filepath.
		Walk(path, func(file string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			for _, filter := range filters {
				if !filter(file) {
					return nil
				}
			}
			files <- file
			return nil
		})

	close(files)
	return err
}

// CreateDir creates a directory specified in `path`.
func CreateDir(path string, createParents bool) error {
	if createParents {
		return os.MkdirAll(path, 0700)
	}
	return os.Mkdir(path, 0700)
}

// CreateFile creates a file with the given filename.
func CreateFile(file string) (*os.File, error) {
	return os.Create(file)
}

// CopyDir copies a directory to `dest` recursively.
func CopyDir(source, dest string) error {
	return copy.Copy(source, dest)
}
