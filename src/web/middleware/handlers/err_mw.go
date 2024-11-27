package handlers

import (
	"errors"
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/middleware/manager"
	"github.com/gin-gonic/gin"
)

type ErrorMiddleware struct {
	l *log.ConsoleLogger
}

func NewErrorMiddleware(
	mwm *manager.MiddlewareManager,
	l *log.ConsoleLogger,
) *ErrorMiddleware {
	e := &ErrorMiddleware{}
	mwm.AddMiddleware(e)
	e.l = l
	return e
}

func (r *ErrorMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// 捕获错误
		var knownError *common.KnownError
		if len(c.Errors) > 0 && errors.As(c.Errors[0], &knownError) {
			r.l.Error(knownError.String())
			c.JSON(200, knownError)
		}
	}
}

func (r *ErrorMiddleware) Order() int {
	return 1
}
