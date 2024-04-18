package geo

import (
	"context"
	"fmt"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/serjvanilla/go-overpass"
)

func (s *Service) BuildingsByCoordinates(_ context.Context, lat, lon float64) (buildings []*model.GeoBuilding, err error) {
	client := overpass.New()
	query := fmt.Sprintf(`
[out:json][timeout:25];
(
  relation(around:50,%f,%f)[building];
  way(around:50,%f,%f)[building];
);
out geom;
`, lat, lon, lat, lon)

	result, _ := client.Query(query)
	//spew.Dump(result.Ways)
	for _, way := range result.Ways {
		building := getBuildingFromRaw(way.Tags, way.Geometry)
		if building != nil {
			buildings = append(buildings, building)
		}
	}
	for _, relation := range result.Relations {
		//todo: geometry from relation->way
		building := getBuildingFromRaw(relation.Tags, nil)
		if building != nil {
			buildings = append(buildings, building)
		}
	}

	for _, node := range result.Nodes {
		building := getBuildingFromRaw(node.Tags, nil)
		if building != nil {
			buildings = append(buildings, building)
		}
	}
	return
}

func getBuildingFromRaw(tags map[string]string, rawGeometry []overpass.Point) *model.GeoBuilding {
	var building model.GeoBuilding
	building.HouseNumber = tags["addr:housenumber"]
	building.Street = tags["addr:street"]
	if building.Street == "" {
		return nil
	}

	var geometry []model.GeoPoint
	for _, point := range rawGeometry {
		geometry = append(geometry, model.GeoPoint{Lat: point.Lat, Lon: point.Lat})
	}
	building.Geometry = geometry
	return &building
}
