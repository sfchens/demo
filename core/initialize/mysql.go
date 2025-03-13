package initialize

import (
	configs "demo/configs"
	initalizeMysql "demo/core/initialize/mysql"
	"demo/core/models"
	"demo/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	var mysqlObj *gorm.DB
	switch global.ConfigAll.Mysql.DbType {
	case "mysql":
		mysqlObj = gormMysql()
		global.DB = mysqlObj
	default:
		mysqlObj = gormMysql()
		global.DB = mysqlObj
	}

	// 初始化相关数据
	if mysqlObj != nil {
		RegisterTables(mysqlObj)
		InitMysqlData(global.DB)
		db, _ := global.DB.DB()
		global.RegisterShutdownFunc(func() {
			_ = db.Close()
		})
	}
	return mysqlObj
}

func gormMysql() *gorm.DB {
	m := global.ConfigAll.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         255,   //string 类型字段的默认长度
		SkipInitializeWithVersion: false, //根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), initalizeMysql.Gorm.Config()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormMysqlByConfig 初始化Mysql数据库用过传入配置
func GormMysqlByConfig(m configs.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         255,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), initalizeMysql.Gorm.Config()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.JwtBlackList{},
	)
	if err != nil {
		global.GetZapLog().Error("Register tables failed", zap.Error(err))
		os.Exit(0)
	}
	global.GetZapLog().Info("Register tables success")
}

func InitMysqlData(db *gorm.DB) {
	//var (
	//	count int64
	//	err   error
	//)
	//
	//if err = db.Model(&model.ExchangeRate{}).Count(&count).Error; err == nil {
	//	if count == 0 { // 没有数据才要插入数据
	//		for _, rate := range source.ExchangeRates {
	//			db.Create(&rate)
	//		}
	//	}
	//}
	//
	//if err = db.Model(&model.ShopifyStore{}).Count(&count).Error; err == nil {
	//	if count == 0 { // 没有数据才要插入数据
	//		for _, store := range source.ShopifyStoresWithCountryCode {
	//			db.Create(&store)
	//		}
	//	}
	//}
}
