package system

import (
	"demo/app/request"
	"demo/core/models"
	"demo/global"
	"demo/utils"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (userService *UserService) Register(u models.SysUser) (userInter models.SysUser, err error) {
	var user models.SysUser
	if errors.Is(global.DB.Where("username == ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err = global.DB.Create(&u).Error
	return u, err
}

func (userService *UserService) Login(u *models.SysUser) (userInter *models.SysUser, err error) {
	if nil == global.DB {
		return nil, fmt.Errorf("DB not init")
	}
	var user models.SysUser
	err = global.DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}

		//var SysAuthorityMenus []system.SysAuthorityMenu
		//err = global.DB.Where("sys_authority_authority_id = ?", user.AuthorityId).Find(&SysAuthorityMenus).Error
		//if err != nil {
		//	return
		//}
		//
		//var MenuIds []string
		//
		//for i := range SysAuthorityMenus {
		//	MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuId)
		//}
		//
		//var am system.SysBaseMenu
		//ferr := global.DB.First(&am, "name = ? and id in (?)", user.Authority.DefaultRouter, MenuIds).Error
		//if errors.Is(ferr, gorm.ErrRecordNotFound) {
		//	user.Authority.DefaultRouter = "404"
		//}
	}
	return &user, err
}

func (userService *UserService) ChangePassword(u *models.SysUser, newPassword string) (userInter *models.SysUser, err error) {
	var user models.SysUser
	//if err = global.DB.Where("username = ?", u.Username).First(&user).Error; err != nil {
	//	return nil, err
	//}
	//if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
	//	return nil, errors.New("原密码错误")
	//}
	//user.Password = utils.BcryptHash(newPassword)
	//err = global.DB.Save(&user).Error
	return &user, err

}

func (userService *UserService) GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&models.SysUser{})
	var userList []models.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

// 设置用户单个权限
func (userService *UserService) SetUserAuthority(id uint, authorityId uint) (err error) {
	//assignErr := global.DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system.SysUserAuthority{}).Error
	//if errors.Is(assignErr, gorm.ErrRecordNotFound) {
	//	return errors.New("该用户无此角色")
	//}
	//err = global.DB.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityId).Error
	return
}

// 设置用户多个权限
func (userService *UserService) SetUserAuthorities(id uint, authorityIds []uint) (err error) {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		//TxErr := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
		//if TxErr != nil {
		//	return TxErr
		//}
		//var useAuthority []system.SysUserAuthority
		//for _, v := range authorityIds {
		//	useAuthority = append(useAuthority, system.SysUserAuthority{
		//		SysUserId: id, SysAuthorityAuthorityId: v,
		//	})
		//}
		//TxErr = tx.Create(&useAuthority).Error
		//if TxErr != nil {
		//	return TxErr
		//}
		//TxErr = tx.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityIds[0]).Error
		//if TxErr != nil {
		//	return TxErr
		//}
		//// 返回 nil 提交事务
		return nil
	})
}

func (userService *UserService) DeleteUser(id int) (err error) {
	//var user models.SysUser
	//err = global.DB.Where("id = ?", id).Delete(&user).Error
	//if err != nil {
	//	return err
	//}
	//err = global.DB.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
	return
}

func (userService *UserService) UpdateUserInfo(user models.SysUser) error {
	return global.DB.Updates(&user).Error
}

func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user models.SysUser, err error) {
	//var reqUser models.SysUser
	//err = global.DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	//if err != nil {
	//	return reqUser, err
	//}
	//
	//var SysAuthorityMenus []system.SysAuthorityMenu
	//err = global.DB.Where("sys_authority_authority_id = ?", reqUser.AuthorityId).Find(&SysAuthorityMenus).Error
	//if err != nil {
	//	return
	//}
	//
	//var MenuIds []string
	//
	//for i := range SysAuthorityMenus {
	//	MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuId)
	//}
	//
	//var am system.SysBaseMenu
	//ferr := global.DB.First(&am, "name = ? and id in (?)", reqUser.Authority.DefaultRouter, MenuIds).Error
	//if errors.Is(ferr, gorm.ErrRecordNotFound) {
	//	reqUser.Authority.DefaultRouter = "404"
	//}
	return
}

func (userService *UserService) FindUserById(id int) (user *models.SysUser, err error) {
	var u models.SysUser
	err = global.DB.Where("`id` = ?", id).First(&u).Error
	return &u, err
}

func (userService *UserService) FindUserByUuid(uuid string) (user *models.SysUser, err error) {
	var u models.SysUser
	if err = global.DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, nil
}

func (userService *UserService) ResetPassword(ID uint) (err error) {
	err = global.DB.Model(&models.SysUser{}).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
	return err
}

func (userService *UserService) SetUserInfo(req models.SysUser) error {
	return global.DB.Updates(&req).Error
}

func (s *UserService) GetByEmail(email string) (user models.SysUser, err error) {
	err = global.DB.Where("email = ?", email).Limit(1).Find(&user).Error
	return
}

func (s *UserService) GetById(id uint) (user models.SysUser, err error) {
	err = global.DB.Where("id = ?", id).Limit(1).Find(&user).Error
	return
}

// GetAllUsers 获取所有用户信息
func (s *UserService) GetAllUsers() ([]models.SysUser, error) {
	var users []models.SysUser
	err := global.DB.Find(&users).Error
	return users, err
}
