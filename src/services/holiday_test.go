package services

import (
	"github.com/crobatair/holiday-api/src/domain/dto"
	"testing"
	"time"
)

type HolidaysRepositoryStubTest struct {
	holidays []dto.Holiday
}

func (h HolidaysRepositoryStubTest) FindAllHolidays(f dto.HolidayFilters) (*[]dto.Holiday, *error) {
	return &h.holidays, nil
}

func TestDefaultClient(t *testing.T) {
	stubRepository := &HolidaysRepositoryStubTest{
		holidays: []dto.Holiday{
			{
				Title:      "test",
				Date:       time.Time{},
				Type:       "test",
				Inalinable: true,
				Extra:      "test",
			},
		},
	}

	service := NewHolidayService(stubRepository)
	f := dto.HolidayFilters{
		Extra: "test",
		FROM:  time.Time{},
		TO:    time.Time{},
	}
	holidays, err := service.FindAllHolidays(f)
	if err != nil {
		t.Error(err)
	}

	if len(*holidays) != 1 {
		t.Error("Expected 1 holiday, got ", len(*holidays))
	}
}
