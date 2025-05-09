package system_service

import (
	"context"
	"demo/app/request"
	"demo/global"
	"demo/internal/models"
	"fmt"
	"github.com/jinzhu/copier"
)

type MenuLogic struct {
}

func NewMenuLogic() *MenuLogic {
	return &MenuLogic{}
}

func (self *MenuLogic) Router(ctx context.Context) (resp *request.MenuResp, err error) {
	var authIds []int64
	global.DB.Table("sys_users").
		Select("sys_role_auths.auth_id").
		Joins("JOIN sys_role_auths ON sys_users.role_id = sys_role_auths.role_id").
		//Where("sys_users.id = ?", xauth.GetTokenData[int64](ctx, "id")).
		Where("sys_users.id = ?", 1).
		Pluck("sys_role_auths.auth_id", &authIds)
	if len(authIds) == 0 {
		return
	}

	var list []*models.SysMenus
	if err = global.DB.Model(&models.SysMenus{}).
		Where("status = ?", 1).
		Where("id IN ?", authIds).
		Where("type != ?", "BUTTON").
		Order("id asc").
		Find(&list).Error; err != nil {
		return
	}

	items := self.getMenuRouter(list, 0)
	resp = &request.MenuResp{}
	resp.Items = items
	resp.Total = int64(len(items))
	return
}

// 根据指定的pid获取菜单路由树
func (self *MenuLogic) getMenuRouter(menuList []*models.SysMenus, pid uint) (treeList []request.RouterResp) {
	for _, v := range menuList {
		if v.ParentId == pid {
			child := self.getMenuRouter(menuList, v.ID)
			node := request.RouterResp{
				Path:      v.Path,
				Component: v.Component,
				Name:      v.RouteName,
				Meta: request.RouterMetaResp{
					Icon:     v.Icon,
					Sort:     v.Sort,
					Title:    v.Name,
					AffixTab: v.AffixTab == 1,
				},
			}
			node.Children = child
			treeList = append(treeList, node)
		}
	}

	return treeList
}

func (self *MenuLogic) Tree(ctx context.Context, params *request.MenuTreeReq) (resp *request.MenuResp, err error) {
	query := global.DB.Model(&models.SysMenus{})
	if params.Name != "" {
		query.Where("name like ?", params.Name+"%")
	}
	if params.Status > 0 {
		query.Where("status = ?", params.Status)
	}
	var list []*models.SysMenus
	if err = query.Find(&list).Error; err != nil {
		return
	}

	var items []*request.MenuInfo
	for _, info := range list {
		item := new(request.MenuInfo)
		_ = copier.Copy(item, info)
		_ = copier.Copy(&item.Meta, info)
		items = append(items, item)
	}

	resp = &request.MenuResp{}
	resp.Items = items
	resp.Total = int64(len(items))

	return
}

func (self *MenuLogic) Add(ctx context.Context, params *request.MenuInfo) (err error) {
	data := new(models.SysMenus)
	_ = copier.Copy(data, params)
	_ = copier.Copy(data, params.Meta)
	err = global.DB.Model(&models.SysMenus{}).Create(data).Error

	return
}

func (self *MenuLogic) Update(ctx context.Context, id int64, params *request.MenuInfo) (err error) {
	err = global.DB.Model(&models.SysMenus{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"parent_id":             params.ParentId,
			"path":                  params.Path,
			"status":                params.Status,
			"type":                  params.Type,
			"route_name":            params.RouteName,
			"perm":                  params.Perm,
			"component":             params.Component,
			"sort":                  params.Meta.Sort,
			"icon":                  params.Meta.Icon,
			"name":                  params.Meta.Name,
			"affix_tab":             params.Meta.AffixTab,
			"hide_children_in_menu": params.Meta.HideChildrenInMenu,
			"hide_in_breadcrumb":    params.Meta.HideInBreadcrumb,
			"hide_in_menu":          params.Meta.HideInMenu,
			"hide_in_tab":           params.Meta.HideInTab,
			"keep_alive":            params.Meta.KeepAlive,
		}).Error

	return
}

func (self *MenuLogic) Info(ctx context.Context, id int64) (resp *request.MenuInfo, err error) {
	// TODO implement

	return resp, nil
}

func (self *MenuLogic) Delete(ctx context.Context, id int64) (err error) {
	var count int64
	global.DB.Model(&models.SysRoleAuths{}).Where("auth_id", id).Count(&count)
	if count > 0 {
		return fmt.Errorf("请删除角色菜单权限！")
	}
	err = global.DB.Delete(&models.SysMenus{}, "id = ?", id).Error
	if err != nil {
		return
	}
	return
}
