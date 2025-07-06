package main

import (
	"github.com/huajianxiaowanzi/zapX/log"
	"go.uber.org/zap"
)

func main() {
	//log.Init(log.LogConfig{
	//	ConsoleOutput: true,
	//	FileOutput:    true,
	//	Level:         zapcore.InfoLevel,
	//})
	log.Init()

	log.Info("SetupLogger")
	log.Errorf("error level test: %s", "111")
	log.With(zap.String("Trace", "12345677")).Info("this is a log")

}
