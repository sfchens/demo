package system_roter

import (
	"demo/app/controller/system_controller"
	"github.com/gin-gonic/gin"
)

var apiApi = system_controller.NewApiApi()

func InitApiRouter(routerGroup *gin.RouterGroup) {
	systemApiGroup := routerGroup.Group("/api")
	{
		systemApiGroup.POST("/add", apiApi.ApiAdd)
		systemApiGroup.GET("/list", apiApi.ApiList)
		systemApiGroup.PUT("/update/:id", apiApi.ApiUpdate)
		systemApiGroup.DELETE("/delete/:id", apiApi.ApiDelete)
	}
}
