package models

type JwtBlackList struct {
	CommonField
	Jwt string `gorm:"type:text;comment:jwt"`
}
