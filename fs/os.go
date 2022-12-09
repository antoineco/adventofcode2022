package fs

import (
	"io/fs"
	"os"
)

// OSFS is a fs.FS implementation that uses functions provided by the os package.
type OSFS struct{}

var _ fs.FS = (*OSFS)(nil)

func (*OSFS) Open(name string) (fs.File, error) {
	return os.Open(name)
}
