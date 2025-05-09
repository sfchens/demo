package system_roter

import (
	"demo/app/controller/system_controller"
	"github.com/gin-gonic/gin"
)

func InitOperateRecordRouter(routerGroup *gin.RouterGroup) {
	systemRecordGroup := routerGroup.Group("/record")
	{
		systemRecordGroup.GET("/list", system_controller.RecordList)
		systemRecordGroup.DELETE("/delete/:id", system_controller.RecordDelete)
	}
}
