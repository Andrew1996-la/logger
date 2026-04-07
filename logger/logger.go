package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Env string // dev, prod
}

var logger *zap.Logger

func Init(cfg Config) {
	var err error

	if cfg.Env == "dev" {
		logger = newDevLogger()
	} else {
		logger, err = zap.NewProduction()
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

func newDevLogger() *zap.Logger {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:     "time",
		LevelKey:    "level",
		MessageKey:  "msg",
		CallerKey:   "caller",
		StacktraceKey: "stacktrace",

		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime:  zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	return zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)
}
