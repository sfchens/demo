package home

import (
	"demo/app/controller"
	"github.com/gin-gonic/gin"
)

var homeController = new(controller.Home)

func InitHomeRouter(r *gin.RouterGroup) {
	r.GET("/home", homeController.Home)
}
