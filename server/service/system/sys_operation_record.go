package system

import (
	"github.com/sunyihoo/gin-vue3-admin/server/global"
	"github.com/sunyihoo/gin-vue3-admin/server/model/system"
)

//@author: [s]()
//@function: CreateSysOperationRecord
//@description: 创建记录
//@Param: sysOperationRecord model.SysOperationRecord

type OperationRecordService struct{}

func (operationRecordService *OperationRecordService) CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}

//@author: [g]()
//@function
