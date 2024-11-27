package base

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method string
	Path   string
	Handle gin.HandlerFunc
}

type ControllerBase interface {
	// GetRoutes 获取controller下面的所有的路由
	GetRoutes() []*Route
	// RegisterController 注册到ControllerManager
	RegisterController()
}
