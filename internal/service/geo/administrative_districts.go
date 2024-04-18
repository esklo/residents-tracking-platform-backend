package geo

import (
	"context"
	"fmt"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/serjvanilla/go-overpass"
	"log"
)

func (s *Service) GetDistricts(_ context.Context, areaId int64, level int) (districts []*model.GeoDistrict, err error) {
	var areaOffset int64 = 3600000000
	client := overpass.New()
	query := fmt.Sprintf(`
[out:json][timeout:25];
(
  area(%d);
  relation(area)[admin_level=%d];
);
out geom;
`, areaOffset+areaId, level)

	result, _ := client.Query(query)
	for _, relation := range result.Relations {
		var district model.GeoDistrict
		district.Id = relation.ID
		district.Name = relation.Tags["name"]
		log.Printf("district: %#v", district)
		if district.Name == "" {
			continue
		}
		districts = append(districts, &district)
	}
	return
}
