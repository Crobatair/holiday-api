package common

import (
	"go.uber.org/zap"
)

var log *zap.Logger

func InitLogger() {
	log, _ := zap.NewProduction()
	defer log.Sync()
}

func Info(topic string) {
	log.Info(topic)
}

func Warning(message string) {
	log.Error(message)
}
