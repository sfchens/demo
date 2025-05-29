package system_roter

import (
	"demo/app/controller/system_controller"
	"github.com/gin-gonic/gin"
)

var menuApi = system_controller.NewMenuApi()

func InitMenuRouter(routerGroup *gin.RouterGroup) {
	systemMenuGroup := routerGroup.Group("/menu")
	{
		systemMenuGroup.GET("/router", menuApi.MenuRouter)
		systemMenuGroup.GET("/tree", menuApi.MenuTree)
		systemMenuGroup.POST("/add", menuApi.MenuAdd)
		systemMenuGroup.PUT("/update/:id", menuApi.MenuUpdate)
		systemMenuGroup.GET("/info/:id", menuApi.MenuInfo)
		systemMenuGroup.DELETE("/delete/:id", menuApi.MenuDelete)
	}
}
