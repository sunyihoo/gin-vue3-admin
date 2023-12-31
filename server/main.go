package main

import (
	"database/sql"
	"fmt"
	"github.com/sunyihoo/gin-vue3-admin/server/core"
	"github.com/sunyihoo/gin-vue3-admin/server/global"
	"github.com/sunyihoo/gin-vue3-admin/server/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title 						Swagger Example API
// @version						0.0.1
// @description					This is a sample Server Pets
// @securityDefinitions.apiKey	ApiKeyAuth
// @in 							header
// @name 						x-token
// @BasePath 					/
func main() {
	fmt.Println("start")
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTable() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(db)
	}
	core.RunWindowsServer()
}
