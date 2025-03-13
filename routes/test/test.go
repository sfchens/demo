package test

import (
	"demo/app/controller"
	"github.com/gin-gonic/gin"
)

var testController = new(controller.Test)

func InitTestRouter(r *gin.RouterGroup) {
	r.GET("/test", testController.Test)
}
