package initialize

import (
	"github.com/sunyihoo/gin-vue3-admin/server/global"
	"github.com/sunyihoo/gin-vue3-admin/server/model/example"
	"github.com/sunyihoo/gin-vue3-admin/server/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTable() {
	db := global.GVA_DB
	err := db.AutoMigrate(
		//系统模块表
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},
		system.SysChatCptOption{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
