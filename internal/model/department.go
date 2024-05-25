package model

import (
	protoDepartment "github.com/esklo/residents-tracking-platform-backend/gen/proto/department"
	"github.com/google/uuid"
)

type Department struct {
	Id         uuid.UUID
	Title      string
	DistrictId uuid.UUID
	FullAccess bool
}

func (d *Department) ToProto() (*protoDepartment.Department, error) {
	if d == nil {
		return nil, ErrorModelIsEmpty
	}
	department := protoDepartment.Department{
		Id:         d.Id.String(),
		Title:      d.Title,
		DistrictId: d.DistrictId.String(),
		FullAccess: d.FullAccess,
	}
	return &department, nil
}
