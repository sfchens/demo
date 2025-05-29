package system_service

import (
	"context"
	"demo/app/request"
	"demo/global"
	"demo/internal/models"
	"demo/utils"
	"demo/utils/helper"
	"demo/utils/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/copier"
)

type UserLogic struct {
}

func NewUserLogic() *UserLogic {
	return &UserLogic{}
}

func (self *UserLogic) Info(ctx *gin.Context) (resp *request.UserInfoResp, err error) {
	user := models.SysUsers{}
	id := jwt.GetUserID(ctx)
	if err = global.MysqlDB.Where("id = ?", id).First(&user).Error; err != nil {
		return
	}

	resp = new(request.UserInfoResp)
	_ = copier.Copy(resp, user)
	//resp.RoleName = user.SysRole.Name
	resp.RealName = user.Username

	// 获取角色信息
	info, err := NewRoleLogic().Info(ctx, user.RoleId)
	if err != nil {
		return
	}

	// 获取权限
	var results []models.SysMenus
	if err = global.MysqlDB.Model(&models.SysMenus{}).
		Select("perm").
		Where("type = ?", "BUTTON").
		Where("id IN ?", info.AuthId).
		Find(&results).Error; err != nil {
		return
	}
	for _, item := range results {
		resp.Permissions = append(resp.Permissions, item.Perm)
	}
	return
}

func (self *UserLogic) UserByName(username string) (user models.SysUsers, err error) {
	user = models.SysUsers{}
	err = global.MysqlDB.Where("username = ?", username).First(&user).Error
	// todo 头像
	// user.Avatar = utils.TransformImageUrl(user.Avatar)

	return
}

func (self *UserLogic) Add(ctx *gin.Context, params *request.UpsertUserReq) (err error) {
	if params.Password == "" {
		params.Password = "123456"
	}

	username := params.Username
	if u, _ := self.UserByName(username); u.ID != 0 {
		return fmt.Errorf("用户已存在！")
	}

	err = global.MysqlDB.Create(&models.SysUsers{
		Username: username,
		Password: utils.BcryptHash(params.Password),
		Nickname: params.Nickname,
		Avatar:   "/uploads/default/logo.png",
		Status:   params.Status,
		Mobile:   params.Mobile,
		Email:    params.Email,
		Remark:   params.Remark,
		RoleId:   params.RoleId,
		CreateBy: jwt.GetUserInfo(ctx).Username,
	}).Error

	return
}

func (self *UserLogic) List(ctx context.Context, params *request.UserListReq) (resp *request.UserListResp, err error) {
	resp = &request.UserListResp{}

	query := global.MysqlDB.Model(&models.SysUsers{}).Order("id desc")
	if params.Username != "" {
		query.Where("username like ?", params.Username+"%")
	}
	if params.Status > 0 {
		query.Where("status = ?", params.Status)
	}
	if err = query.Count(&resp.Total).Error; err != nil {
		return
	}

	var list []*models.SysUsers
	if err = query.Scopes(models.GormPaginate(params.Page, params.PageSize)).Find(&list).Error; err != nil {
		return
	}

	items := make([]*request.UserInfoResp, 0)
	for _, info := range list {
		item := new(request.UserInfoResp)
		_ = copier.Copy(item, info)
		//item.RoleName = info.SysRole.Name
		item.CreatedAt = info.CreatedAt.UnixMilli()
		items = append(items, item)
	}
	resp.Items = items

	return
}

func (self *UserLogic) Update(ctx context.Context, id int64, params *request.UpsertUserReq) (err error) {
	user := models.SysUsers{}
	if err = global.MysqlDB.First(&user, id).Error; err != nil {
		return err
	}

	if user.Username != params.Username {
		var count int64
		global.MysqlDB.Model(&models.SysUsers{}).Where("username = ?", params.Username).Count(&count)
		if count >= 1 {
			return errors.New("用户已存在")
		}
	}

	uMap := map[string]interface{}{
		"username": params.Username,
		"nickname": params.Nickname,
		"role_id":  params.RoleId,
		"status":   params.Status,
		"mobile":   params.Mobile,
		"email":    params.Email,
		"remark":   params.Remark,
		//"update_by": xauth.GetTokenData[int64](ctx, "id"),
	}
	if len(params.Password) > 0 {
		newSalt := helper.RandString(6)
		uMap["salt"] = newSalt
		uMap["password"] = utils.BcryptHash(params.Password)
	}

	err = global.MysqlDB.Model(&models.SysUsers{}).Where("id = ?", id).Updates(uMap).Error

	return
}

func (self *UserLogic) Delete(ctx context.Context, id int64) (err error) {
	err = global.MysqlDB.Delete(&models.SysUsers{}, "id = ?", id).Error
	return
}

func (self *UserLogic) GetById(id uint) (user models.SysUsers, err error) {
	err = global.MysqlDB.Where("id = ?", id).Limit(1).Find(&user).Error
	return
}
