package handlers

import (
	"family-web-server/src/log"
	"family-web-server/src/web/middleware/manager"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type SessionMiddleware struct {
	l *log.ConsoleLogger
}

func NewSessionMiddleware(
	mwm *manager.MiddlewareManager,
	l *log.ConsoleLogger,
) *SessionMiddleware {
	s := &SessionMiddleware{}
	mwm.AddMiddleware(s)
	s.l = l
	return s
}

func (s *SessionMiddleware) Handle() gin.HandlerFunc {
	// 使用 Cookie 存储 session（可以自定义密钥）
	store := cookie.NewStore([]byte("ledger")) // 替换为你的密钥
	return sessions.Sessions("session", store)
}

func (s *SessionMiddleware) Order() int {
	return 3
}
