package system_controller

import (
	"demo/app/request"
	"demo/internal/services/system_service"
	"demo/utils/helper"
	"demo/utils/response"
	"github.com/gin-gonic/gin"
)

func DictAdd(ctx *gin.Context) {
	var req request.UpsertDictReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	if err := system_service.NewDictLogic().Add(ctx, &req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func DictList(ctx *gin.Context) {
	var req request.DictListReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	resp, err := system_service.NewDictLogic().List(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, resp)
}

func DictUpdate(ctx *gin.Context) {
	id := helper.StringToInt64(ctx.Param("id"))
	if helper.IsValidNumber(id) == false {
		response.FailWithMessage(ctx, "参数错误")

		return
	}
	var req request.UpsertDictReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	if err := system_service.NewDictLogic().Update(ctx, id, &req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Ok(ctx)
}

func DictDelete(ctx *gin.Context) {
	id := helper.StringToInt64(ctx.Param("id"))
	if helper.IsValidNumber(id) == false {
		response.FailWithMessage(ctx, "参数错误")
		return
	}
	if err := system_service.NewDictLogic().Delete(ctx, id); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}
