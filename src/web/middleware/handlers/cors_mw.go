package handlers

import (
	"family-web-server/src/log"
	"family-web-server/src/web/middleware/manager"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CorsMiddleware struct {
	l *log.ConsoleLogger
}

func NewCorsMiddleware(
	mwm *manager.MiddlewareManager,
	l *log.ConsoleLogger,
) *CorsMiddleware {
	c := &CorsMiddleware{}
	mwm.AddMiddleware(c)
	c.l = l
	return c
}

func (c *CorsMiddleware) Handle() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 允许跨域的请求来源
		context.Header("Access-Control-Allow-Origin", "*") // 允许前端的域名
		// 必须设置为 true，表示允许携带凭证（cookie）
		context.Header("Access-Control-Allow-Credentials", "true")
		// 允许的方法（OPTIONS 是预检请求的标准方法）
		context.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		// 允许的请求头
		context.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")
		// 设置支持的最大请求时间
		context.Header("Access-Control-Max-Age", "86400")

		// 如果是预检请求（OPTIONS 请求），直接返回 200
		if context.Request.Method == http.MethodOptions {
			context.AbortWithStatus(http.StatusOK)
			return
		}

		// 执行后续的处理逻辑
		context.Next()
	}
}

func (c *CorsMiddleware) Order() int {
	return 0
}
