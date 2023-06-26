package request

import (
	"appnest/server/model/common/request"
	"appnest/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
