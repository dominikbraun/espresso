// Package filesystem provides utility functions serving as unified
// interfaces for creating and accessing filesystem objects.
package filesystem

import "os"

// CreateDir creates a directory. createParents indicates if the
// parent directories specified in the path should be created if
// they don't exist yet.
func CreateDir(path string, createParents bool) error {
	if createParents {
		return os.MkdirAll(path, 0700)
	}
	return os.Mkdir(path, 0700)
}

// CreateFile creates a file and returns an OS handle to it.
func CreateFile(file string) (*os.File, error) {
	return os.Create(file)
}
