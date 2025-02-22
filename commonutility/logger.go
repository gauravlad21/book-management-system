package commonutility

import (
	"log"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

var logger *Logger

func GetLogger() *Logger {
	if logger == nil {
		logger = initializeLogger()
	}
	return logger
}

func initializeLogger() *Logger {
	var config zap.Config
	config = zapdriver.NewDevelopmentConfig()
	config.Encoding = "console"
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	clientLogger, err := config.Build()
	if err != nil {
		log.Fatalf("zap.config.Build(): %v", err)
	}
	return &Logger{clientLogger}
}
