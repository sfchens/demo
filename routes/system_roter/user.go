package system_roter

import (
	"demo/app/controller/system_controller"
	"github.com/gin-gonic/gin"
)

var userApi = system_controller.NewUserApi()

func InitUserRouter(routerGroup *gin.RouterGroup) {
	systemUserGroup := routerGroup.Group("/user")
	{
		systemUserGroup.GET("/info", userApi.UserInfo)
		systemUserGroup.POST("/add", userApi.UserAdd)
		systemUserGroup.GET("/list", userApi.UserList)
		systemUserGroup.PUT("/update/:id", userApi.UserUpdate)
		systemUserGroup.DELETE("/delete/:id", userApi.UserDelete)
	}
}
