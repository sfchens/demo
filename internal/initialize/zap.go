package initialize

import (
	"demo/global"
	initLogger "demo/internal/initialize/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ZapLog(fileNames ...string) *zap.Logger {
	var (
		fileName string
	)
	if len(fileNames) > 0 {
		fileName = fileNames[0]
	}

	if global.Logger[fileName] != nil {
		return global.Logger[fileName]
	}

	if len(fileNames) == 0 {
		fileName = global.DefaultLoggerKey
		if global.Logger[fileName] != nil {
			return global.Logger[fileName]
		}
		fileName = global.DefaultLoggerKey
	}

	customZap := initLogger.NewZap()
	customZap.FileName = fileName
	customZap.Level = global.ConfigAll.Zap.Level
	customZap.Format = "console"

	cores := customZap.GetZapCores()
	loggerT := zap.New(zapcore.NewTee(cores...))
	if global.ConfigAll.Zap.ShowLine { // 显示文件路径和行号
		loggerT = loggerT.WithOptions(zap.AddCaller())
	}

	global.Logger[fileName] = loggerT
	return loggerT
}
