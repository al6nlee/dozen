package initialize

import (
	"dozen/backend/internal/global"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
	var (
		logger *zap.Logger
		err    error
	)

	if global.ServerConfig.Mode == "prod" {
		cfg := zap.NewProductionConfig()

		// 时间格式（推荐）
		cfg.EncoderConfig.TimeKey = "ts"
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		// 日志级别（可配置化）
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

		logger, err = cfg.Build(zap.AddCaller())
	} else {
		cfg := zap.NewDevelopmentConfig()
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

		logger, err = cfg.Build(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	}

	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(logger)
}
