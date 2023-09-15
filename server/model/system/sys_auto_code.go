package system

import "github.com/sunyihoo/gin-vue3-admin/server/global"

type SysAutoCode struct {
	global.GVA_MODEL
	PackageName string `json:"packageName" gorm:"comment:包名"`
	Label       string `json:"label" gorm:"comment:展示名"`
	Desc        string `json:"desc" gorm:"comment:描述"`
}
