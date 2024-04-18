package model

import (
	protoTheme "github.com/esklo/residents-tracking-platform-backend/gen/proto/theme"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Theme struct {
	Id           uuid.UUID
	Title        string
	Priority     int64
	DepartmentId uuid.UUID
	DeletedAt    *time.Time
}

func (t *Theme) ToProto() (*protoTheme.Theme, error) {
	if t == nil {
		return nil, ErrorModelIsEmpty
	}
	theme := protoTheme.Theme{
		Id:           t.Id.String(),
		Title:        t.Title,
		Priority:     t.Priority,
		DepartmentId: t.DepartmentId.String(),
	}
	if t.DeletedAt != nil {
		theme.DeletedAt = timestamppb.New(*t.DeletedAt)
	}
	return &theme, nil
}
