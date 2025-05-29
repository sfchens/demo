package models

import (
	"demo/global"
	"fmt"
	"gorm.io/gorm"
)

const (
	SysUserStatusNormal  = 1
	SysUserStatusDisable = 2
)

// SysUsers 系统用户表
type SysUsers struct {
	CommonField
	Username      string     `json:"username" gorm:"username"`             // 用户名
	Nickname      string     `json:"nickname" gorm:"nickname"`             // 用户昵称
	Password      string     `json:"password" gorm:"password"`             // 密码
	Salt          string     `json:"salt" gorm:"salt"`                     // 加密盐
	Mobile        string     `json:"mobile" gorm:"mobile"`                 // 手机号
	Gender        int64      `json:"gender" gorm:"gender"`                 // 性别(1-男 2-女 0-保密)
	Email         string     `json:"email" gorm:"email"`                   // 邮箱
	Avatar        string     `json:"avatar" gorm:"avatar"`                 // 头像
	Status        int64      `json:"status" gorm:"status"`                 // 状态 1:正常,2:禁用
	DeptId        uint       `json:"deptId" gorm:"dept_id"`                // 部门ID
	RoleId        uint       `json:"roleId" gorm:"role_id"`                // 角色ID
	Remark        string     `json:"remark" gorm:"remark"`                 // 备注
	CreateBy      string     `json:"createBy" gorm:"create_by"`            // 创建者ID
	UpdateBy      string     `json:"updateBy" gorm:"update_by"`            // 更新者ID
	LastLoginTime int64      `json:"lastLoginTime" gorm:"last_login_time"` // 最后一次登录的时间
	LastLoginIp   string     `json:"lastLoginIp" gorm:"last_login_ip"`     // 最后一次登录的IP
	UserRoles     []SysRoles `gorm:"-" json:"userRoles" `
	RoleIds       []int      `gorm:"-" json:"role_ids"`
}

// TableName 表名称
func (m *SysUsers) TableName() string {
	return "sys_users"
}

// AfterFind 获取用户后，获取用户的角色信息
func (m *SysUsers) AfterFind(tx *gorm.DB) (err error) {
	if e := global.MysqlDB.Model(&SysUserRoles{}).Where("user_id = ?", m.ID).Pluck("role_id", &m.RoleIds).Error; e != nil {
		err = fmt.Errorf("根据权限ID<%d>获取菜单ID异常: <%s>", m.ID, e.Error())
		return
	}
	if e := global.MysqlDB.Model(&SysRoles{}).Where("id in ?", m.RoleIds).Find(&m.UserRoles).Error; e != nil {
		err = fmt.Errorf("根据菜单ID<%+v>获取菜单信息异常: <%s>", m.RoleIds, e.Error())
		return
	}
	return
}

// AfterCreate 添加用户后，创建用户与角色的关联
//func (t *SysUsers) AfterCreate(tx *gorm.DB) (err error) {
//	// 添加用户后，要添加用户角色并且添加到casbin
//	if _, e := system_service.NewCasbin().AddUserRoles([]string{t.Username}, t.RoleIds); e != nil {
//		err = fmt.Errorf("关联用户和角色到casbin异常: <%s>", e.Error())
//		return
//	}
//	err = t.bulkCreateUserRole()
//	return
//}

//// 批量创建用户与角色的关联
//func (t *SysUsers) bulkCreateUserRole() (err error) {
//	bulks := make([]*SysUserRoles, 0)
//	for _, v := range t.RoleIds {
//		bulks = append(bulks, &SysUserRoles{RoleId: v, UserId: int(t.ID)})
//	}
//	if e := global.MysqlDB.Create(&bulks).Error; e != nil {
//		err = fmt.Errorf("关联用户角色异常: <%s>", e.Error())
//	}
//	return err
//}
//
//// 删除用户与角色的关联
//func (t *SysUsers) deleteUserRole() (err error) {
//	if e := global.MysqlDB.Where("user_id = ?", t.ID).Delete(&SysUserRoles{}).Error; e != nil {
//		err = fmt.Errorf("删除用户角色关联异常: <%s>", e.Error())
//	}
//	return
//}
//
//// BeforeUpdate 更新用户信息前，先清除用户角色关联，然后再重新添加
//func (t *SysUsers) BeforeUpdate(tx *gorm.DB) (err error) {
//	// 清除casbin用户和角色关联
//	if _, e := system_service.NewCasbin().DeleteUserRole(t.Username); e != nil {
//		err = fmt.Errorf("删除用户<%s>的casbin角色关联异常: <%s>", t.Username, e.Error())
//		return
//	}
//	// 添加用户和角色关联
//	if err = t.deleteUserRole(); err != nil {
//		return
//	}
//	// 重新构建casbin用户和角色
//	return t.AfterCreate(tx)
//}
//
//// BeforeDelete 删除用户前清除用户与角色的关联信息
//func (t *SysUsers) BeforeDelete(tx *gorm.DB) (err error) {
//	if err = t.deleteUserRole(); err != nil {
//		return
//	}
//	return
//}
