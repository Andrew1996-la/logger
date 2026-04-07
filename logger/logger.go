package logger

import (
	"log"

	"go.uber.org/zap"
)

type Config struct {
	Env string // dev, prod
}

var logger *zap.Logger

func Init(cfg Config) {
	var err error

	file := initLogFile()

	if cfg.Env == "dev" {
		logger = newDevLogger(file)
	} else {
		logger = newProdLogger(file)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func Sync() {
	logger.Sync()
}

func Info(msg string, fields ...Field) {
	logger.Info(msg, fields...)
}

func Error(msg string, fields ...Field) {
	logger.Error(msg, fields...)
}
