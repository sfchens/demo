package routes

import (
	"demo/routes/home"
	"demo/routes/public"
	"demo/routes/system_roter"
	"demo/routes/test"
	"github.com/gin-gonic/gin"
)

func InitPrivateRouter(PrivateGroup *gin.RouterGroup) {
	home.InitHomeRouter(PrivateGroup)
	test.InitTestRouter(PrivateGroup)
	system_roter.InitSystemRouter(PrivateGroup) // 加载系统路由
}

func InitPublicRouter(PublicGroup *gin.RouterGroup) {
	public.InitStorage(PublicGroup)
	system_roter.InitAuthRouter(PublicGroup)
}
