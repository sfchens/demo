package initialize

import (
	"demo/global"
	"demo/internal/services/system_service"
)

func Middleware() {
	global.Casbin = system_service.NewCasbinLogic().Casbin()

}
