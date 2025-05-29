package system_roter

import (
	"demo/app/controller/system_controller"
	"github.com/gin-gonic/gin"
)

func InitRoleRouter(routerGroup *gin.RouterGroup) {
	systemRoleGroup := routerGroup.Group("/role")
	roleApi := system_controller.NewRoleApi()
	{
		systemRoleGroup.POST("/add", roleApi.RoleAdd)
		systemRoleGroup.GET("/list", roleApi.RoleList)
		systemRoleGroup.GET("/info/:id", roleApi.RoleInfo)
		systemRoleGroup.PUT("/update/:id", roleApi.RoleUpdate)
		systemRoleGroup.PUT("/assign/:id", roleApi.RoleAssign)
		systemRoleGroup.DELETE("/delete/:id", roleApi.RoleDelete)
	}
}
