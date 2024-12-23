package test5

import (
	"context"
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/pkg"
	"family-web-server/src/web"
	"family-web-server/src/web/services/v1/interfaces"
	"go.uber.org/fx"
	"os"
	"testing"
)

var (
	albumService interfaces.IAlbumService
)

func TestMain(m *testing.M) {
	app := fx.New(
		config.Module, // 配置文件
		log.Module,    // 日志
		web.Module,    // web服务
		pkg.Module,    // 包
		fx.Populate(&albumService),
	)
	app.Start(context.Background())
	code := m.Run()
	os.Exit(code)
}
