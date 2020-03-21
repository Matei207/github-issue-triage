package main

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	log, _ := zap.NewProduction()
	logger = log.Sugar()
}

func main() {
	defer logger.Sync()


}