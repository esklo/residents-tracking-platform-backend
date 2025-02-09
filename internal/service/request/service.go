package request

import (
	"context"
	"fmt"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/esklo/residents-tracking-platform-backend/internal/repository"
	def "github.com/esklo/residents-tracking-platform-backend/internal/service"
	"github.com/google/uuid"
	"github.com/paulmach/orb/geojson"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"reflect"
	"time"
)

var _ def.RequestService = (*Service)(nil)

type Service struct {
	requestRepository repository.RequestRepository
	contactService    def.ContactService
	fileService       def.FileService
	themeService      def.ThemeService
	departmentService def.DepartmentService
	userService       def.UserService
	logger            *zap.Logger
}

func NewService(
	requestRepository repository.RequestRepository,
	contactService def.ContactService,
	fileService def.FileService,
	themeService def.ThemeService,
	departmentService def.DepartmentService,
	userService def.UserService,
	logger *zap.Logger,
) *Service {
	return &Service{
		requestRepository: requestRepository,
		contactService:    contactService,
		fileService:       fileService,
		themeService:      themeService,
		departmentService: departmentService,
		userService:       userService,
		logger:            logger,
	}
}

func (s *Service) Create(ctx context.Context, request *model.Request) (*model.Request, error) {
	var createdAt time.Time
	createdAt = time.Now()

	contactCreated, err := s.contactService.Create(ctx, request.Contact)
	if err != nil {
		return nil, err
	}
	request.Contact = contactCreated

	request.CreatedAt = createdAt
	request.Status = model.RequestStatusOpen

	request, err = s.requestRepository.Create(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, file := range request.Files {
		fileModel, err := s.fileService.GetById(ctx, &file.Id)
		if err != nil {
			return nil, err
		}
		request.Files = append(request.Files, fileModel)
		err = s.requestRepository.AddFile(ctx, &request.Id, &file.Id)
		if err != nil {
			return nil, err
		}
	}

	for _, file := range request.ReportFiles {
		fileModel, err := s.fileService.GetById(ctx, &file.Id)
		if err != nil {
			return nil, err
		}
		request.ReportFiles = append(request.ReportFiles, fileModel)
		err = s.requestRepository.AddReportFile(ctx, &request.Id, &file.Id)
		if err != nil {
			return nil, err
		}
	}
	return request, nil
}

func (s *Service) Get(ctx context.Context, id *uuid.UUID) (*model.Request, error) {
	request, err := s.requestRepository.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "can not get request by id")
	}
	contact, err := s.contactService.Get(ctx, &request.Contact.Id)
	if err != nil {
		return nil, errors.Wrap(err, "can not get contact by id")
	}
	request.Contact = contact
	fileIds, err := s.requestRepository.GetFiles(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "can not get files by id")
	}
	var files []*model.File
	for _, fileId := range fileIds {
		file, err := s.fileService.GetById(ctx, fileId)
		if err != nil {
			return nil, errors.Wrap(err, "can not get file by id")
		}
		files = append(files, file)
	}
	request.Files = files

	reportFileIds, err := s.requestRepository.GetReportFiles(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "can not get report files by id")
	}
	var reportFiles []*model.File
	for _, reportFileId := range reportFileIds {
		file, err := s.fileService.GetById(ctx, reportFileId)
		if err != nil {
			return nil, errors.Wrap(err, "can not get report file by id")
		}
		reportFiles = append(reportFiles, file)
	}
	request.ReportFiles = reportFiles
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

		fileIds, err := s.requestRepository.GetFiles(ctx, &request.Id)
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

		reportFileIds, err := s.requestRepository.GetReportFiles(ctx, &request.Id)
		if err != nil {
			return nil, err
		}
		var reportFiles []*model.File
		for _, fileId := range reportFileIds {
			file, err := s.fileService.GetById(ctx, fileId)
			if err != nil {
				return nil, err
			}
			reportFiles = append(reportFiles, file)
		}
		request.ReportFiles = reportFiles
	}
	return requests, nil
}

func (s *Service) GetAllWithDepartment(ctx context.Context, departmentId *uuid.UUID) ([]*model.Request, error) {
	department, err := s.departmentService.Get(ctx, departmentId)
	if department.FullAccess {
		return s.GetAll(ctx)
	}

	themes, err := s.themeService.GetAllWithDepartment(ctx, departmentId)
	if err != nil {
		return nil, err
	}
	var themeIds []*uuid.UUID
	for _, theme := range themes {
		themeIds = append(themeIds, &theme.Id)
	}
	requests, err := s.requestRepository.GetAllWithThemeIds(ctx, themeIds)

	if err != nil {
		return nil, err
	}
	for _, request := range requests {
		contact, err := s.contactService.Get(ctx, &request.Contact.Id)
		if err != nil {
			return nil, err
		}
		request.Contact = contact

		fileIds, err := s.requestRepository.GetFiles(ctx, &request.Id)
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

		reportFileIds, err := s.requestRepository.GetFiles(ctx, &request.Id)
		if err != nil {
			return nil, err
		}
		var reportFiles []*model.File
		for _, fileId := range reportFileIds {
			file, err := s.fileService.GetById(ctx, fileId)
			if err != nil {
				return nil, err
			}
			reportFiles = append(reportFiles, file)
		}
		request.ReportFiles = reportFiles
	}
	return requests, nil
}

func (s *Service) GetAllAsGeoJsonWithDepartment(ctx context.Context, department *uuid.UUID) ([]byte, error) {
	requests, err := s.GetAllWithDepartment(ctx, department)
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

func (s *Service) GetCountWithThemeId(ctx context.Context, from time.Time, to time.Time, themeId *uuid.UUID) (float64, error) {
	return s.requestRepository.GetCountWithThemeId(ctx, from, to, themeId)
}
func (s *Service) GetCountWithThemeIdAndStatus(ctx context.Context, themeId *uuid.UUID, status model.RequestStatus) (float64, error) {
	return s.requestRepository.GetCountWithThemeIdAndStatus(ctx, themeId, int(status))
}

func (s *Service) Update(ctx context.Context, request *model.Request) error {
	requestOG, err := s.requestRepository.GetByID(ctx, &request.Id)
	if err != nil {
		return errors.Wrap(err, "can not get requestOG")
	}

	err = s.requestRepository.Update(ctx, request)

	if err != nil {
		return errors.Wrap(err, "can not update request")
	}

	if requestOG.Contact != nil {
		err := s.contactService.Update(ctx, &model.Contact{
			Id:    requestOG.Contact.Id,
			Phone: request.Contact.Phone,
			Email: request.Contact.Email,
			Name:  request.Contact.Name,
			Note:  request.Contact.Note,
		})
		if err != nil {
			return errors.Wrap(err, "can not update contact")
		}
	}

	currentFileIds, err := s.requestRepository.GetFiles(ctx, &request.Id)
	if err != nil {
		return errors.Wrap(err, "can not get request's files")
	}
	for _, fileId := range currentFileIds {
		err := s.requestRepository.RemoveFile(ctx, &request.Id, fileId)
		if err != nil {
			return errors.Wrap(err, "can not remove file")
		}
	}
	for _, file := range request.Files {
		fileModel, err := s.fileService.GetById(ctx, &file.Id)
		if err != nil {
			return errors.Wrap(err, "can not get file by id")
		}
		request.Files = append(request.Files, fileModel)
		err = s.requestRepository.AddFile(ctx, &request.Id, &file.Id)
		if err != nil {
			return errors.Wrap(err, "can not add file")
		}
	}

	currentReportFileIds, err := s.requestRepository.GetReportFiles(ctx, &request.Id)
	if err != nil {
		return errors.Wrap(err, "can not get request's report files")
	}
	s.logger.Warn("currentReportFileIds", zap.Any("currentReportFileIds", currentReportFileIds))
	for _, fileId := range currentReportFileIds {
		err := s.requestRepository.RemoveReportFile(ctx, &request.Id, fileId)
		if err != nil {
			return errors.Wrap(err, "can not remove report file")
		}
	}
	for _, file := range request.ReportFiles {
		fileModel, err := s.fileService.GetById(ctx, &file.Id)
		if err != nil {
			return errors.Wrap(err, "can not get report file by id")
		}
		request.ReportFiles = append(request.ReportFiles, fileModel)
		err = s.requestRepository.AddReportFile(ctx, &request.Id, &file.Id)
		if err != nil {
			return errors.Wrap(err, "can not add report file")
		}
	}
	return nil

}

func (s *Service) ExportExcel(ctx context.Context, departmentId *uuid.UUID) (*model.File, error) {
	sheet := "Sheet1"

	f := excelize.NewFile()
	defer f.Close()

	headers := []string{
		"id",
		"status",
		"theme_id",
		"theme_name",
		"theme_department_id",
		"theme_department_name",
		"user_id",
		"user_last_name",
		"user_first_name",
		"user_father_name",
		"user_email",
		"description",
		"address",
		"lon",
		"lat",
		"contact_id",
		"contact_name",
		"contact_phone",
		"contact_email",
		"contact_note",
		"created_at",
	}
	for i, header := range headers {
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}

	requests, err := s.GetAllWithDepartment(ctx, departmentId)
	if err != nil {
		return nil, err
	}

	for i, request := range requests {
		theme, _ := s.themeService.Get(ctx, request.ThemeId)
		dep, _ := s.departmentService.Get(ctx, &theme.DepartmentId)
		request.Contact, _ = s.contactService.Get(ctx, &request.Contact.Id)
		var row []any
		row = appendExcelRow(row, request.Id)
		row = appendExcelRow(row, request.Status)
		row = appendExcelRow(row, request.ThemeId)
		row = appendExcelRow(row, theme.Title)
		row = appendExcelRow(row, theme.DepartmentId)
		row = appendExcelRow(row, dep.Title)
		row = appendExcelRow(row, request.UserId)
		if request.UserId != nil {
			user, _ := s.userService.Get(ctx, request.UserId)
			row = appendExcelRow(row, *user.LastName)
			row = appendExcelRow(row, user.FirstName)
			row = appendExcelRow(row, *user.FatherName)
			row = appendExcelRow(row, user.Email)
		} else {
			row = appendExcelRow(row, nil)
			row = appendExcelRow(row, nil)
			row = appendExcelRow(row, nil)
			row = appendExcelRow(row, nil)
		}
		row = appendExcelRow(row, request.Description)
		row = appendExcelRow(row, request.Address)
		row = appendExcelRow(row, request.Geo.Lon)
		row = appendExcelRow(row, request.Geo.Lat)
		row = appendExcelRow(row, request.Contact.Id)
		row = appendExcelRow(row, request.Contact.Name)
		row = appendExcelRow(row, request.Contact.Phone)
		row = appendExcelRow(row, request.Contact.Email)
		row = appendExcelRow(row, request.Contact.Note)
		row = appendExcelRow(row, request.CreatedAt.Format("02.01.2006 15:04:05"))
		for j, r := range row {
			f.SetCellValue(sheet, fmt.Sprintf("%s%d", string(rune(65+j)), i+2), r)
		}
	}

	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return s.fileService.Upload(
		ctx,
		fmt.Sprintf("requests_export_%s.xlsx", time.Now().Format("02_01_2006_15_04_05")),
		buf.Bytes(),
	)
}

func appendExcelRow(row []any, value any) []any {
	val := reflect.ValueOf(value)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return append(row, "")
		}
		return append(row, val.Elem().Interface())
	}
	return append(row, value)
}

func (s *Service) Delete(ctx context.Context, id *uuid.UUID) error {
	request, err := s.requestRepository.GetByID(ctx, id)
	if err != nil {
		return errors.Wrap(err, "can not get request by id")
	}
	//for _, fileId := range files {
	//	if err := s.requestRepository.RemoveFile(ctx, id, fileId); err != nil {
	//		return errors.Wrap(err, "can not remove file")
	//	}
	//
	//	if err := s.fileService.Delete(ctx, fileId); err != nil {
	//		return errors.Wrap(err, "can not delete file")
	//	}
	//}
	//return s.requestRepository.Delete(ctx, id)
	t := time.Now()
	request.DeletedAt = &t
	return s.requestRepository.Update(ctx, request)
}
