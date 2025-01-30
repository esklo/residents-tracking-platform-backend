package model

import (
	protoFile "github.com/esklo/residents-tracking-platform-backend/gen/proto/file"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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

func (f *File) FromProto(file *protoFile.File) error {
	fileId, err := uuid.Parse(file.Id)
	if err != nil {
		return errors.Wrap(err, "can not parse file id")
	}
	f.Id = fileId
	f.FileName = file.Filename
	f.MimeType = file.Mimetype
	f.Extension = file.Extension
	f.PublicUrl = file.Url
	return nil
}
