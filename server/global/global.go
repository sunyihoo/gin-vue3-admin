package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/sunyihoo/gin-vue3-admin/server/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG config.Server
	GVA_REDIS  *redis.Client
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
	//GVA_Timer  = timer.New

	//BlockCache local
	lock sync.RWMutex
)
