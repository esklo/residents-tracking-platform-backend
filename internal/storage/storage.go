package storage

import (
	"github.com/gabriel-vasile/mimetype"
	"io"
)

type Storage interface {
	PutFile(r io.ReadSeeker) (path string, mime *mimetype.MIME, err error)
	GetFile(path string) (r io.ReadCloser, err error)
	DeleteFile(path string) (err error)
	GetFileUrl(path string) (url string, err error)
}
