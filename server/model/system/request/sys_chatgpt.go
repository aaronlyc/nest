package request

import (
	"appnest/server/model/common/request"
	"appnest/server/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
