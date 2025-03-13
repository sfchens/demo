package config

type ServerConfig struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Jwt    Jwt    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
}
