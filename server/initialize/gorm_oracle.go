package initialize

import (
	"github.com/dzwvip/oracle"
	"github.com/sunyihoo/gin-vue3-admin/server/global"
	"github.com/sunyihoo/gin-vue3-admin/server/initialize/internal"
	"gorm.io/gorm"

	_ "github.com/godror/godror"
)

// GormOracle 初始化oracle数据库
// 如果需要Oracle库 放开import 里面的注释 把下方 mysql.Config 改为 oracle.Config ; mysql.New 改为 oracle.New
func GormOracle() *gorm.DB {
	m := global.GVA_CONFIG.Oracle
	if m.Dbname == "" {
		return nil
	}
	oracleConfig := oracle.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	if db, err := gorm.Open(oracle.New(oracleConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}

}
