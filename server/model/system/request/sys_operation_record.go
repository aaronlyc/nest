package request

import (
	"appnest/server/model/common/request"
	"appnest/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
