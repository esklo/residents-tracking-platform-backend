package s3

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gabriel-vasile/mimetype"
	"io"
)

func getMimeInfo(r io.ReadSeeker) (*mimetype.MIME, error) {
	_, err := r.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	mime, err := mimetype.DetectReader(r)
	if err != nil {
		return nil, err
	}
	return mime, nil
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
