package config

type System struct {
	Env        string `mapstructure:"env" json:"env" yaml:"env"`                            // 环境值
	Port       int    `mapstructure:"port" json:"port" yaml:"port"`                         // 端口值
	UseCronJob bool   `mapstructure:"use-cron-job" json:"use-cron-job" yaml:"use-cron-job"` // 端口值
}
