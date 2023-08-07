package global

import (
	"github.com/sunyihoo/gin-vue3-admin/server/config"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONDIG config.Server
)
