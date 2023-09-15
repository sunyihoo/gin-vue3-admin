package system

import "github.com/sunyihoo/gin-vue3-admin/server/global"

type SysOperationRecord struct {
	global.GVA_MODEL
	Ip           string  `json:"ip" form:"ip" gorm:"column:id;comment:请求IP"`                                  // 请求IP
	Method       string  `json:"method" form:"method" gorm:"column:method;comment:请求方法"`                      // 请求方法
	Path         string  `json:"path" form:"path" gorm:"column:path;comment:请求路径"`                            // 请求路径
	Status       string  `json:"status"  form:"status" gorm:"column:status;comment:请求状态"`                     // 请求状态
	Latency      string  `json:"latency" form:"latency" gorm:"column:latency;comment:延迟"`                     // 延迟
	Agent        string  `json:"agent" form:"agent" gorm:"column:agent;comment:代理"`                           // 代理
	ErrorMessage string  `json:"error_message" form:"error_message" gorm:"column:error_message;comment:错误信息"` // 错误信息
	Body         string  `json:"body" form:"body" gorm:"column:body;comment:请求Body"`                          // 请求body
	Resp         string  `json:"resp" form:"resp" gorm:"column:resp;comment:响应body"`                          // 响应body
	UserID       int     `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`                   // 用户ID
	User         SysUser `json:"user"`
}
