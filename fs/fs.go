package fs

import (
	"github.com/otiai10/copy"
	"os"
	"path/filepath"
	"strings"
)

var (
	MarkdownOnly = func(file string) bool {
		return filepath.Ext(file) == ".md"
	}
	IgnoreUnderscores = func(file string) bool {
		return !strings.HasPrefix(filepath.Base(file), "_")
	}
)

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

func CreateDir(path string, createParents bool) error {
	if createParents {
		return os.MkdirAll(path, 0700)
	}
	return os.Mkdir(path, 0700)
}

func CreateFile(file string) (*os.File, error) {
	return os.Create(file)
}

func CopyDir(source, dest string) error {
	return copy.Copy(source, dest)
}
