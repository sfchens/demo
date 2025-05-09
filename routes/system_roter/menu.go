package system_roter

import (
	"demo/app/controller/system_controller"
	"github.com/gin-gonic/gin"
)

func InitMenuRouter(routerGroup *gin.RouterGroup) {
	systemMenuGroup := routerGroup.Group("/menu")
	{
		systemMenuGroup.GET("/router", system_controller.MenuRouter)
		systemMenuGroup.GET("/tree", system_controller.MenuTree)
		systemMenuGroup.POST("/add", system_controller.MenuAdd)
		systemMenuGroup.PUT("/update/:id", system_controller.MenuUpdate)
		systemMenuGroup.GET("/info/:id", system_controller.MenuInfo)
		systemMenuGroup.DELETE("/delete/:id", system_controller.MenuDelete)
	}
}
