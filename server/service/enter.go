package service

import (
	"github.com/sunyihoo/gin-vue3-admin/server/service/example"
	"github.com/sunyihoo/gin-vue3-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
