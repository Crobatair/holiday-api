package dto

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Holiday struct {
	Title      string `json:"title"`
	Date       string `json:"date" time_format:"2006-01-31"`
	Type       string `json:"type"`
	Inalinable bool   `json:"inlinable"`
	Extra      string `json:"extra"`
}

type HolidayRepository interface {
	FindAllHolidays(HolidayFilters) (*[]Holiday, *error)
}

func (h Holiday) GetDate() time.Time {
	date, _ := time.Parse("2006-01-02", h.Date)
	return date
}

func DTOArrayToResponseJSON(holidays []Holiday) *ResponseJson {
	return &ResponseJson{
		Success: "true",
		Data:    holidays,
	}
}

type HolidayFilters struct {
	Extra string
	FROM  string
	TO    string
}

func GetFiltersFromRequest(c *gin.Context) HolidayFilters {
	return HolidayFilters{
		Extra: c.Query("extra"),
		FROM:  c.Query("from"),
		TO:    c.Query("to"),
	}
}

func (h HolidayFilters) GetFromDate() time.Time {
	from, _ := time.Parse(time.RFC3339, h.FROM)
	return from
}

func (h HolidayFilters) GetToDate() time.Time {
	to, _ := time.Parse(time.RFC3339, h.TO)
	return to
}
