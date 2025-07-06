package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	BlueColor   = "\033[34m"
	YellowColor = "\033[33m"
	RedColor    = "\033[31m"
	ResetColor  = "\033[0m"
)

func myEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.InfoLevel:
		enc.AppendString(BlueColor + "INFO" + ResetColor)
	case zapcore.WarnLevel:
		enc.AppendString(YellowColor + "WARN" + ResetColor)
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		enc.AppendString(RedColor + "ERROR" + ResetColor)
	default:
		enc.AppendString(level.String())
	}
}

func createConsoleEncoder() zapcore.Encoder {
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	consoleEncoderConfig.EncodeLevel = myEncodeLevel
	return zapcore.NewConsoleEncoder(consoleEncoderConfig)
}

func createFileEncoder() zapcore.Encoder {
	jsonEncoderConfig := zap.NewProductionEncoderConfig()
	jsonEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	return zapcore.NewJSONEncoder(jsonEncoderConfig)
}
