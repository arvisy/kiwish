package config

import (
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Newlogger(service string, level zapcore.Level) *zap.Logger {
	esEncoderCfg := ecszap.NewDefaultEncoderConfig()
	escore := ecszap.NewCore(esEncoderCfg, zapcore.AddSync(os.Stdout), level)

	core := zapcore.NewTee(escore)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel), zap.Fields(zap.String("service", service)))
	defer logger.Sync()

	return logger
}
