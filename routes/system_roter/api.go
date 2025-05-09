package system_roter

import (
	"demo/app/controller/system_controller"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(routerGroup *gin.RouterGroup) {
	systemApiGroup := routerGroup.Group("/api")
	{
		systemApiGroup.POST("/add", system_controller.ApiAdd)
		systemApiGroup.GET("/list", system_controller.ApiList)
		systemApiGroup.PUT("/update/:id", system_controller.ApiUpdate)
		systemApiGroup.DELETE("/delete/:id", system_controller.ApiDelete)
	}
}
