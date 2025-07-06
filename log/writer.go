package log

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

func getFileWriter(level zapcore.Level) zapcore.WriteSyncer {
	now := time.Now().Format("2006-01-02")
	_ = os.MkdirAll(filepath.Join("logs", now), 0755)

	var filename string
	if level >= zapcore.ErrorLevel {
		filename = fmt.Sprintf("logs/%s/err.log", now)
	} else {
		filename = fmt.Sprintf("logs/%s/out.log", now)
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("failed to open log file: %v", err))
	}
	return zapcore.AddSync(file)
}
