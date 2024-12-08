package main

import (
	_ "family-web-server/docs"
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/pkg"
	"family-web-server/src/web"
	"go.uber.org/fx"
)

// @title           My family-web-server API
// @version         1.0
// @description     family-web-server API description
// @termsOfService
// @contact.name   Ledgerbiggg
// @contact.url    https://github.com/Ledgerbiggg/family-web/issues
// @contact.email  ledgerbiggg@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8001
// @BasePath  /v1
// @securityDefinitions.basic  BasicAuth
func main() {
	app := fx.New(
		config.Module, // 配置文件
		log.Module,    // 日志
		web.Module,    // web服务
		pkg.Module,    // 包
	)
	app.Run()
}
