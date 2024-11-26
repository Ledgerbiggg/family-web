package manager

import (
	"family-web-server/src/web/middleware/base"
	"sort"
)

type MiddlewareManager struct {
	middlewares []base.MiddlewareBase
}

func NewMiddlewareManager() *MiddlewareManager {
	return &MiddlewareManager{}
}

func (mm *MiddlewareManager) GetMiddlewares() []base.MiddlewareBase {
	// 排序 middlewares
	sort.Slice(mm.middlewares, func(i, j int) bool {
		return mm.middlewares[i].Order() < mm.middlewares[j].Order()
	})
	return mm.middlewares
}

func (mm *MiddlewareManager) AddMiddleware(mw base.MiddlewareBase) {
	mm.middlewares = append(mm.middlewares, mw)
}
