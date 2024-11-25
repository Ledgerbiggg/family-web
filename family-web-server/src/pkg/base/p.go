package base

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	LC     fx.Lifecycle
	Log    *log.ConsoleLogger
	Config *config.GConfig
}
