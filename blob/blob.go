package blob

import (
	"errors"
	"io"
	"time"
)

var ErrNotFound = errors.New("blob: not found")

type Store interface {
	Get(name string) (io.ReadCloser, error)
	Set(name string, data io.Reader) error
	Del(name string) error
	Stat(name string) (Info, string)
}

type Info struct {
	Name    string
	Size    int
	Created time.Time
	Updated time.Time
}
