package global

import configs "demo/configs"

var (
	ConfigDefaultPath = "config.yaml"
	ConfigReleasePath = "config-uat.yaml"
	ConfigUatPath     = "config-uat.yaml"
	ConfigProdPath    = "config-prod.yaml"
)

var (
	DefaultLoggerKey = "defaultLogger"

	ConfigAll configs.ServerConfig
)
