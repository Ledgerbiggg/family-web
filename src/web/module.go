package web

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	controllerManager "family-web-server/src/web/controllers"
	controllerHandlers "family-web-server/src/web/controllers/v1"
	middlewareManager "family-web-server/src/web/middlewares"
	middlewareHandlers "family-web-server/src/web/middlewares/handlers"
	"family-web-server/src/web/services/impls"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module("web",
	// 注册中间件
	fx.Provide(middlewareManager.NewMiddlewareManager), // 中间件管理者
	fx.Invoke(middlewareHandlers.NewErrorMiddleware),   // 错误中间件
	fx.Invoke(middlewareHandlers.NewCorsMiddleware),    // 跨域中间件
	fx.Invoke(middlewareHandlers.NewSessionMiddleware), // session中间件
	fx.Invoke(middlewareHandlers.NewJwtMiddleware),     // jwt中间件
	fx.Invoke(middlewareHandlers.NewCaptchaMiddleware), // 验证码中间件
	// 注册控制器
	fx.Provide(controllerManager.NewControllerManager), // 控制器管理者
	fx.Invoke(controllerHandlers.NewLoginController),   // 登录控制器
	fx.Invoke(controllerHandlers.NewHomeController),    // 主页控制器
	fx.Invoke(controllerHandlers.NewInviteController),  // 邀请控制器
	fx.Invoke(controllerHandlers.NewAlbumController),   // 相册控制器
	// 注册服务
	fx.Provide(impls.NewLoginService),  // 登录服务
	fx.Provide(impls.NewHomeService),   // 登录服务
	fx.Provide(impls.NewInviteService), // 邀请服务
	fx.Invoke(func(
		c *config.GConfig,
		l *log.ConsoleLogger,
		mwm *middlewareManager.MiddlewareManager,
		cm *controllerManager.ControllerManager,
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
				r.Handle(controller.GetRoutes()[j].Method,
					"/"+c.ServerLevel+"/"+controller.GetRoutes()[j].Path,
					controller.GetRoutes()[j].Handle)
			}
		}

		err := r.Run(fmt.Sprintf(":%d", c.Address.Port))
		if err != nil {
			l.Error(fmt.Sprintf("failed to start server: %v", err))
		}
	}),
)
