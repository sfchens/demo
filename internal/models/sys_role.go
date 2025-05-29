package models

import (
	"demo/global"
	"demo/internal/libs"
	"fmt"
	"gorm.io/gorm"
)

// SysRoles 角色管理
type SysRoles struct {
	CommonField
	Name   string `json:"name" gorm:"name"`     // 角色名称
	Code   string `json:"code" gorm:"code"`     // 角色编码
	Sort   int64  `json:"sort" gorm:"sort"`     // 排序
	Status int64  `json:"status" gorm:"status"` // 状态（1-显示 2-隐藏）
	Remark string `json:"remark" gorm:"remark"` // 备注

	RoleAuths []SysRoleAuths `json:"roleAuths" gorm:"foreignKey:RoleId;references:id"` // 角色与菜单的关系
	RoleApis  []SysRoleApis  `json:"roleApis"  gorm:"foreignKey:RoleId;references:id"` // 菜单与接口的关系
}

// TableName 表名称
func (m *SysRoles) TableName() string {
	return "sys_roles"
}

func (m *SysRoles) BeforeDelete(tx *gorm.DB) (err error) {
	// 清除casbin用户与角色关联
	if _, e := libs.NewCasbinLogic().DeleteRole(int(m.ID)); e != nil {
		err = fmt.Errorf("清除casbin角色权限异常: <%s>", e.Error())
	}
	// 清除数据库中用户与角色的关联
	return m.deleteRoleUser()
}

// 删除用户与角色的关联
func (m *SysRoles) deleteRoleUser() (err error) {
	if e := global.MysqlDB.Where("role_id = ?", m.ID).Delete(&SysUserRoles{}).Error; e != nil {
		err = fmt.Errorf("删除用户角色关联异常: <%s>", e.Error())
	}
	return
}
