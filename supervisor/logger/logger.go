package logger

import "go.uber.org/zap"

var I *zap.Logger

func init() {
	var err error
	I, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}
