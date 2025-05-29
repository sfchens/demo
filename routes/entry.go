package routes

import (
	"demo/routes/home"
	"demo/routes/system_roter"
	"demo/routes/test"
	"github.com/gin-gonic/gin"
)

func InitRouter(PrivateGroup *gin.RouterGroup, PublicGroup *gin.RouterGroup) {
	// 私有路由
	home.InitHomeRouter(PrivateGroup)
	test.InitTestRouter(PrivateGroup)
	system_roter.InitSystemRouter(PrivateGroup) // 加载系统路由

	// 公共路由
	system_roter.InitPublicRouter(PublicGroup)
}
