package handlers

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/middlewares"
	"family-web-server/src/web/services/v1/interfaces"
	"family-web-server/src/web/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

type JwtMiddleware struct {
	l  *log.ConsoleLogger
	c  *config.GConfig
	ls interfaces.ILoginService
}

func NewJwtMiddleware(
	mwm *middlewares.MiddlewareManager,
	l *log.ConsoleLogger,
	c *config.GConfig,
	ls interfaces.ILoginService,
) *JwtMiddleware {
	j := &JwtMiddleware{}
	mwm.AddMiddleware(j)
	j.l = l
	j.c = c
	j.ls = ls
	return j
}

func (j *JwtMiddleware) Handle() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := strings.Replace(context.Request.URL.Path, "/"+j.c.ServerLevel, "", 1)
		// 如果是/Invite 或者 /register 请求 或者 /verify，不需要验证 JWT
		if path == "/captcha" ||
			path == "/login" ||
			path == "/register" ||
			path == "/verify" ||
			path == "/invite/info" ||
			path == "/invite/register" ||
			strings.HasPrefix(path, "/swagger") {
			context.Next()
			return
		}
		// 获取 Token
		tokenString := strings.TrimSpace(strings.TrimPrefix(context.Request.Header.Get("Authorization"), "Bearer"))
		// 解析 Token
		claims, err := utils.ParseToken(tokenString, j.c.Jwt.SecretKey)
		// Token 校验不通过
		if err != nil {
			j.l.Error("Token 校验失败:" + err.Error())
			context.Error(common.LoginExpiredError)
			context.Abort()
			return
		}
		// 获取用户的权限
		role, permissions, _ := j.ls.GetRoleAndPermissionByUserId(claims.UserId)
		// 如果是登出请求
		if path == "/logout" {
			// 直接进入路由
			context.Next()
			return
		}

		// 校验是否有权限进入路由
		for i := range permissions {
			p := permissions[i].Path
			if j.matchPath(path, p) {
				context.Set("userId", claims.UserId)
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
		// 获取通配符前的路径部分（即 `prefix`），假设是以 * 结尾
		prefix := strings.Split(permissionPath, "*")[0]

		// 检查路径是否以该前缀开头，且允许完全等于该前缀（即匹配 /a 和 /a/xxx 都符合）
		return strings.HasPrefix(path, prefix) || path+"/" == prefix || path == prefix
	}
	// 完全匹配
	return path == permissionPath
}
