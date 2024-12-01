package handlers

import (
	"errors"
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/middlewares"
	"github.com/gin-gonic/gin"
)

type ErrorMiddleware struct {
	l *log.ConsoleLogger
}

func NewErrorMiddleware(
	mwm *middlewares.MiddlewareManager,
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
		if len(c.Errors) > 0 {
			if errors.As(c.Errors[0], &knownError) {
				r.l.Error(knownError.String())
				c.JSON(200, knownError)
			} else {
				r.l.Error(c.Errors[0].Error())
				c.JSON(200, common.UnknownError)
			}
		}
	}
}

func (r *ErrorMiddleware) Order() int {
	return 1
}
