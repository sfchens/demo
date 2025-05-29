package system_roter

import "github.com/gin-gonic/gin"

func InitSystemRouter(PrivateGroup *gin.RouterGroup) {
	InitApiRouter(PrivateGroup)
	InitDictRouter(PrivateGroup)
	InitOperateRecordRouter(PrivateGroup)
	InitRoleRouter(PrivateGroup)
	InitMenuRouter(PrivateGroup)
	InitUserRouter(PrivateGroup)

	InitPrivateAuthRouter(PrivateGroup)

}

func InitPublicRouter(PublicGroup *gin.RouterGroup) {
	InitPublicAuthRouter(PublicGroup)
}
