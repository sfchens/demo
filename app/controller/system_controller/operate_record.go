package system_controller

import (
	"demo/app/request"
	"demo/internal/services/system_service"
	"demo/utils/helper"
	"demo/utils/response"
	"github.com/gin-gonic/gin"
)

func RecordList(ctx *gin.Context) {
	var req request.RecordListReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	resp, err := system_service.NewOperateRecordLogic().List(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

func RecordDelete(ctx *gin.Context) {
	id := helper.StringToInt64(ctx.Param("id"))
	if helper.IsValidNumber(id) == false {
		response.FailWithMessage(ctx, "参数错误")
		return
	}
	if err := system_service.NewOperateRecordLogic().Delete(ctx, id); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Ok(ctx)
}
