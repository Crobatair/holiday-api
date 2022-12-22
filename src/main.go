package main

import (
	"github.com/crobatair/holiday-api/src/app/handlers"
	"github.com/crobatair/holiday-api/src/config"
	"github.com/crobatair/holiday-api/src/domain/repository"
	"github.com/crobatair/holiday-api/src/services"
	"github.com/gin-gonic/gin"
)

func main() {
	config.DefaultLogger()
	config.NewGenericHolidayConfig()

	hRepository := repository.NewHolidayRepository(config.Get().ServiceURL)
	hService := services.NewHolidayService(hRepository)

	r := gin.Default()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/holiday", handlers.HolidayHandler(hService))
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
