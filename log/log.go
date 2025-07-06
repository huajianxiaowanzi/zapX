package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func Init(configs ...LogConfig) {
	config := getConfig(configs...)

	var cores []zapcore.Core

	if config.ConsoleOutput {
		consoleEncoder := createConsoleEncoder()
		cores = append(cores, createConsoleCore(config, consoleEncoder))
	}

	if config.FileOutput {
		fileEncoder := createFileEncoder()
		cores = append(cores, createFileCore(config, fileEncoder)...)
	}

	core := zapcore.NewTee(cores...)
	log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(log)
}

// 封装日志方法
func Info(msg string) {
	log.Info(msg)
}

func Infof(format string, args ...interface{}) {
	log.Sugar().Infof(format, args...)
}

func Error(msg string) {
	log.Error(msg)
}

func Errorf(format string, args ...interface{}) {
	log.Sugar().Errorf(format, args...)
}

func Warn(msg string) {
	log.Warn(msg)
}

func Warnf(format string, args ...interface{}) {
	log.Sugar().Warnf(format, args...)
}

func Debug(msg string) {
	log.Debug(msg)
}

func Debugf(format string, args ...interface{}) {
	log.Sugar().Debugf(format, args...)
}

func With(fields ...zap.Field) *zap.Logger {
	return log.With(fields...)
}

func WithError(err error) *zap.Logger {
	return log.With(zap.Error(err))
}
