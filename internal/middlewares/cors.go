package middlewares

import (
	"demo/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")

		if global.ConfigAll.System.UseCors {
			context.Header("Access-Control-Allow-Origin", origin)
			context.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token, X-Token, X-User-Id, Sso-Request-Id")
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			context.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
