package middlewares

import (
	"demo/internal/services/system_service"
	"demo/utils/jwt"
	"demo/utils/response"
	"github.com/gin-gonic/gin"
	"strings"
)

var (
	jwtService  = system_service.NewJwtLogic()
	userService = system_service.NewUserLogic()
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := GetToken(ctx)
		if token == "" {
			response.UnauthorizedWithDetailed(ctx, gin.H{"reload": true}, "token required")
			ctx.Abort()
			return
		}

		if jwtService.IsBlackList(token) {
			response.UnauthorizedWithDetailed(ctx, gin.H{"reload": true}, "invalid token")
			ctx.Abort()
			return
		}

		jwtObj := jwt.NewJWT()
		claims, err := jwtObj.ParseToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				response.UnauthorizedWithDetailed(ctx, gin.H{"reload": true}, "token expired")
				ctx.Abort()
				return
			}

			response.UnauthorizedWithDetailed(ctx, gin.H{"reload": true}, err.Error())
			ctx.Abort()
			return
		}

		if claims.ID == 0 {
			response.UnauthorizedWithDetailed(ctx, gin.H{"reload": true}, "user not found")
			ctx.Abort()
			return
		}

		user, _ := userService.GetById(claims.ID)
		if user.ID == 0 || user.Status != 1 {
			response.UnauthorizedWithDetailed(ctx, gin.H{"reload": true}, "user not found or disabled")
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}

func GetToken(ctx *gin.Context) (token string) {
	token = ctx.Request.Header.Get("x-token")
	if token == "" {
		tokenTmpArr := strings.Split(ctx.Request.Header.Get("authorization"), " ")
		if len(tokenTmpArr) == 2 {
			token = tokenTmpArr[1]
		}
	}
	return token
}
