package model

import (
	protoAnalytics "github.com/esklo/residents-tracking-platform-backend/gen/proto/analytics"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type RequestsPerTheme struct {
	Theme Theme
	Count uint64
}

func (r *RequestsPerTheme) ToProto() (*protoAnalytics.RequestsPerTheme, error) {
	if r == nil {
		return nil, ErrorModelIsEmpty
	}
	themeProto, err := r.Theme.ToProto()
	if err != nil {
		return nil, err
	}
	requestPerTheme := protoAnalytics.RequestsPerTheme{
		Count: int64(r.Count),
		Theme: themeProto,
	}
	return &requestPerTheme, nil
}

type StatsElement struct {
	WeekCount, MonthCount, OpenCount uint64
	WeekDelta, MonthDelta            float64
}

func (s *StatsElement) ToProto() (*protoAnalytics.StatsElement, error) {
	if s == nil {
		return nil, ErrorModelIsEmpty
	}

	statsProto := protoAnalytics.StatsElement{
		WeekCount:  int64(s.WeekCount),
		WeekDelta:  float32(s.WeekDelta),
		MonthCount: int64(s.MonthCount),
		MonthDelta: float32(s.MonthDelta),
		OpenCount:  int64(s.OpenCount),
	}
	return &statsProto, nil
}

type RequestsPerThemePerDateElement struct {
	RequestsPerTheme
	Date time.Time
}

func (r *RequestsPerThemePerDateElement) ToProto() (*protoAnalytics.RequestsPerThemePerDate, error) {
	if r == nil {
		return nil, ErrorModelIsEmpty
	}
	themeProto, err := r.Theme.ToProto()
	if err != nil {
		return nil, err
	}
	requestPerTheme := protoAnalytics.RequestsPerThemePerDate{
		Count: int64(r.Count),
		Theme: themeProto,
		Date:  timestamppb.New(r.Date),
	}
	return &requestPerTheme, nil
}
