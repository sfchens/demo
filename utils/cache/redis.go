package cache

import (
	"context"
	"demo/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strings"
	"time"
)

type RedisCache struct {
	prefix string
	client *redis.Client
	ctx    context.Context
	logger *zap.SugaredLogger
}

var (
	Instance *RedisCache
)

// Redis 返回 RedisCache 单实例
func Redis() *RedisCache {
	if Instance == nil {
		Instance = NewRedisCache()
	}
	return Instance
}

// NewRedisCache 创建 RedisCache 实例
func NewRedisCache() *RedisCache {
	prefix := global.ConfigAll.Redis.Prefix
	if len(prefix) == 0 {
		prefix = "RedisCache"
	}

	return &RedisCache{
		prefix: prefix,
		client: global.Redis,
		ctx:    context.Background(),
		logger: global.GetZapLog().Sugar(),
	}
}

// Set 设置指定 key 的值
func (ru RedisCache) Set(key string, value interface{}, expirationSec int) (err error) {
	err = ru.client.Set(ru.ctx, ru.Key(key), value, ru.Expiration(expirationSec)).Err()
	if err != nil {
		ru.logger.Errorf("redisUtil.Set err: err=[%+v]", err)
	}
	return err
}

// Get 获取指定 key 的值
func (ru RedisCache) Get(key string) (value string, err error) {
	value, err = ru.client.Get(ru.ctx, ru.Key(key)).Result()
	return value, err
}

// SAdd 将数据放入集合缓存
func (ru RedisCache) SAdd(key string, values ...interface{}) error {
	err := ru.client.SAdd(ru.ctx, ru.Key(key), values...).Err()
	if err != nil {
		ru.logger.Errorf("redisUtil.SSet err: err=[%+v]", err)
	}
	return err
}

// SAddEX 将数据放入集合，并指定过期时间
func (ru RedisCache) SAddEX(key string, expirationSec int, values ...interface{}) error {
	err := ru.client.SAdd(ru.ctx, ru.Key(key), values...).Err()
	if err != nil {
		ru.logger.Errorf("redisUtil.SSet err: err=[%+v]", err)
	}

	if expirationSec > 0 {
		// 注意需要传递原来的 key，因为 Expire 中已有添加前缀的逻辑
		err = ru.Expire(key, expirationSec)
	}

	return err
}

// SMembers 根据 key 返回集合中的所有成员
func (ru RedisCache) SMembers(key string) ([]string, error) {
	res, err := ru.client.SMembers(ru.ctx, ru.Key(key)).Result()
	if err != nil {
		ru.logger.Errorf("redisUtil.SGet err: err=[%+v]", err)
		return []string{}, err
	}
	return res, nil
}

// Del 删除一个或多个键
//
//	注意：键名会自动添加前缀，如果想删除给定键名的值，可使用 Delete
func (ru RedisCache) Del(keys ...string) (err error) {
	fullKeys := ru.toFullKeys(keys)
	err = ru.client.Del(ru.ctx, fullKeys...).Err()
	if err != nil {
		ru.logger.Errorf("redisUtil.HDel err: err=[%+v]", err)
	}
	return err
}

// Delete 删除给定键名的值
func (ru RedisCache) Delete(keys ...string) (err error) {
	err = ru.client.Del(ru.ctx, keys...).Err()
	if err != nil {
		ru.logger.Errorf("redisUtil.Delete err: err=[%+v]", err)
	}
	return err
}

// Exists 判断多项 key 是否存在
func (ru RedisCache) Exists(keys ...string) int64 {
	fullKeys := ru.toFullKeys(keys)
	cnt, err := ru.client.Exists(ru.ctx, fullKeys...).Result()
	if err != nil {
		ru.logger.Errorf("redisUtil.Exists err: err=[%+v]", err)
		return -1
	}
	return cnt
}

// Expire 指定缓存失效时间
func (ru RedisCache) Expire(key string, expirationSec int) error {
	err := ru.client.Expire(context.Background(), ru.Key(key), ru.Expiration(expirationSec)).Err()
	if err != nil {
		ru.logger.Errorf("redisUtil.Expire err: err=[%+v]", err)
	}
	return err
}

// Keys 查找给定表达式的所有 Key
func (ru RedisCache) Keys(pattern string) ([]string, error) {
	return ru.client.Keys(ru.ctx, ru.Key(pattern)).Result()
}

// Scan 迭代查找
func (ru RedisCache) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return ru.client.Scan(ru.ctx, cursor, ru.Key(match), count).Result()
}

// FlushAll 清空整个 Redis 服务器的数据
func (ru RedisCache) FlushAll() error {
	return ru.client.FlushAll(ru.ctx).Err()
}

// TTL 根据 key 获取过期时间
func (ru RedisCache) TTL(key string) int {
	td, err := ru.client.TTL(context.Background(), ru.Key(key)).Result()
	if err != nil {
		ru.logger.Errorf("redisUtil.TTL err: err=[%+v]", err)
	}
	return int(td / time.Second)
}

// Key 返回添加统一前缀的键
func (ru RedisCache) Key(key string) string {
	return ru.prefix + ":" + strings.Trim(key, ":")
}

// Expiration 返回统一的时间
func (ru RedisCache) Expiration(sec int) time.Duration {
	return time.Duration(sec) * time.Second
}

// toFullKeys 为 keys 批量增加前缀
func (ru RedisCache) toFullKeys(keys []string) (fullKeys []string) {
	for _, name := range keys {
		fullKeys = append(fullKeys, ru.Key(name))
	}
	return
}

func (ru RedisCache) RedisKey(keys ...string) string {
	return RedisKey(keys...)
}

func RedisKey(keys ...string) string {
	key := ""
	for _, s := range keys {
		if len(s) > 0 {
			key += ":" + s
		}
	}
	return strings.Trim(key, ":")
}
