package initialize

import (
	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	InstallPlugin(Router) // 安装插件
	//systemRouter := router.RouterGroupApp
	return nil
}
