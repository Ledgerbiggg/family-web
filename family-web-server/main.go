package main

import (
	"family-web-server/src/config"
	"family-web-server/src/logs"
	"family-web-server/src/pkg"
	"family-web-server/src/web"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module, // 配置文件
		logs.Module,   // 日志
		web.Module,    // web服务
		pkg.Module,    // 包
	)
	app.Run()
}
