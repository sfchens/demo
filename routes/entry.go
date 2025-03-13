package routes

import (
	"demo/routes/home"
	"demo/routes/public"
	"demo/routes/test"
	"github.com/gin-gonic/gin"
)

func InitPrivateRouter(PrivateGroup *gin.RouterGroup) {
	home.InitHomeRouter(PrivateGroup)
	test.InitTestRouter(PrivateGroup)
}

func InitPublicRouter(PublicGroup *gin.RouterGroup) {
	public.InitStorage(PublicGroup)
}
