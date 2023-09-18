package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sunyihoo/gin-vue3-admin/server/global"
	systemReq "github.com/sunyihoo/gin-vue3-admin/server/model/system/request"
)

func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJTW()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析信息失败,请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}
