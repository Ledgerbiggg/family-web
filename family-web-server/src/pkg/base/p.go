package base

import (
	"family-web-server/src/config"
	"family-web-server/src/logs"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	LC     fx.Lifecycle
	Log    *logs.ConsoleLogger
	Config *config.GConfig
}
