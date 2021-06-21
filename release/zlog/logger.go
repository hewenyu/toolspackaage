package zlog

import "go.uber.org/zap"

var loggers = Zap()

func Debug(msg string, fields ...zap.Field) {
	loggers.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	loggers.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	loggers.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	loggers.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	loggers.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	loggers.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	loggers.Fatal(msg, fields...)
}
