package system_roter

import (
	"demo/app/controller/system_controller"
	"github.com/gin-gonic/gin"
)

func InitRoleRouter(routerGroup *gin.RouterGroup) {
	systemRoleGroup := routerGroup.Group("/role")
	{
		systemRoleGroup.POST("/add", system_controller.RoleAdd)
		systemRoleGroup.GET("/list", system_controller.RoleList)
		systemRoleGroup.GET("/info/:id", system_controller.RoleInfo)
		systemRoleGroup.PUT("/update/:id", system_controller.RoleUpdate)
		systemRoleGroup.PUT("/assign/:id", system_controller.RoleAssign)
		systemRoleGroup.DELETE("/delete/:id", system_controller.RoleDelete)
	}
}
