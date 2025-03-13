package initialize

import (
	"demo/core/middlewares"
	"demo/global"
	"demo/routes"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func Routers() *gin.Engine {
	r := gin.Default()
	// 公共API
	PublicGroup := r.Group("")
	routes.InitPublicRouter(PublicGroup)

	// 私有API
	PrivateGroup := r.Group("")
	PrivateGroup.Use(middlewares.JwtAuth()).Use(middlewares.CasbinAuth())
	routes.InitPrivateRouter(PrivateGroup)
	return r
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func RunWindowsServer() {
	Router := Routers()
	address := fmt.Sprintf(":%d", global.ConfigAll.System.Port)
	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	global.GetZapLog().Info("server run success on ", zap.String("address", address))
	global.GetZapLog().Error(s.ListenAndServe().Error())
}
