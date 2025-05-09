package system_roter

import (
	"demo/app/controller/system_controller"
	"github.com/gin-gonic/gin"
)

func InitAuthRouter(routerGroup *gin.RouterGroup) {
	systemAuthGroup := routerGroup.Group("/api/auth")
	{
		systemAuthGroup.POST("/login", system_controller.AuthLogin)
		systemAuthGroup.POST("/logout", system_controller.AuthLogout)
	}
}
