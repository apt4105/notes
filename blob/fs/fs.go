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

func (fs *FS) fileName(name string) string {
	return filepath.Join(fs.Basepath, name)
}

func (fs *FS) Get(name string) (io.ReadCloser, error) {
	file, err := os.OpenFile(fs.fileName(name), os.O_RDONLY, 0644)

	if errAs := (&os.PathError{}); errors.As(err, &errAs) {
		return nil, fmt.Errorf("%w: %v", blob.ErrNotFound, err)
	}

	return file, err
}

func (fs *FS) Set(name string, data io.Reader) error {
	file, err := os.OpenFile(fs.fileName(name),
		os.O_WRONLY | os.O_TRUNC | os.O_CREATE ,
		0644)

	if err != nil {
		return err
	}

	_, err = io.Copy(file, data)

	return err
}

func (fs *FS) Del(name string) error {
	return os.Remove(fs.fileName(name))
}

func (fs *FS) Stat(name string) (blob.Info, error) {
	return blob.Info{}, nil
}
