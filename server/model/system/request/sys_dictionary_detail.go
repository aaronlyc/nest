package request

import (
	"appnest/server/model/common/request"
	"appnest/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
