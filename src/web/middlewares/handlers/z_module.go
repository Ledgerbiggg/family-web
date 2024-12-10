package handlers

import (
	"family-web-server/src/web/middlewares"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"middlewares",
	// 中间件管理者
	fx.Provide(middlewares.NewMiddlewareManager),
	// 注册服务
	fx.Invoke(NewErrorMiddleware),   // 错误中间件
	fx.Invoke(NewCorsMiddleware),    // 跨域中间件
	fx.Invoke(NewSessionMiddleware), // session中间件
	fx.Invoke(NewJwtMiddleware),     // jwt中间件
	fx.Invoke(NewCaptchaMiddleware), // 验证码中间件
)
