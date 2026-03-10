package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(env string) *zap.Logger {
	var config zap.Config

	if env == "production" {
		// Production: JSON format, Info level and above, no colors
		config = zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		// Development: Console format, Debug level, colors enabled
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Build the logger
	logger, err := config.Build(zap.AddCaller())
	if err != nil {
		// If logger fails, we have no choice but to panic
		panic(err)
	}

	return logger
}
