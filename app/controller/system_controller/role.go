package system_controller

import (
	"demo/app/request"
	"demo/internal/services/system_service"
	"demo/utils/helper"
	"demo/utils/response"
	"github.com/gin-gonic/gin"
)

type RoleApi struct{}

func NewRoleApi() *RoleApi {
	return &RoleApi{}
}

func (c *RoleApi) RoleAdd(ctx *gin.Context) {
	var req request.UpsertRoleReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	if err := system_service.NewRoleLogic().Add(ctx, &req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func (c *RoleApi) RoleList(ctx *gin.Context) {
	var req request.RoleListReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	resp, err := system_service.NewRoleLogic().List(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

func (c *RoleApi) RoleInfo(ctx *gin.Context) {
	id := helper.StringToUint(ctx.Param("id"))
	if helper.IsValidNumber(id) == false {
		response.FailWithMessage(ctx, "参数错误")
		return
	}
	resp, err := system_service.NewRoleLogic().Info(ctx, id)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

func (c *RoleApi) RoleUpdate(ctx *gin.Context) {
	id := helper.StringToUint(ctx.Param("id"))
	if helper.IsValidNumber(id) == false {
		response.FailWithMessage(ctx, "参数错误")
		return
	}
	var req request.UpsertRoleReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	if err := system_service.NewRoleLogic().Update(ctx, id, &req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func (c *RoleApi) RoleAssign(ctx *gin.Context) {
	id := helper.StringToUint(ctx.Param("id"))
	if helper.IsValidNumber(id) == false {
		response.FailWithMessage(ctx, "参数错误")
		return
	}
	var req request.AssignRoleReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	if err := system_service.NewRoleLogic().Assign(ctx, id, &req); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func (c *RoleApi) RoleDelete(ctx *gin.Context) {
	id := helper.StringToUint(ctx.Param("id"))
	if helper.IsValidNumber(id) == false {
		response.FailWithMessage(ctx, "参数错误")
		return
	}
	if err := system_service.NewRoleLogic().Delete(ctx, id); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}
