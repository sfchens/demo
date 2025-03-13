package main

import (
	"demo/core/initialize"
)

func main() {
	initialize.Viper()            // 初始化配置
	initialize.ZapLog()           // 初始化日志
	initialize.Gorm()             // 初始化数据库
	initialize.Validate()         // 初始化验证器
	initialize.Cron()             // 初始化定时器
	initialize.Cache()            // 本地缓存
	initialize.RunWindowsServer() // 初始化启动服务器
}
