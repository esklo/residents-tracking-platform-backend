package geo

import (
	"context"
	"github.com/esklo/residents-tracking-platform/gen/proto/empty"
	proto "github.com/esklo/residents-tracking-platform/gen/proto/geo"
	"github.com/esklo/residents-tracking-platform/internal/model"
	"github.com/esklo/residents-tracking-platform/internal/service"
	"strconv"
)

type Implementation struct {
	proto.UnimplementedGeoServiceServer
	geoService  service.GeoService
	authService service.AuthService
}

func NewImplementation(geoService service.GeoService, authService service.AuthService) *Implementation {
	return &Implementation{
		geoService:  geoService,
		authService: authService,
	}
}

func (i Implementation) BuildingsByCoordinates(ctx context.Context, req *proto.GeoPoint) (*proto.BuildingByCoordinatesResponse, error) {
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	buildings, err := i.geoService.BuildingsByCoordinates(ctx, float64(req.Latitude), float64(req.Longitude))
	if err != nil {
		return nil, err
	}
	var protoBuildings []*proto.GeoBuilding
	for _, building := range buildings {
		var geometry []*proto.GeoPoint
		for _, point := range building.Geometry {
			geometry = append(geometry, &proto.GeoPoint{
				Latitude:  float32(point.Lat),
				Longitude: float32(point.Lon),
			})
		}
		protoBuilding := &proto.GeoBuilding{
			Street:   building.Street,
			House:    building.HouseNumber,
			Geometry: geometry,
		}
		protoBuildings = append(protoBuildings, protoBuilding)
	}
	return &proto.BuildingByCoordinatesResponse{
		Buildings: protoBuildings,
	}, nil
}

func (i Implementation) GetAdministrativeDistricts(ctx context.Context, _ *empty.Empty) (*proto.GetDistrictsResponse, error) {
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	//102269 – moscow osm id
	administrativeDistricts, err := i.geoService.GetDistricts(ctx, 102269, 5)
	if err != nil {
		return nil, err
	}

	var protoAdministrativeDistricts []*proto.GeoDistrict
	for _, administrativeDistrict := range administrativeDistricts {
		protoAdministrativeDistricts = append(protoAdministrativeDistricts, &proto.GeoDistrict{
			Id:   administrativeDistrict.Id,
			Name: administrativeDistrict.Name,
		})
	}
	return &proto.GetDistrictsResponse{Districts: protoAdministrativeDistricts}, err
}

func (i Implementation) GetDistricts(ctx context.Context, req *proto.GetDistrictsRequest) (*proto.GetDistrictsResponse, error) {
	_, err := i.authService.ExchangeTokenFromContext(ctx)
	if err != nil {
		return nil, model.ErrorUnauthenticated
	}

	districts, err := i.geoService.GetDistricts(ctx, req.AdministrativeDistrictId, 8)
	if err != nil {
		return nil, err
	}

	var protoDistricts []*proto.GeoDistrict
	for _, district := range districts {
		protoDistricts = append(protoDistricts, &proto.GeoDistrict{
			Id:   district.Id,
			Name: district.Name,
		})
	}
	return &proto.GetDistrictsResponse{Districts: protoDistricts}, err
}

func (i Implementation) Suggest(ctx context.Context, req *proto.SuggestRequest) (*proto.SuggestResponse, error) {
	suggestions, err := i.geoService.Suggest(ctx, req.Query)
	if err != nil {
		return nil, err
	}
	var items []*proto.SuggestItem
	for _, suggestion := range suggestions {
		latitude, err := strconv.ParseFloat(suggestion.Data.GeoLat, 64)
		if err != nil {
			return nil, err
		}

		longitude, err := strconv.ParseFloat(suggestion.Data.GeoLon, 64)
		if err != nil {
			return nil, err
		}

		items = append(items, &proto.SuggestItem{
			Address: suggestion.Value,
			Geo: &proto.GeoPoint{
				Latitude:  float32(latitude),
				Longitude: float32(longitude),
			},
		})
	}
	return &proto.SuggestResponse{Items: items}, nil
}

func (i Implementation) GeoLocate(ctx context.Context, req *proto.GeoPoint) (*proto.GeoLocateResponse, error) {
	suggestion, err := i.geoService.Locate(ctx, float64(req.Latitude), float64(req.Longitude))
	if err != nil {
		return nil, err
	}
	latitude, err := strconv.ParseFloat(suggestion.Data.GeoLat, 64)
	if err != nil {
		return nil, err
	}

	longitude, err := strconv.ParseFloat(suggestion.Data.GeoLon, 64)
	if err != nil {
		return nil, err
	}
	return &proto.GeoLocateResponse{Item: &proto.SuggestItem{
		Address: suggestion.Value,
		Geo: &proto.GeoPoint{
			Latitude:  float32(latitude),
			Longitude: float32(longitude),
		},
	}}, nil
}
