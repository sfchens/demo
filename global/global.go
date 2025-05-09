package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	Logger             = make(map[string]*zap.Logger)
	DB                 *gorm.DB
	Cron               *cron.Cron
	LocalCache         local_cache.Cache
	Redis              *redis.Client
	ConcurrencyControl = &singleflight.Group{}
	Casbin             *casbin.SyncedEnforcer
)
