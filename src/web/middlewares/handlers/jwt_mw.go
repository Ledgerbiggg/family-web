package handlers

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/middlewares"
	"family-web-server/src/web/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

type JwtMiddleware struct {
	l *log.ConsoleLogger
	c *config.GConfig
}

func NewJwtMiddleware(
	mwm *middlewares.MiddlewareManager,
	l *log.ConsoleLogger,
	c *config.GConfig,
) *JwtMiddleware {
	j := &JwtMiddleware{}
	mwm.AddMiddleware(j)
	j.l = l
	j.c = c
	return j
}

func (j *JwtMiddleware) Handle() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		// 如果是/login 或者 /register 请求 或者 /verify，不需要验证 JWT
		if path == "/captcha" ||
			path == "/login" ||
			path == "/register" ||
			path == "/verify" {
			context.Next()
			return
		}
		// 获取 Token
		tokenString := strings.TrimPrefix(context.Request.Header.Get("Authorization"), "Bearer ")
		// 解析 Token
		claims, err := utils.ParseToken(tokenString, j.c.Jwt.SecretKey)
		// Token 校验不通过
		if err != nil {
			context.Error(common.UnauthorizedError)
			return
		}
		// 获取用户的权限
		username := claims.Username
		role := claims.Role
		permissions := claims.Permissions

		// 校验是否有权限进入路由
		for i := range permissions {
			p := permissions[i].Path
			if j.matchPath(path, p) {
				context.Set("username", username)
				context.Set("role", role)
				context.Set("permissions", permissions)
				context.Next()
				return
			}
		}
		// 没有权限进入路由
		context.Error(common.UnauthorizedError)
		context.Abort()
	}
}

func (j *JwtMiddleware) Order() int {
	return 2
}

// 路由匹配函数
func (j *JwtMiddleware) matchPath(path string, permissionPath string) bool {
	// 如果权限路径包含通配符
	if strings.Contains(permissionPath, "*") {
		// 判断是否以指定路径开头
		prefix := strings.Split(permissionPath, "*")[0]
		return strings.HasPrefix(path, prefix)
	}
	// 完全匹配
	return path == permissionPath
}
