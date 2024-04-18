package model

import (
	protoDepartment "github.com/esklo/residents-tracking-platform-backend/gen/proto/department"
	"github.com/google/uuid"
)

type Department struct {
	Id         uuid.UUID
	Title      string
	DistrictId uuid.UUID
}

func (d *Department) ToProto() (*protoDepartment.Department, error) {
	if d == nil {
		return nil, ErrorModelIsEmpty
	}
	department := protoDepartment.Department{
		Id:         d.Id.String(),
		Title:      d.Title,
		DistrictId: d.DistrictId.String(),
	}
	return &department, nil
}
