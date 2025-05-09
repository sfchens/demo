package initialize

import (
	global2 "demo/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper(path ...string) *viper.Viper {
	var configPath string

	if len(path) == 0 {
		flag.StringVar(&configPath, "c", "", "choose config file.")
		flag.Parse()
		if configPath == "" { // 判断命令行参数是否为空
			switch gin.Mode() {
			case gin.DebugMode:
				configPath = global2.ConfigDefaultPath
				fmt.Printf("您正在使用%s环境名称,config的路径为%s\n", gin.EnvGinMode, global2.ConfigDefaultPath)
			case gin.ReleaseMode:
				configPath = global2.ConfigReleasePath
				fmt.Printf("您正在使用%s环境名称,config的路径为%s\n", gin.EnvGinMode, global2.ConfigReleasePath)
			case gin.TestMode:
				configPath = global2.ConfigUatPath
				fmt.Printf("您正在使用%s环境名称,config的路径为%s\n", gin.EnvGinMode, global2.ConfigUatPath)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", configPath)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		configPath = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", configPath)
	}

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global2.ConfigAll); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global2.ConfigAll); err != nil {
		fmt.Println(err)
	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	//global.INTRA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	//global.BlackCache = local_cache.NewCache(
	//	local_cache.SetDefaultExpire(time.Second * time.Duration(global.INTRA_CONFIG.JWT.ExpiresTime)),
	//)
	return v
}
