package system_roter

import (
	"demo/app/controller/system_controller"
	"github.com/gin-gonic/gin"
)

func InitDictRouter(routerGroup *gin.RouterGroup) {
	systemDictGroup := routerGroup.Group("/dict")
	{
		systemDictGroup.POST("/add", system_controller.DictAdd)
		systemDictGroup.GET("/list", system_controller.DictList)
		systemDictGroup.PUT("/update/:id", system_controller.DictUpdate)
		systemDictGroup.DELETE("/delete/:id", system_controller.DictDelete)
	}
}
