package global

import "go.uber.org/zap"

var (
	shutdownFunc = make([]func(), 0) // 在关闭时执行的函数
)

// RegisterShutdownFunc 注册在关闭时执行的函数
func RegisterShutdownFunc(f func()) {
	shutdownFunc = append(shutdownFunc, f)
}

func GetZapLog(name ...string) *zap.Logger {
	if len(name) == 0 {
		return Logger[DefaultLoggerKey]
	}
	if Logger[name[0]] == nil {
		return Logger[DefaultLoggerKey]
	}
	return Logger[name[0]]
}
