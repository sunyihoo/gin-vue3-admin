package core

import (
	"github.com/sunyihoo/gin-vue3-admin/server/global"
	"github.com/sunyihoo/gin-vue3-admin/server/initialize"
	"github.com/sunyihoo/gin-vue3-admin/server/service/system"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")
}
