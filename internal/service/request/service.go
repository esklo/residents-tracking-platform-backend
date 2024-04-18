package request

import (
	"context"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	def "github.com/esklo/residents-tracking-platform-backend/internal/service"
	"github.com/google/uuid"
	"github.com/paulmach/orb/geojson"
	"time"
)

var _ def.RequestService = (*Service)(nil)

type Service struct {
	requestRepository repository.RequestRepository
	contactService    def.ContactService
	fileService       def.FileService
}

func NewService(
	requestRepository repository.RequestRepository,
	contactService def.ContactService,
	fileService def.FileService,
) *Service {
	return &Service{
		requestRepository: requestRepository,
		contactService:    contactService,
		fileService:       fileService,
	}
}

func (s *Service) Create(ctx context.Context, themeId *uuid.UUID, description, address string, contact *model.Contact, geo model.GeoPoint, fileIds []*uuid.UUID) (*model.Request, error) {
	contactCreated, err := s.contactService.Create(ctx, contact)
	if err != nil {
		return nil, err
	}
	request, err := s.requestRepository.Create(ctx, &model.Request{
		Description: description,
		Geo:         geo,
		Address:     address,
		CreatedAt:   time.Now(),
		Status:      model.RequestStatusOpen,
		ThemeId:     themeId,
		Contact:     contactCreated,
	})
	if err != nil {
		return nil, err
	}
	for _, fileId := range fileIds {
		file, err := s.fileService.GetById(ctx, fileId)
		if err != nil {
			return nil, err
		}
		request.Files = append(request.Files, file)
		err = s.requestRepository.AddFile(ctx, request.Id.String(), fileId.String())
		if err != nil {
			return nil, err
		}
	}
	return request, nil
}

func (s *Service) Get(ctx context.Context, id *uuid.UUID) (*model.Request, error) {
	request, err := s.requestRepository.GetByID(ctx, id.String())
	if err != nil {
		return nil, err
	}
	contact, err := s.contactService.Get(ctx, &request.Contact.Id)
	if err != nil {
		return nil, err
	}
	request.Contact = contact
	fileIds, err := s.requestRepository.GetFiles(ctx, id.String())
	if err != nil {
		return nil, err
	}
	var files []*model.File
	for _, fileId := range fileIds {
		file, err := s.fileService.GetById(ctx, fileId)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	request.Files = files
	return request, nil
}

func (s *Service) GetAll(ctx context.Context) ([]*model.Request, error) {
	requests, err := s.requestRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, request := range requests {
		contact, err := s.contactService.Get(ctx, &request.Contact.Id)
		if err != nil {
			return nil, err
		}
		request.Contact = contact

		fileIds, err := s.requestRepository.GetFiles(ctx, request.Id.String())
		if err != nil {
			return nil, err
		}
		var files []*model.File
		for _, fileId := range fileIds {
			file, err := s.fileService.GetById(ctx, fileId)
			if err != nil {
				return nil, err
			}
			files = append(files, file)
		}
		request.Files = files
	}
	return requests, nil
}

func (s *Service) GetAllAsGeoJson(ctx context.Context) ([]byte, error) {
	requests, err := s.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	collection := geojson.NewFeatureCollection()
	for _, request := range requests {
		feature, err := request.AsGeoJson()
		if err != nil {
			return nil, err
		}
		collection.Append(feature)
	}
	return collection.MarshalJSON()
}
