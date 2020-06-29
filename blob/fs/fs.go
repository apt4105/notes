// package fs contains a BlobStore implementation backed
// by files on local disk (using Go's os package).
package fs

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/apt4105/notes/blob"
)

// FS implements BlobStore
type FS struct {
	Basepath string
}

func (fs *FS) Get(name string) (io.ReadCloser, error) {
	name = filepath.Join(fs.Basepath, name)
	file, err := os.Open(name)
	if errAs := (&os.PathError{}); errors.As(err, &errAs) {
		return nil, fmt.Errorf("%w: %v", blob.ErrNotFound, err)
	}
	return file, err
}

func (fs *FS) Set(name string, data io.Reader) error {
	return nil
}

func (fs *FS) Del(name string) error {
	return nil
}

func (fs *FS) Stat(name string) (blob.Info, error) {
	return blob.Info{}, nil
}
