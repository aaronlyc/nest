package response

import "appnest/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
