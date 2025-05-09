package system_service

import (
	"context"
	"demo/app/request"
	"demo/global"
	"demo/internal/models"
	"fmt"

	"github.com/jinzhu/copier"
)

var apiInfoMap = make(map[string]*models.SysApis)
var recordUseApiInfoMap = make(map[string]*models.SysApis)

type ApiLogic struct {
}

func NewApiLogic() *ApiLogic {
	return &ApiLogic{}
}

func (self *ApiLogic) Add(ctx context.Context, params *request.UpsertApiReq) (err error) {
	var has int64
	global.DB.Model(&models.SysApis{}).Where("parent_id != ?", 0).Where("path = ?", params.Path).Count(&has)
	if has > 0 {
		return fmt.Errorf("API已存在！")
	}

	err = global.DB.Create(&models.SysApis{
		ParentId:    params.ParentId,
		Description: params.Description,
		Method:      params.Method,
		Path:        params.Path,
	}).Error

	return
}

func (self *ApiLogic) List(ctx context.Context, params *request.ApiListReq) (resp *request.ApiListResp, err error) {
	resp = &request.ApiListResp{}
	query := global.DB.Model(&models.SysApis{})
	if params.Description != "" {
		query.Where("description like ?", params.Description+"%")
	}
	if params.Path != "" {
		query.Where("path = ?", params.Path)
	}
	if params.OnlyParent {
		query.Where("parent_id = ?", 0)
	}

	var list []*models.SysApis
	if err = query.Find(&list).Error; err != nil {
		return
	}

	items := make([]*request.ApiInfoResp, 0)
	for _, info := range list {
		item := new(request.ApiInfoResp)
		_ = copier.Copy(item, info)
		item.CreatedAt = info.CreatedAt.UnixMilli()
		items = append(items, item)
	}
	if params.OnlyParent {
		items = append(items, &request.ApiInfoResp{
			Id:          0,
			CreatedAt:   0,
			ParentId:    0,
			Description: "根API",
			Method:      "",
			Path:        "",
		})
	}
	resp.Items = items
	resp.Total = int64(len(items))

	return
}

func (self *ApiLogic) Update(ctx context.Context, id int64, params *request.UpsertApiReq) (err error) {
	var has int64
	global.DB.Model(&models.SysApis{}).Where("parent_id != ?", 0).Where("path = ?", params.Path).Count(&has)
	if has > 0 {
		return fmt.Errorf("API已存在！")
	}

	err = global.DB.Model(&models.SysApis{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"parent_id":   params.ParentId,
			"description": params.Description,
			"method":      params.Method,
			"path":        params.Path,
		}).Error

	return
}

func (self *ApiLogic) Delete(ctx context.Context, id int64) (err error) {
	var count int64
	global.DB.Where("api_id", id).Model(&models.SysRoleApis{}).Count(&count)
	if count > 0 {
		return fmt.Errorf("请先删除角色API权限后再操作")
	}

	err = global.DB.Delete(&models.SysApis{}, "id = ?", id).Error

	return
}

func (self *ApiLogic) CacheApiInfo() {
	var list []*models.SysApis
	if err := global.DB.Model(&models.SysApis{}).Find(&list).Error; err != nil {
		//gina.Log.Error("[CacheApiInfo]获取API信息失败！", zap.Error(err))
		return
	}

	for _, item := range list {
		apiInfoMap[item.Path] = item
		recordUseApiInfoMap[fmt.Sprintf("%s_%s", item.Path, item.Method)] = item
	}
}

func (self *ApiLogic) GetRecordDescription(path, method string) string {
	info, ok := recordUseApiInfoMap[fmt.Sprintf("%s_%s", path, method)]
	if !ok {
		return ""
	}

	return info.Description
}
