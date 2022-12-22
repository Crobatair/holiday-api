package config

import "os"

var c HolidayConfig

type HolidayConfig struct {
	ServiceURL string
}

func NewGenericHolidayConfig() {
	url := os.Getenv("HOLIDAY_SERVICE_URL")
	if url == "" {
		url = "https://api.victorsanmartin.com/feriados/en.json"
	}
	c = HolidayConfig{
		ServiceURL: url,
	}

}

func Get() HolidayConfig {
	return c
}
