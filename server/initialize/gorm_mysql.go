package initialize

import (
	"github.com/sunyihoo/gin-vue3-admin/server/global"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	//mysqlConfig := mysql.Config{
	//	DSN:                       m.Dsn(), // DSN data source name
	//	DefaultStringSize:         191,     // string 类型字段的默认长度
	//	SkipInitializeWithVersion: false,   // 根据版本自动配置
	//}
	//if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm)
	return &gorm.DB{}
}
