package pkg

import (
	"family-web-server/src/logs"
	"family-web-server/src/pkg/mysql"
	"go.uber.org/fx"
)

var Module = fx.Module("pkg",
	fx.Provide(mysql.NewGorm),
	fx.Invoke(func(log *logs.ConsoleLogger, g *mysql.GormDb) {
		log.Info("mysql register success")
	}),
)
