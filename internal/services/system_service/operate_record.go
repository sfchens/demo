package system_service

import (
	"context"
	"demo/app/request"
	"demo/global"
	"demo/internal/models"

	"github.com/jinzhu/copier"
)

type OperateRecordLogic struct {
}

func NewOperateRecordLogic() *OperateRecordLogic {
	return &OperateRecordLogic{}
}

func (self *OperateRecordLogic) List(ctx context.Context, params *request.RecordListReq) (resp *request.RecordListResp, err error) {
	resp = &request.RecordListResp{}

	query := global.MysqlDB.Model(&models.SysOperateRecords{}).Order("id desc")
	if params.Username != "" {
		query.Where("username like ?", params.Username+"%")
	}
	if len(params.CreateTime) > 0 {
		query.Where("created_at between ? and ?", params.CreateTime[0], params.CreateTime[1])
	}
	if err = query.Count(&resp.Total).Error; err != nil {
		return
	}

	var list []*models.SysOperateRecords
	if err = query.Scopes(models.GormPaginate(params.Page, params.PageSize)).Find(&list).Error; err != nil {
		return
	}

	items := make([]*request.RecordInfoResp, 0)
	for _, info := range list {
		item := new(request.RecordInfoResp)
		_ = copier.Copy(item, info)
		item.CreatedAt = info.CreatedAt.UnixMilli()
		items = append(items, item)
	}
	resp.Items = items

	return
}

func (self *OperateRecordLogic) Delete(ctx context.Context, id int64) (err error) {
	err = global.MysqlDB.Delete(&models.SysOperateRecords{}, "id = ?", id).Error
	return
}
