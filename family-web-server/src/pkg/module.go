package pkg

import (
	"family-web-server/src/log"
	"family-web-server/src/pkg/mysql"
	"go.uber.org/fx"
)

var Module = fx.Module("pkg",
	fx.Provide(mysql.NewGorm),
	fx.Invoke(func(log *log.ConsoleLogger, g *mysql.GormDb) {
		log.Info("mysql register success")
	}),
)
