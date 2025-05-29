package system_service

import (
	"context"
	"demo/global"
	"demo/internal/models"
	"go.uber.org/zap"
	"time"
)

type JwtLogic struct {
}

func NewJwtLogic() *JwtLogic {
	return &JwtLogic{}
}

func (JwtLogic *JwtLogic) JsonInBlackList(jwtList models.JwtBlackList) (err error) {
	err = global.MysqlDB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.LocalCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

func (JwtLogic *JwtLogic) IsBlackList(jwt string) bool {
	_, ok := global.LocalCache.Get(jwt)
	return ok
}

func (JwtLogic *JwtLogic) GetRedisJWT(userName string) (err error, redisJwt string) {
	redisJwt, err = global.Redis.Get(context.Background(), userName).Result()
	return err, redisJwt
}

func (JwtLogic *JwtLogic) SetRedisJWT(jwt string, userName string) (err error) {
	timer := time.Duration(global.ConfigAll.Jwt.ExpiresTime) * time.Second
	err = global.Redis.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAllJwt() {
	if global.MysqlDB.Migrator().HasTable(&models.JwtBlackList{}) {
		var data []string
		err := global.MysqlDB.Model(&models.JwtBlackList{}).Select("jwt").Find(&data).Error
		if err != nil {
			global.GetZapLog().Error("加载数据库jwt黑名单失败", zap.Error(err))
		}
		for i := 0; i < len(data); i++ {
			global.LocalCache.SetDefault(data[i], struct{}{})
		}
	}
}
