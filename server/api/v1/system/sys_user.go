package system

import (
	"github.com/gin-gonic/gin"
	"github.com/sunyihoo/gin-vue3-admin/server/model/common/response"
	systemReq "github.com/sunyihoo/gin-vue3-admin/server/model/system/request"
)

// Login
// @Tags Base
// @Summary 用户登录
// @Produce application/json
// @Param data body systemReq.length 													true "用户名,密码,验证码"
// @Success 200 {object} response.Response{data=systemRes.LoginResponse, msg=string}	"返回包括用户信息,token,过期时间"
// @Router /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify()
}
