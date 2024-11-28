package middlewares

import (
	"sort"
)

type MiddlewareManager struct {
	middlewares []Base
}

func NewMiddlewareManager() *MiddlewareManager {
	return &MiddlewareManager{}
}

func (mm *MiddlewareManager) GetMiddlewares() []Base {
	// 排序 middlewares
	sort.Slice(mm.middlewares, func(i, j int) bool {
		return mm.middlewares[i].Order() < mm.middlewares[j].Order()
	})
	return mm.middlewares
}

func (mm *MiddlewareManager) AddMiddleware(mw Base) {
	mm.middlewares = append(mm.middlewares, mw)
}
