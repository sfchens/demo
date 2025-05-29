package models

// 用户与角色的关联模型
type SysUserRoles struct {
	Id     int  `gorm:"primary_key" json:"id"`
	UserId uint `gorm:"column:user_id" json:"user_id"`
	RoleId uint `gorm:"column:role_id" json:"role_id"`
}

func (*SysUserRoles) TableName() string {
	return "sys_user_roles"
}
