package system_roter

import (
	"demo/app/controller/system_controller"
	"github.com/gin-gonic/gin"
)

var (
	authApi = system_controller.NewAuthApi()
)

func InitPrivateAuthRouter(router *gin.RouterGroup) {
	systemAuthGroup := router.Group("/auth")
	{
		systemAuthGroup.GET("codes", authApi.AuthCodes)
	}
}

func InitPublicAuthRouter(router *gin.RouterGroup) {
	systemAuthGroup := router.Group("/api/auth")
	{
		systemAuthGroup.POST("/login", authApi.AuthLogin)
		systemAuthGroup.POST("/logout", authApi.AuthLogout)
	}
}
