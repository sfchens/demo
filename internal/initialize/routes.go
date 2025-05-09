package initialize

import (
	"demo/global"
	"demo/internal/middlewares"
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
	r.Use(middlewares.Cors()) // 直接放行全部跨域请求

	// 公共API
	PublicGroup := r.Group("")
	routes.InitPublicRouter(PublicGroup)
	// 私有API
	PrivateGroup := r.Group("/api")
	PrivateGroup.Use(middlewares.JwtAuth())
	PrivateGroup.Use(middlewares.OperateRecord())
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
