package log

import "go.uber.org/zap/zapcore"

type LogConfig struct {
	ConsoleOutput bool
	FileOutput    bool
	Level         zapcore.Level
}

func getConfig(configs ...LogConfig) LogConfig {
	if len(configs) > 0 {
		return configs[0]
	}
	return LogConfig{
		ConsoleOutput: true,
		FileOutput:    false,
		Level:         zapcore.InfoLevel,
	}
}
