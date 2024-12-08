package web

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/web/controllers"
	controllersV1 "family-web-server/src/web/controllers/v1"
	"family-web-server/src/web/middlewares"
	"family-web-server/src/web/middlewares/handlers"
	servicesV1 "family-web-server/src/web/services/v1"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

var Module = fx.Module("web",
	// 注册中间件
	handlers.Module,
	// 注册控制器
	controllersV1.Module,
	// 注册服务
	servicesV1.Module,
	fx.Invoke(func(
		c *config.GConfig,
		l *log.ConsoleLogger,
		mwm *middlewares.MiddlewareManager,
		cm *controllers.ControllerManager,
	) {
		// 创建一个 Gin 引擎实例
		r := gin.Default()
		// 获取所有的控制器
		cs := cm.GetControllers()
		// 获取所有的中间件
		middlewares := mwm.GetMiddlewares()

		// 将所有的中间件注册进去
		for i := range middlewares {
			r.Use(middlewares[i].Handle())
		}

		// 将所有的路由和处理函数注册进去
		for i := range cs {
			controller := cs[i]
			for j := range controller.GetRoutes() {
				// 使用路由版本号+控制器主路由+控制器子路由拼接
				r.Handle(controller.GetRoutes()[j].Method,
					"/"+c.ServerLevel+"/"+
						controller.GetRoot()+
						controller.GetRoutes()[j].Path,
					controller.GetRoutes()[j].Handle,
				)
			}
		}
		// 如果是本地开发环境的话 启用 Swagger UI
		if c.Mode == "dev" {
			r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		err := r.Run(fmt.Sprintf(":%d", c.Address.Port))
		if err != nil {
			l.Error(fmt.Sprintf("failed to start server: %v", err))
		}
	}),
)
