package fs

import (
	"bytes"
	"io/fs"
)

// MemFS is an in-memory fs.FS implementation backed by a map of files indexed
// by name (path).
type MemFS map[string][]byte

var _ fs.FS = (MemFS)(nil)

// NewMemFS returns a new initialized MemFS.
func NewMemFS() MemFS {
	return make(MemFS)
}

// Open implements fs.FS.
func (mfs MemFS) Open(name string) (fs.File, error) {
	f, exists := mfs[name]
	if !exists {
		return nil, &fs.PathError{Op: "read", Path: name, Err: fs.ErrNotExist}
	}

	return &memFD{b: *bytes.NewBuffer(f)}, nil
}

// CreateFile creates a new file in the MemFS.
func (mfs MemFS) CreateFile(name string, data []byte) error {
	if _, exists := mfs[name]; exists {
		return &fs.PathError{Op: "create", Path: name, Err: fs.ErrExist}
	}

	mfs[name] = data
	return nil
}

// memFD is an in-memory representation of a file descriptor backed by a
// bytes.Buffer.
type memFD struct {
	b bytes.Buffer
}

var _ fs.File = (*memFD)(nil)

// Read implements fs.File and satisfies io.Reader.
func (fd *memFD) Read(b []byte) (int, error) {
	return fd.b.Read(b)
}

// Close implements fs.File.
func (fd *memFD) Close() error {
	fd.b.Reset()
	return nil
}

// Stat implements fs.File.
func (*memFD) Stat() (fs.FileInfo, error) {
	panic("not implemented")
}
