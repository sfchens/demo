package system_service

import (
	"context"
	"demo/app/request"
	"demo/global"
	"demo/internal/models"
	"demo/utils/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthLogic struct {
}

func NewAuthLogic() *AuthLogic {
	return &AuthLogic{}
}

func (self *AuthLogic) Login(ctx context.Context, params *request.LoginReq) (resp request.LoginResp, err error) {
	var (
		user = new(models.SysUsers)
	)

	if err = global.MysqlDB.Model(&models.SysUsers{}).Where("username = ?", params.Username).First(user).Error; err != nil {
		return
	}
	if user.Status != models.SysUserStatusNormal {
		err = errors.New("用户已被禁用")
		return
	}
	fmt.Printf("user:  %+v\n", user)

	//if ok := utils.BcryptCheck(params.Password, user.Password); !ok {
	//	err = errors.New("密码错误")
	//	return
	//}

	var token string
	j := &jwt.JWT{SigningKey: []byte(global.ConfigAll.Jwt.SigningKey)} // 唯一签名
	token, err = j.CreateToken(request.CustomClaims{
		BaseClaims: request.BaseClaims{
			ID:       user.ID,
			NickName: user.Nickname,
			Username: user.Username,
			RoleId:   user.RoleId,
			Email:    user.Email,
		},
	})
	if err != nil {
		return
	}

	resp = request.LoginResp{
		AccessToken: token,
		Id:          user.ID,
		Password:    "",
		RealName:    user.Username,
		Roles:       []string{""},
		Username:    user.Username,
	}
	fmt.Printf("resp:  %+v\n", resp)

	return
}

func (self *AuthLogic) Logout(ctx context.Context, params *request.LogoutReq) (err error) {
	// TODO implement

	return nil
}

func (self *AuthLogic) Codes(ctx *gin.Context) (resp []string, err error) {
	userID := jwt.GetUserID(ctx)
	if userID <= 0 {
		err = errors.New("账号异常")
		return
	}
	var sysUser models.SysUsers
	err = global.MysqlDB.Where("id =?", userID).Find(&sysUser).Error
	if err != nil {
		return
	}

	// 取出拥有的角色
	var sysUserRolesList []models.SysUserRoles
	err = global.MysqlDB.Model(&models.SysUserRoles{}).Where("id=?", userID).Find(&sysUserRolesList).Error
	if err != nil {
		return
	}

	var RoleIds []uint
	var isSuperAdmin bool
	for _, userRole := range sysUserRolesList {
		RoleIds = append(RoleIds, userRole.RoleId)
		if userRole.RoleId == 1 {
			isSuperAdmin = true
		}
	}
	var menuIds []uint
	if !isSuperAdmin {
		// 取出角色拥有的菜单
		var menuAuthList []models.SysRoleAuths
		err = global.MysqlDB.Model(&models.SysRoleAuths{}).Where("role_id in ?", RoleIds).Find(&menuAuthList).Error
		if err != nil {
			return
		}
		for _, menuItem := range menuAuthList {
			menuIds = append(menuIds, menuItem.AuthId)
		}
	}

	// 取出具体的菜单信息
	var sysMenuList []models.SysMenus
	db := global.MysqlDB.Model(&models.SysMenus{})
	if len(menuIds) > 0 {
		db = db.Where("id in ?", menuIds)
	}
	err = db.Find(&sysMenuList).Error
	if err != nil {
		return
	}

	for _, menu := range sysMenuList {
		if menu.Perm != "" {
			resp = append(resp, menu.Perm)
		}
	}
	return
}
