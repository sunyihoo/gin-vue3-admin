package v1

import (
	"github.com/sunyihoo/gin-vue3-admin/server/api/v1/example"
	"github.com/sunyihoo/gin-vue3-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
