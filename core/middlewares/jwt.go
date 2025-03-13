package middlewares

import (
	"demo/core/services/system"
	"demo/utils/jwt"
	"demo/utils/response"
	"github.com/gin-gonic/gin"
)

var (
	jwtService  = system.NewJwtService()
	userService = system.NewUserService()
)

func JwtAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := context.Request.Header.Get("x-token")
		if token == "" {
			response.UnauthorizedWithDetailed(gin.H{"reload": true}, "token required", context)
			context.Abort()
			return
		}

		if jwtService.IsBlackList(token) {
			response.UnauthorizedWithDetailed(gin.H{"reload": true}, "invalid token", context)
			context.Abort()
			return
		}

		jwtObj := jwt.NewJWT()
		claims, err := jwtObj.ParseToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				response.UnauthorizedWithDetailed(gin.H{"reload": true}, "token expired", context)
				context.Abort()
				return
			}

			response.UnauthorizedWithDetailed(gin.H{"reload": true}, err.Error(), context)
			context.Abort()
			return
		}

		if claims.ID == 0 {
			response.UnauthorizedWithDetailed(gin.H{"reload": true}, "user not found", context)
			context.Abort()
			return
		}

		user, _ := userService.GetById(claims.ID)

		if user.ID == 0 || user.Enable != 1 {
			response.UnauthorizedWithDetailed(gin.H{"reload": true}, "user not found or disabled", context)
			context.Abort()
			return
		}

		context.Set("claims", claims)
		context.Next()
	}
}
