// Package fs contains implementations of the fs.FS interface.
package fs

import "io/fs"

// FS is a type alias to avoid having to import both fs.FS and this package.
type FS = fs.FS
