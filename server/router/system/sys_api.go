package system

import (
	"github.com/gin-gonic/gin"
	"github.com/sunyihoo/gin-vue3-admin/server/middleware"
)

type ApiRouter struct{}

func (s *ApiRouter) InitRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("api")

	apiPublicRouterWithoutRecord := RouterPub.Group("api")
	apiRouterApi := v1.ApiGroup
}
