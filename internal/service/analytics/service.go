package analytics

import (
	"context"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	def "github.com/esklo/residents-tracking-platform-backend/internal/service"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

var _ def.AnalyticsService = (*Service)(nil)

type Service struct {
	themeService      def.ThemeService
	requestService    def.RequestService
	departmentService def.DepartmentService
	logger            *zap.Logger
}

func NewService(
	themeService def.ThemeService,
	requestService def.RequestService,
	departmentService def.DepartmentService,
	logger *zap.Logger,
) *Service {
	return &Service{
		themeService:      themeService,
		requestService:    requestService,
		departmentService: departmentService,
		logger:            logger,
	}
}

func (s *Service) GetDepartmentThemes(ctx context.Context, departmentId *uuid.UUID) ([]*model.Theme, error) {
	if departmentId == nil {
		return nil, errors.New("department id is nil")
	}
	department, err := s.departmentService.Get(ctx, departmentId)
	if err != nil {
		return nil, err
	}

	var themes []*model.Theme
	if department.FullAccess {
		themes, err = s.themeService.GetAll(ctx, &department.DistrictId)
		if err != nil {
			return nil, err
		}
	} else {
		themes, err = s.themeService.GetAllWithDepartment(ctx, departmentId)
		if err != nil {
			return nil, err
		}
	}
	return themes, nil
}

func (s *Service) RequestsPerTheme(ctx context.Context, from time.Time, to time.Time, departmentId *uuid.UUID) ([]*model.RequestsPerTheme, error) {
	themes, err := s.GetDepartmentThemes(ctx, departmentId)
	if err != nil {
		return nil, err
	}
	var data []*model.RequestsPerTheme
	for _, theme := range themes {
		count, err := s.requestService.GetCountWithThemeId(ctx, from, to, &theme.Id)
		if err != nil {
			return nil, err
		}
		data = append(data, &model.RequestsPerTheme{
			Theme: *theme,
			Count: uint64(count),
		})
	}
	return data, nil
}

func (s *Service) RequestsPerThemePerDate(ctx context.Context, departmentId *uuid.UUID, from time.Time, to time.Time) (data []*model.RequestsPerThemePerDateElement, err error) {
	themes, err := s.GetDepartmentThemes(ctx, departmentId)
	if err != nil {
		return nil, err
	}

	year, month, day := from.Date()
	current := time.Date(year, month, day, 0, 0, 0, 0, from.Location())
	for current.Before(to) {
		to := current.Add(time.Hour * 24)
		for _, theme := range themes {
			count, err := s.requestService.GetCountWithThemeId(ctx, current, to, &theme.Id)
			if err != nil {
				return nil, err
			}
			data = append(data, &model.RequestsPerThemePerDateElement{
				RequestsPerTheme: model.RequestsPerTheme{
					Theme: *theme,
					Count: uint64(count),
				},
				Date: current,
			})
		}
		current = to
	}
	return data, nil
}

func (s *Service) Stats(ctx context.Context, departmentId *uuid.UUID) (*model.StatsElement, error) {
	themes, err := s.GetDepartmentThemes(ctx, departmentId)
	if err != nil {
		return nil, err
	}
	var weekCount float64
	now := time.Now()
	year, week := now.ISOWeek()
	//this week count
	for _, theme := range themes {
		from, to := weekRange(year, week)
		count, err := s.requestService.GetCountWithThemeId(ctx, from, to, &theme.Id)
		if err != nil {
			return nil, err
		}
		weekCount += count
	}

	var lastWeekCount float64
	//last week count
	for _, theme := range themes {
		lastWeek := week
		if lastWeek == 1 {
			lastWeek = 53
			year -= 1
		} else {
			lastWeek -= 1
		}
		from, to := weekRange(year, lastWeek)
		count, err := s.requestService.GetCountWithThemeId(ctx, from, to, &theme.Id)
		if err != nil {
			return nil, err
		}
		lastWeekCount += count
	}

	var weekDelta float64 = 0
	if lastWeekCount > 0 {
		weekDelta = (weekCount/lastWeekCount)*100 - 100
	}

	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	var monthCount float64
	//this month count
	for _, theme := range themes {
		from := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		to := from.AddDate(0, 1, -1)
		count, err := s.requestService.GetCountWithThemeId(ctx, from, to, &theme.Id)
		if err != nil {
			return nil, err
		}
		monthCount += count
	}

	var lastMonthCount float64
	//last month count
	for _, theme := range themes {
		if currentMonth == 1 {
			currentMonth = 12
			currentYear -= 1
		} else {
			currentMonth -= 1
		}
		from := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		to := from.AddDate(0, 1, -1)
		count, err := s.requestService.GetCountWithThemeId(ctx, from, to, &theme.Id)
		if err != nil {
			return nil, err
		}
		lastMonthCount += count
	}
	var monthDelta float64 = 0

	if lastMonthCount > 0 {
		monthDelta = (monthCount/lastMonthCount)*100 - 100
	}

	var openCount float64
	for _, theme := range themes {
		count, err := s.requestService.GetCountWithThemeIdAndStatus(ctx, &theme.Id, model.RequestStatusOpen)
		if err != nil {
			return nil, err
		}
		openCount += count
	}
	return &model.StatsElement{
		WeekCount:  uint64(weekCount),
		MonthCount: uint64(monthCount),
		OpenCount:  uint64(openCount),
		WeekDelta:  weekDelta,
		MonthDelta: monthDelta,
	}, nil
}

func weekStart(year, week int) time.Time {
	// Start from the middle of the year:
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}

func weekRange(year, week int) (start, end time.Time) {
	start = weekStart(year, week)
	end = start.AddDate(0, 0, 7)
	return
}
