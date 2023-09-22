package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/sunyihoo/gin-vue3-admin/server/global"
	"github.com/sunyihoo/gin-vue3-admin/server/router"
	"net/http"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	InstallPlugin(Router) // 安装插件
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	// 如果想要不适用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH =  http://localhost
	// 然后执行打包命令 npm run build，在打开下面4行注释
	//Router.LoadHTMLGlob("./dist/*.html") // npm 打包成dist的路径
	//Router.Static("/favicon.ico", "./dist/favicon.ico")
	//Router.Static("/static", "./dist/assets") // dist里面的静态资源
	//Router.Static("/", "./dist/index.html")   // 前端网页入口页面

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, http.Dir(global.GVA_CONFIG.Local.StorePath)) // 为用户头像和文件提供静态地址
	//Router.Use(middleware.LoadTls)                                                                  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口", "你的cre/pem文件", "你的key文件")

	//跨域,如需跨域可以打开下面的注释
	//Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CoresByRules()) // 按照配置的规则放行跨域请求
	// global.GVA_LOG.Info("use middleware  cors")
	//docs.Swagger()

	//Router.GET(global.GVA_CONFIG.System.RouterPrefix + "/swagger/*any", gin)
	global.GVA_LOG.Info("register swagger handler")

	// 方便统一添加路由组前缀 多服务器上使用
	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		// 健康检测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		//systemRouter.
	}

	return nil
}
