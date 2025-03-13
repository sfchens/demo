package config

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	Prefix   string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`       // Key 前缀
	Lifetime string `mapstructure:"lifetime" json:"lifetime" yaml:"lifetime"` // Key 前缀
}
