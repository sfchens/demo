package system_controller

import (
	"demo/app/request"
	"demo/internal/services/system_service"
	"demo/utils/helper"
	"demo/utils/response"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func NewUserApi() *UserApi {
	return &UserApi{}
}
func (c *UserApi) UserInfo(ctx *gin.Context) {
	resp, err := system_service.NewUserLogic().Info(ctx)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

func (c *UserApi) UserAdd(ctx *gin.Context) {
	var req request.UpsertUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	if err := system_service.NewUserLogic().Add(ctx, &req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Ok(ctx)
}

func (c *UserApi) UserList(ctx *gin.Context) {
	var req request.UserListReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	resp, err := system_service.NewUserLogic().List(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

func (c *UserApi) UserUpdate(ctx *gin.Context) {
	id := helper.StringToInt64(ctx.Param("id"))
	if helper.IsValidNumber(id) == false {
		response.FailWithMessage(ctx, "参数错误")
		return
	}
	var req request.UpsertUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	if err := system_service.NewUserLogic().Update(ctx, id, &req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Ok(ctx)
}

func (c *UserApi) UserDelete(ctx *gin.Context) {
	id := helper.StringToInt64(ctx.Param("id"))
	if helper.IsValidNumber(id) == false {
		response.FailWithMessage(ctx, "参数错误")
		return
	}
	if err := system_service.NewUserLogic().Delete(ctx, id); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}
