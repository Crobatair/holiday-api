package services

import "github.com/crobatair/holiday-api/src/domain/dto"

type HolidayService struct {
	repository dto.HolidayRepository
}

func (r HolidayService) FindAllHolidays(filters dto.HolidayFilters) (*[]dto.Holiday, *error) {
	return r.repository.FindAllHolidays(filters)
}

func NewHolidayService(repository dto.HolidayRepository) HolidayService {
	return HolidayService{
		repository: repository,
	}
}
