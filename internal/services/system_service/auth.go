package system_service

import (
	"context"
	"demo/app/request"
	"demo/global"
	"demo/internal/models"
	"demo/utils"
	"demo/utils/jwt"
	"errors"
	"fmt"
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

	if err = global.DB.Model(&models.SysUsers{}).Where("username = ?", params.Username).First(user).Error; err != nil {
		return
	}
	if user.Status != models.SysUserStatusNormal {
		err = errors.New("用户已被禁用")
		return
	}
	fmt.Printf("user:  %+v\n", user)

	if ok := utils.BcryptCheck(params.Password, user.Password); !ok {
		err = errors.New("密码错误")
		return
	}

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

func (self *AuthLogic) Codes(ctx context.Context) (resp []string, err error) {
	resp = []string{
		"AC_100100",
		"AC_100110",
		"AC_100120",
		"AC_100010",
	}
	return
}
