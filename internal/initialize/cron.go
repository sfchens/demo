package initialize

import (
	"demo/global"
	"fmt"
	"github.com/robfig/cron/v3"
	"sync"
)

var m sync.Mutex

func Cron() {
	global.Cron = cron.New()
	fmt.Println("开启定时器")

	global.Cron.AddFunc("@hourly", func() {
		fmt.Println("Cron Every hour")
	})
	global.Cron.AddFunc("@daily", func() {
		fmt.Println("Cron Every Day")
	})

	global.Cron.AddFunc("@every 0h10m0s", func() {
		fmt.Println("Cron Every 10 Min")
	})
	global.Cron.AddFunc("* * * * *", func() {
		fmt.Println("Cron Linux * * * * *")
	})
}
