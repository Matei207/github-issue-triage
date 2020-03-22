package helper

import "go.uber.org/zap"

func NewZapSugaredLogger() *zap.SugaredLogger {
	log, _ := zap.NewProduction()
	return log.Sugar()
}
