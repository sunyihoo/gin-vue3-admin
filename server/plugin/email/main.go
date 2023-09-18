package email

import (
	"github.com/gin-gonic/gin"
	"github.com/sunyihoo/gin-vue3-admin/server/plugin/email/global"
	"github.com/sunyihoo/gin-vue3-admin/server/plugin/email/router"
)

type EmailPlugin struct{}

func CreateEmailPlug(To, From, Host, Secret, Nickname string, Port int, IsSSL bool) *EmailPlugin {
	global.Config.To = To
	global.Config.From = From
	global.Config.Host = Host
	global.Config.Secret = Secret
	global.Config.Nickname = Nickname
	global.Config.Port = Port
	return &EmailPlugin{}
}

func (*EmailPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitEmailRouter(group)
}

func (*EmailPlugin) RouterPath() string {
	return "email"
}
