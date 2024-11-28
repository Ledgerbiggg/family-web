package middlewares

import "github.com/gin-gonic/gin"

// Base 中间件接口
type Base interface {
	Handle() gin.HandlerFunc
	Order() int
}
