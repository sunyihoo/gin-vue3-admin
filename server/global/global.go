package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"github.com/sunyihoo/gin-vue3-admin/server/config"
	"github.com/sunyihoo/gin-vue3-admin/server/utils/timer"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"sync"
)

var (
	GVA_DB                  *gorm.DB
	GVA_CONFIG              config.Server
	GVA_REDIS               *redis.Client
	GVA_VP                  *viper.Viper
	GVA_LOG                 *zap.Logger
	GVA_Timer               = timer.NewTimerTask()
	GVA_Concurrency_Control = &singleflight.Group{}

	BlockCache local_cache.Cache
	lock       sync.RWMutex
)
