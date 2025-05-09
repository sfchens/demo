package system_service

import (
	"context"
	"demo/app/request"
	"demo/global"
	"demo/internal/models"

	"github.com/jinzhu/copier"
)

type DictLogic struct {
}

func NewDictLogic() *DictLogic {
	return &DictLogic{}
}

func (self *DictLogic) Add(ctx context.Context, params *request.UpsertDictReq) (err error) {
	data := new(models.SysDicts)
	_ = copier.Copy(data, params)
	err = global.DB.Model(&models.SysDicts{}).Create(data).Error

	return
}

func (self *DictLogic) List(ctx context.Context, params *request.DictListReq) (resp *request.DictListResp, err error) {
	resp = &request.DictListResp{}

	query := global.DB.Model(&models.SysDicts{}).Order("id desc")
	if params.DictName != "" {
		query.Where("dict_name like ?", params.DictName+"%")
	}
	if params.DictType != "" {
		query.Where("dict_type like ?", params.DictType+"%")
	}
	if params.Status > 0 {
		query.Where("status = ?", params.Status)
	}
	if err = query.Count(&resp.Total).Error; err != nil {
		return
	}

	var list []*models.SysDicts
	if err = query.Scopes(models.GormPaginate(params.Page, params.PageSize)).Find(&list).Error; err != nil {
		return
	}

	items := make([]*request.DictInfoResp, 0)
	for _, info := range list {
		item := new(request.DictInfoResp)
		_ = copier.Copy(item, info)
		item.CreatedAt = info.CreatedAt.UnixMilli()
		items = append(items, item)
	}
	resp.Items = items

	return
}

func (self *DictLogic) Update(ctx context.Context, id int64, params *request.UpsertDictReq) (err error) {
	err = global.DB.Model(&models.SysDicts{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"dict_name":  params.DictName,
			"dict_type":  params.DictType,
			"status":     params.Status,
			"item_key":   params.ItemKey,
			"item_value": params.ItemValue,
			"sort":       params.Sort,
			"remark":     params.Remark,
		}).Error

	return
}

func (self *DictLogic) Delete(ctx context.Context, id int64) (err error) {
	err = global.DB.Delete(&models.SysDicts{}, "id = ?", id).Error

	return
}
