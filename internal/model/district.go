package model

import (
	protoDistrict "github.com/esklo/residents-tracking-platform-backend/gen/proto/district"
	"github.com/google/uuid"
	"github.com/paulmach/orb/geojson"
)

type District struct {
	Id               uuid.UUID
	Title            string
	GeoJson          []byte
	CoatOfArmsFileId *uuid.UUID
}

func (d *District) ToProto() (*protoDistrict.District, error) {
	if d == nil {
		return nil, ErrorModelIsEmpty
	}
	district := protoDistrict.District{
		Id:      d.Id.String(),
		Title:   d.Title,
		Geojson: d.GeoJson,
	}
	if d.CoatOfArmsFileId != nil {
		fileId := d.CoatOfArmsFileId.String()
		district.CoatOfArmsFileId = &fileId
	}
	return &district, nil
}

func (d *District) SetGeoJson(data *geojson.Feature) error {
	geojsonBytes, err := data.MarshalJSON()
	if err != nil {
		return err
	}
	d.GeoJson = geojsonBytes
	return nil
}
