package logger

import (
	"log"

	"go.uber.org/zap"
)

var I *zap.SugaredLogger

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("logger failed:", err)
	}
	I = logger.Sugar()
}
