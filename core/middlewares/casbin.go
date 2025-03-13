package middlewares

import (
	"demo/core/services/system"
	"demo/global"
	"demo/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CasbinAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		//tacticInUse, _ := jwt.GetClaims(context)
		path := context.Request.URL.Path
		method := context.Request.Method
		//auth := strconv.Itoa(int(tacticInUse.AuthorityId))
		e := system.NewCasbin().Casbin()
		success, _ := e.Enforce("amie", path, method)
		fmt.Println("success: ", success)
		if global.ConfigAll.System.Env == "develop" || success {
			context.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "权限不足", context)
			context.Abort()
			return
		}
	}
}
