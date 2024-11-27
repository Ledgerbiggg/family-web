package log

import (
	"go.uber.org/fx"
)

var Module = fx.Module("log",
	fx.Provide(NewConsoleLogger),
	fx.Invoke(func(c *ConsoleLogger) {
		c.Info("log register success")
	}),
)
