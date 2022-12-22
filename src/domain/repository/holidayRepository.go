package repository

import (
	"encoding/json"
	"github.com/crobatair/holiday-api/src/domain/dto"
	"io"
	"net/http"
	"strings"
	"time"
)

type HolidaysRepositoryStub struct {
	holidays []dto.Holiday
}

func (h HolidaysRepositoryStub) FindAllHolidays(f dto.HolidayFilters) (*[]dto.Holiday, *error) {
	results := filterOutExtra(h.holidays, f.Extra)
	if f.GetFromDate().IsZero() == false && f.GetToDate().IsZero() == false {
		results = filterByDate(results, f.GetFromDate(), f.GetToDate())
	}

	return &results, nil
}

func filterOutExtra(holidays []dto.Holiday, extra string) []dto.Holiday {
	results := make([]dto.Holiday, 0)
	for _, holiday := range holidays {
		if strings.Contains(holiday.Extra, extra) {
			results = append(results, holiday)
		}
	}
	return results
}

func filterByDate(holidays []dto.Holiday, from time.Time, to time.Time) []dto.Holiday {

	results := make([]dto.Holiday, 0)

	for _, holiday := range holidays {
		if holiday.GetDate().After(from) && holiday.GetDate().Before(to) {
			results = append(results, holiday)
		}

	}
	return results
}

func NewHolidayRepository(serviceUrl string) dto.HolidayRepository {
	client := http.Client{Timeout: time.Second * 3}
	req, err := http.NewRequest(http.MethodGet, serviceUrl, nil)
	if err != nil {
		return &HolidaysRepositoryStub{
			holidays: []dto.Holiday{},
		}
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		return &HolidaysRepositoryStub{
			holidays: []dto.Holiday{},
		}
	}

	if res.Body != nil {
		defer func() {
			err := res.Body.Close()
			if err != nil {
				println(err.Error())
			}
		}()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return &HolidaysRepositoryStub{
			holidays: []dto.Holiday{},
		}
	}

	response := GetDTOFromResponse(body)
	if response == nil {
		println("response is nil")
		return &HolidaysRepositoryStub{
			holidays: []dto.Holiday{},
		}
	}

	for _, holiday := range response.Data {
		println(holiday.Title)
	}

	return &HolidaysRepositoryStub{
		holidays: response.Data,
	}
}

func GetDTOFromResponse(body []byte) *dto.ResponseJson {
	r := dto.ResponseJson{}
	err := json.Unmarshal(body, &r)
	if err != nil {
		return nil
	}
	return &r
}
