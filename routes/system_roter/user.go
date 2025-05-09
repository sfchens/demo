package system_roter

import (
	"demo/app/controller/system_controller"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(routerGroup *gin.RouterGroup) {
	systemUserGroup := routerGroup.Group("/user")
	{
		systemUserGroup.GET("/info", system_controller.UserInfo)
		systemUserGroup.POST("/add", system_controller.UserAdd)
		systemUserGroup.GET("/list", system_controller.UserList)
		systemUserGroup.PUT("/update/:id", system_controller.UserUpdate)
		systemUserGroup.DELETE("/delete/:id", system_controller.UserDelete)
	}
}
