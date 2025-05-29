package system_roter

import (
	"demo/app/controller/system_controller"
	"github.com/gin-gonic/gin"
)

var recordApi = system_controller.NewRecordApi()

func InitOperateRecordRouter(routerGroup *gin.RouterGroup) {
	systemRecordGroup := routerGroup.Group("/record")
	{
		systemRecordGroup.GET("/list", recordApi.RecordList)
		systemRecordGroup.DELETE("/delete/:id", recordApi.RecordDelete)
	}
}
