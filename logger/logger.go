package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Env string // dev, prod
}

var logger *zap.Logger

func Init(cfg Config) {
	var err error

	// создание дирректории и файла логирования

	dir := "logs"
	now := time.Now()
	timeStamp := now.Format("2006-01-02 15:04:05.000")
	fileName := fmt.Sprintf("log-%s.log", timeStamp)
	filePath := filepath.Join(dir, fileName)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

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

func newDevLogger(file zapcore.WriteSyncer) *zap.Logger {
	encoderCfgConsole := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		MessageKey:    "msg",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",

		EncodeLevel:  zapcore.CapitalColorLevelEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	encoderCfgFile := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		MessageKey:    "msg",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",

		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfgConsole),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfgFile),
		zapcore.AddSync(file),
		zapcore.DebugLevel,
	)

	core := zapcore.NewTee(consoleCore, fileCore)

	return zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)
}

func newProdLogger(file zapcore.WriteSyncer) *zap.Logger {
	encoderCfgConsole := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		MessageKey:    "msg",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",

		EncodeLevel:  zapcore.CapitalColorLevelEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	encoderCfgFile := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		MessageKey:    "msg",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",

		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfgConsole),
		zapcore.AddSync(os.Stdout),
		zapcore.InfoLevel,
	)

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfgFile),
		zapcore.AddSync(file),
		zapcore.InfoLevel,
	)

	core := zapcore.NewTee(consoleCore, fileCore)

	return zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)
}
