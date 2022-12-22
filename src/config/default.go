package config

import "github.com/crobatair/holiday-api/src/common"

func DefaultLogger() {
	common.InitLogger()
	//	common.Info("Startup")
}
