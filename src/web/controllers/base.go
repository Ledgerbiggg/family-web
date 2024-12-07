package controllers

import (
	"github.com/gin-gonic/gin"
)

// Route 路由映射
type Route struct {
	Method string
	Path   string
	Handle gin.HandlerFunc
}

// Base 所有的controller需要实现的接口
type Base interface {
	//GetRoot 获取一级路由
	GetRoot() string
	// GetRoutes 获取controller下面的所有的路由
	GetRoutes() []*Route
	// RegisterController 注册到ControllerManager
	RegisterController()
}
