package router

import (
	"github.com/sunyihoo/gin-vue3-admin/server/router/example"
	"github.com/sunyihoo/gin-vue3-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
