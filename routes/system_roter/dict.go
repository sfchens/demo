package system_roter

import (
	"demo/app/controller/system_controller"
	"github.com/gin-gonic/gin"
)

var dictApi = system_controller.NewDictApi()

func InitDictRouter(routerGroup *gin.RouterGroup) {
	systemDictGroup := routerGroup.Group("/dict")
	{
		systemDictGroup.POST("/add", dictApi.DictAdd)
		systemDictGroup.GET("/list", dictApi.DictList)
		systemDictGroup.PUT("/update/:id", dictApi.DictUpdate)
		systemDictGroup.DELETE("/delete/:id", dictApi.DictDelete)
	}
}
