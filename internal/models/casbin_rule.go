package models

type CasbinRules struct {
	ID    int    `json:"id" gorm:"comment:名称"`
	Ptype string `json:"Ptype" gorm:"NOT NULL;comment:策略p,角色G"`
	V0    string `json:"v0" gorm:"comment:角色或用户"`
	V1    int    `json:"v1" gorm:"comment:用户或路由"`
	V2    int    `json:"v2" gorm:"comment:请求方式"`
	V3    int    `json:"v3" gorm:"comment:允许标识"`
	V4    int    `json:"v4" gorm:"comment:请求方式"`
	V5    int    `json:"v5" gorm:"comment:请求方式"`
}

func (*CasbinRules) TableName() string {
	return "casbin_rules"
}
