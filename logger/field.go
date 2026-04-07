package logger

import "go.uber.org/zap"

type Field = zap.Field

func String(key, value string) Field {
	return zap.String(key, value)
}

func Err (err error) Field {
	return zap.Error(err)
}

func Int (key string, value int) Field {
	return zap.Int(key, value)
}
