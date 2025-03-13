package internal

import (
	"demo/global"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

func (r *fileRotatelogs) GetWriteSyncer(filename, level string) (zapcore.WriteSyncer, error) {
	var pathLog = global.ConfigAll.Zap.Director
	var filNameT = level + ".log"

	if filename != "" {
		pathLog = path.Join(pathLog, "%Y-%m-%d", filename, filNameT)
	} else {
		pathLog = path.Join(pathLog, "%Y-%m-%d", filNameT)
	}
	fileWriter, err := rotatelogs.New(
		pathLog,
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(global.ConfigAll.Zap.MaxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if global.ConfigAll.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
