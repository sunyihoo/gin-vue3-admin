package system

import "github.com/sunyihoo/gin-vue3-admin/server/global"

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
