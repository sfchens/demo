package system_controller

import (
	"demo/app/request"
	"demo/internal/services/system_service"
	"demo/utils/response"
	"github.com/gin-gonic/gin"
)

type Casbin struct {
}

// ListRole 角色列表
func (c *Casbin) ListRole(ctx *gin.Context) {
	var (
		err error
		req request.ListRoleReq
		res request.PageResult
	)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	res, err = system_service.NewCasbinLogic().ListRole(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, res)
}

// AddOrEditRole 添加或编辑角色
func (c *Casbin) AddOrEditRole(ctx *gin.Context) {

}

// UserListRole 用户角色列表
func (c *Casbin) UserListRole(ctx *gin.Context) {

}

// AddOrEditUserRole 添加或编辑用户角色
func (c *Casbin) AddOrEditUserRole(ctx *gin.Context) {
	var (
		err error
		req request.AddOrEditRoleCasbinReq
	)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err = system_service.NewCasbinLogic().AddOrEditUserRole(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

// AddOrEditRolePermission 添加角色权限
func (c *Casbin) AddOrEditRolePermission(ctx *gin.Context) {
	var (
		err error
		req request.AddOrEditRolePermissionReq
	)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err = system_service.NewCasbinLogic().AddOrEditRolePermission(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Ok(ctx)
}
