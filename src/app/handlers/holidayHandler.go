package handlers

import (
	"github.com/crobatair/holiday-api/src/domain/dto"
	"github.com/crobatair/holiday-api/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// @BasePath /api/v1

// Get Holidays
// @Summary Retrieve a list of Holidays
// @Schemes
// @Description Retrieve a list of Holidays
// @Produce json
// @Success 200
// @Router /api/v1/holidays [get]
func HolidayHandler(holidayService services.HolidayService) func(c *gin.Context) {
	return func(c *gin.Context) {
		filters := dto.GetFiltersFromRequest(c)
		holidays, _ := holidayService.FindAllHolidays(filters)

		renderer := c.GetHeader("Accept")

		if renderer == "application/json" || renderer == "*/*" {
			c.JSON(http.StatusOK, dto.DTOArrayToResponseJSON(*holidays))

		} else if strings.HasPrefix(renderer, "application/xml") {
			c.XML(http.StatusOK, dto.DTOArrayToResponseJSON(*holidays))
		} else {
			c.JSON(http.StatusOK, dto.DTOArrayToResponseJSON(*holidays))
		}
	}
}
