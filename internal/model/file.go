package model

import (
	protoFile "github.com/esklo/residents-tracking-platform-backend/gen/proto/file"
	"github.com/google/uuid"
)

type File struct {
	Id        uuid.UUID
	Path      string
	FileName  string
	MimeType  string
	Extension string
	PublicUrl string
}

func (f *File) ToProto() (*protoFile.File, error) {
	if f == nil {
		return nil, ErrorModelIsEmpty
	}
	fileProto := protoFile.File{
		Id:        f.Id.String(),
		Filename:  f.FileName,
		Mimetype:  f.MimeType,
		Extension: f.Extension,
		Url:       f.PublicUrl,
	}
	return &fileProto, nil
}
