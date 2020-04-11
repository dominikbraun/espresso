// Package filesystem provides utility functions serving as unified
// interfaces for creating and accessing filesystem objects.
package filesystem

import cp "github.com/otiai10/copy"

// CopyDir copies a directory recursively to the specified destination
// under consideration of a copy-policy suitable for Espresso.
func CopyDir(source, dest string) error {
	// So far, the recursive copy functionality is not implemented yet
	// so otiai10/copy is used for this purpose.
	return cp.Copy(source, dest)
}
