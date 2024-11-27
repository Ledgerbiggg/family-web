package test5

import (
	"context"
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/pkg"
	"family-web-server/src/pkg/mysql"
	"go.uber.org/fx"
	"os"
	"testing"
)

var (
	gConfig *config.GConfig
	logger  *log.ConsoleLogger
	gorm    *mysql.GormDb
)

func TestMain(m *testing.M) {
	app := fx.New(
		config.Module, // 配置文件
		log.Module,    // 日志
		//web.Module,    // web服务
		pkg.Module, // 包
		fx.Populate(
			&gConfig, &logger, &gorm,
		),
	)
	app.Start(context.Background())
	code := m.Run()
	os.Exit(code)
}
