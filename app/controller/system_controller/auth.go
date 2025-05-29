package system_controller

import (
	"demo/app/request"
	"demo/internal/services/system_service"
	"demo/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthApi struct{}

func NewAuthApi() *AuthApi {
	return &AuthApi{}
}
func (c *AuthApi) AuthLogin(ctx *gin.Context) {
	var (
		req request.LoginReq
		res request.LoginResp

		err error
	)
	if err = ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	fmt.Printf("req:  %+v\n", req)
	res, err = system_service.NewAuthLogic().Login(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, res)
}

func (c *AuthApi) AuthLogout(ctx *gin.Context) {
	var req request.LogoutReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	if err := system_service.NewAuthLogic().Logout(ctx, &req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Ok(ctx)
}

func (c *AuthApi) AuthCodes(ctx *gin.Context) {

	resp, err := system_service.NewAuthLogic().Codes(ctx)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, resp)
}
