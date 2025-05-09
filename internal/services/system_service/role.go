package system_service

import (
	"context"
	"demo/app/request"
	"demo/global"
	"demo/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type RoleLogic struct {
}

func NewRoleLogic() *RoleLogic {
	return &RoleLogic{}
}

func (self *RoleLogic) Add(ctx context.Context, params *request.UpsertRoleReq) (err error) {
	var has int64
	global.DB.Model(&models.SysRoles{}).Where("code = ?", params.Code).Count(&has)
	if has > 0 {
		return fmt.Errorf("角色已存在！")
	}

	err = global.DB.Create(&models.SysRoles{
		Name:   params.Name,
		Code:   params.Code,
		Status: params.Status,
		Remark: params.Remark,
		Sort:   params.Sort,
	}).Error

	return
}

func (self *RoleLogic) List(ctx context.Context, params *request.RoleListReq) (resp *request.RoleListResp, err error) {
	resp = &request.RoleListResp{}

	query := global.DB.Model(&models.SysRoles{}).Order("id asc")
	if params.Name != "" {
		query.Where("name like ?", params.Name+"%")
	}
	if params.Status > 0 {
		query.Where("status = ?", params.Status)
	}
	if err = query.Count(&resp.Total).Error; err != nil {
		return
	}

	var list []*models.SysRoles
	if err = query.Scopes(models.GormPaginate(params.Page, params.PageSize)).Find(&list).Error; err != nil {
		return
	}

	items := make([]*request.RoleInfoResp, 0)
	for _, info := range list {
		item := new(request.RoleInfoResp)
		_ = copier.Copy(item, info)
		item.CreatedAt = info.CreatedAt.UnixMilli()
		items = append(items, item)
	}
	resp.Items = items

	return
}

func (self *RoleLogic) Update(ctx context.Context, id uint, params *request.UpsertRoleReq) (err error) {
	err = global.DB.Model(&models.SysRoles{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"name":   params.Name,
			"code":   params.Code,
			"sort":   params.Sort,
			"status": params.Status,
			"remark": params.Remark,
		}).Error

	return
}

func (self *RoleLogic) deleteRoleAuth(ctx context.Context, tx *gorm.DB, id uint) (err error) {
	if err = tx.Where("role_id = ?", id).Unscoped().Delete(&models.SysRoleAuths{}).Error; err != nil {
		return
	}

	err = tx.Where("role_id = ?", id).Unscoped().Delete(&models.SysRoleApis{}).Error

	return
}

func (self *RoleLogic) Delete(ctx context.Context, id uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Delete(&models.SysRoles{}, "id = ?", id).Error; err != nil {
			return err
		}

		return self.deleteRoleAuth(ctx, tx, id)
	})

	return
}

type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

func (self *RoleLogic) Assign(ctx context.Context, id uint, params *request.AssignRoleReq) (err error) {
	//var casbinInfos []CasbinInfo
	//err = global.DB.Transaction(func(tx *gorm.DB) error {
	//	if err = self.deleteRoleAuth(ctx, tx, id); err != nil {
	//		return err
	//	}
	//
	//	if err = tx.Where("v0 = ?", id).Unscoped().Delete(&gormadapter.CasbinRule{}).Error; err != nil {
	//		return err
	//	}
	//
	//	var roleAuths []models.SysRoleAuths
	//	for _, authId := range params.AuthId {
	//		roleAuths = append(roleAuths, models.SysRoleAuths{
	//			RoleId: id,
	//			AuthId: authId,
	//		})
	//	}
	//	if err = tx.Model(&models.SysRoleAuths{}).Create(&roleAuths).Error; err != nil {
	//		return err
	//	}
	//
	//	var apis []models.SysApis
	//	if err = tx.Model(&models.SysApis{}).Where("id in ?", params.ApiId).Find(&apis).Error; err != nil {
	//		return err
	//	}
	//
	//	var roleApis []models.SysRoleApis
	//	for _, api := range apis {
	//		roleApis = append(roleApis, models.SysRoleApis{
	//			RoleId: id,
	//			ApiId:  api.ID,
	//		})
	//
	//		if api.Path != "" {
	//			casbinInfos = append(casbinInfos, CasbinInfo{
	//				Path:   api.Path,
	//				Method: api.Method,
	//			})
	//		}
	//	}
	//
	//	return tx.Create(&roleApis).Error
	//})
	//
	//if err == nil {
	//	err = casbinmodule.UpsertCasbin(ctx, id, casbinInfos)
	//}

	return
}

func (self *RoleLogic) Info(ctx *gin.Context, id uint) (resp *request.RoleInfoResp, err error) {
	role := &models.SysRoles{}
	if err = global.DB.Preload("RoleAuths").Preload("RoleApis").First(&role, id).Error; err != nil {
		return
	}

	resp = &request.RoleInfoResp{}
	_ = copier.Copy(resp, role)
	for _, auth := range role.RoleAuths {
		resp.AuthId = append(resp.AuthId, auth.AuthId)
	}
	for _, api := range role.RoleApis {
		resp.ApiId = append(resp.ApiId, api.ApiId)
	}

	return resp, nil
}
