package system_roter

import "github.com/gin-gonic/gin"

func InitSystemRouter(Router *gin.RouterGroup) {
	InitApiRouter(Router)
	InitDictRouter(Router)
	InitOperateRecordRouter(Router)
	InitRoleRouter(Router)
	InitMenuRouter(Router)
	InitUserRouter(Router)
}

func InitPublicRouter(Router *gin.RouterGroup) {
	InitAuthRouter(Router)
}
