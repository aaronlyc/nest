package initialize

import (
	_ "appnest/server/source/example"
	_ "appnest/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
