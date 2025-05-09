package initialize

import (
	"context"
	"demo/global"
	"demo/utils/cache"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"time"
)

func Cache() {
	LocalCache()
	RedisCache()
}

func LocalCache() {
	global.LocalCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(global.ConfigAll.Jwt.ExpiresTime)),
	)
}

func RedisCache() {
	redisCfg := global.ConfigAll.Redis
	if global.ConfigAll.System.Env != "prod" {
		return
	}

	fmt.Printf("RedisCache")

	// @link https://redis.uptrace.dev/zh/guide/go-redis-debugging.html#%E8%BF%9E%E6%8E%A5%E6%B1%A0%E5%A4%A7%E5%B0%8F
	redisOptions := redis.Options{
		Addr:        redisCfg.Addr,
		Username:    redisCfg.Username,
		Password:    redisCfg.Password,
		DB:          redisCfg.DB,
		PoolSize:    80,
		PoolTimeout: 10 * time.Second, // 默认 4 秒，现在改为 10 秒（默认值：ReadTimeout + 1 秒, ReadTimeout 为 3 秒）
	}

	client := redis.NewClient(&redisOptions)
	rs, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GetZapLog().Error("redis connection ping failed, err:", zap.Error(err))
	} else {
		global.GetZapLog().Info("redis connection ping response:", zap.String("RedisPing", rs))
		global.Redis = client
		// 初始化使用 Redis 的模块
		cache.Redis()
	}
}
