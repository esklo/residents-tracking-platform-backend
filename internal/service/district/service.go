package district

import (
	"context"
	"fmt"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	def "github.com/esklo/residents-tracking-platform-backend/internal/service"
	"github.com/google/uuid"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/pkg/errors"
	"github.com/serjvanilla/go-overpass"
	"go.uber.org/zap"
	"slices"
)

var _ def.DistrictService = (*Service)(nil)

type Service struct {
	districtRepository repository.DistrictRepository
	fileRepository     repository.FileRepository
	logger             *zap.Logger
}

func NewService(
	districtRepository repository.DistrictRepository,
	fileRepository repository.FileRepository,
	logger *zap.Logger,
) *Service {
	return &Service{
		districtRepository: districtRepository,
		fileRepository:     fileRepository,
		logger:             logger,
	}
}

func (s *Service) Get(ctx context.Context, id *uuid.UUID) (*model.District, error) {
	district, err := s.districtRepository.GetByID(ctx, id.String())
	if err != nil {
		s.logger.Error("can not get district", zap.Error(err))
		return nil, err
	}
	if district == nil {
		s.logger.Error("district with id not found", zap.String("id", id.String()))
		return nil, model.ErrorNotFound
	}

	return district, nil
}

func (s *Service) Create(ctx context.Context, areaId int64, fileId *uuid.UUID) (*model.District, error) {
	if fileId != nil {
		_, err := s.fileRepository.GetByID(ctx, fileId.String())
		if err != nil {
			return nil, errors.Wrap(err, "can not get file by id")
		}
	}

	client := overpass.New()
	query := fmt.Sprintf(`
	[out:json][timeout:25];
	relation(%d)[admin_level=%d];
	(._;>;);
	out body;
	`, areaId, 8)
	result, err := client.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "overpass")
	}
	relation, ok := result.Relations[areaId]
	if !ok {
		return nil, errors.New("relation not found")
	}
	if len(relation.Members) == 0 {
		return nil, errors.New("members not found")
	}

	var multiPolygon orb.MultiPolygon
	var ring orb.Ring
	var proceededWays []int64
	currentMember := findNextMember(relation.Members, nil, proceededWays)
	for currentMember != nil {
		nodes := currentMember.Way.Nodes
		//log.Printf("w.id: %#v; %#v; %#v", currentMember.Way.ID, nodes[0].ID, nodes[len(nodes)-1].ID)
		for i, node := range nodes {
			point := orb.Point{node.Lon, node.Lat}
			ring = append(ring, point)
			if len(currentMember.Way.Nodes) == i+1 && ring[0].Equal(point) {
				//log.Printf("EQUAL!")
				if ring.Orientation() == orb.CCW {
					ring.Reverse()
				}
				multiPolygon = append(multiPolygon, orb.Polygon{ring})
				ring = orb.Ring{}
			}
		}
		proceededWays = append(proceededWays, currentMember.Way.ID)
		currentMember = findNextMember(relation.Members, currentMember, proceededWays)
	}
	feature := geojson.NewFeature(multiPolygon)

	props := make(geojson.Properties, len(relation.Tags))
	for k, v := range relation.Tags {
		props[k] = v
	}
	feature.Properties = props

	geojsonBytes, err := feature.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "can not marshal geojson")
	}
	district, err := s.districtRepository.Create(ctx, &model.District{
		Title:            relation.Tags["name"],
		GeoJson:          geojsonBytes,
		CoatOfArmsFileId: fileId,
	})
	if err != nil {
		s.logger.Error("can not create district", zap.Error(err))
		return nil, err
	}

	return district, nil
}

func findNextMember(members []overpass.RelationMember, currentMember *overpass.RelationMember, proceededWays []int64) *overpass.RelationMember {
	for i, member := range members {
		if member.Way == nil || member.Way.Nodes == nil {
			members = append(members[:i], members[i+1:]...)
		}
	}

	for i, member := range members {
		if member.Way == nil || member.Way.Nodes == nil || slices.Contains(proceededWays, member.Way.ID) {
			continue
		}
		if currentMember == nil {
			return &member
		}
		nodes := member.Way.Nodes
		currentNodes := currentMember.Way.Nodes
		nextExpectedNodeId := currentNodes[len(currentNodes)-1].ID
		if nextExpectedNodeId == nodes[0].ID {
			return &member
		}
		if nextExpectedNodeId == nodes[len(nodes)-1].ID {
			slices.Reverse(member.Way.Nodes)
			return &member
		}
		if len(members)-1 == i {
			return &member
		}
	}
	return nil
}

func (s *Service) GetAll(ctx context.Context) ([]*model.District, error) {
	districts, err := s.districtRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return districts, nil
}
