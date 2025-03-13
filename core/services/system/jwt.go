package system

import (
	"context"
	"demo/core/models"
	"demo/global"
	"go.uber.org/zap"
	"time"
)

type JwtService struct {
}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (jwtService *JwtService) JsonInBlackList(jwtList models.JwtBlackList) (err error) {
	err = global.DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.LocalCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

func (jwtService *JwtService) IsBlackList(jwt string) bool {
	_, ok := global.LocalCache.Get(jwt)
	return ok
}

func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJwt string) {
	redisJwt, err = global.Redis.Get(context.Background(), userName).Result()
	return err, redisJwt
}

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	timer := time.Duration(global.ConfigAll.Jwt.ExpiresTime) * time.Second
	err = global.Redis.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAllJwt() {
	if global.DB.Migrator().HasTable(&models.JwtBlackList{}) {
		var data []string
		err := global.DB.Model(&models.JwtBlackList{}).Select("jwt").Find(&data).Error
		if err != nil {
			global.GetZapLog().Error("加载数据库jwt黑名单失败", zap.Error(err))
		}
		for i := 0; i < len(data); i++ {
			global.LocalCache.SetDefault(data[i], struct{}{})
		}
	}
}
