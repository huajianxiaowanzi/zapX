package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func createConsoleCore(config LogConfig, encoder zapcore.Encoder) zapcore.Core {
	return zapcore.NewCore(
		encoder,
		zapcore.Lock(os.Stdout),
		config.Level,
	)
}

func createFileCore(config LogConfig, encoder zapcore.Encoder) []zapcore.Core {
	infoWS := getFileWriter(zapcore.InfoLevel)
	errWS := getFileWriter(zapcore.ErrorLevel)

	return []zapcore.Core{
		zapcore.NewCore(encoder, infoWS, zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l < zapcore.ErrorLevel && l >= config.Level
		})),
		zapcore.NewCore(encoder, errWS, zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l >= zapcore.ErrorLevel
		})),
	}
}
