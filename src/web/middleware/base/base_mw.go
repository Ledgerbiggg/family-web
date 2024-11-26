package base

import "github.com/gin-gonic/gin"

type MiddlewareBase interface {
	Handle() gin.HandlerFunc
	Order() int
}
