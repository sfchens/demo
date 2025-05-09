package models

// SysRoleApis 角色接口权限
type SysRoleApis struct {
	CommonField
	RoleId uint `json:"roleId" gorm:"role_id"` // sys_roles表Id
	ApiId  uint `json:"apiId" gorm:"api_id"`   // sys_apis表Id
}

// TableName 表名称
func (*SysRoleApis) TableName() string {
	return "sys_role_apis"
}
