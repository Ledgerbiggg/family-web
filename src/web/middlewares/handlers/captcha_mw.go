package handlers

import (
	"bytes"
	"encoding/json"
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/middlewares"
	"family-web-server/src/web/services/interfaces"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
)

type CaptchaMiddleware struct {
	l            *log.ConsoleLogger
	c            *config.GConfig
	loginService interfaces.ILoginService
}

func (cm *CaptchaMiddleware) Handle() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := strings.Replace(context.Request.URL.Path, "/"+cm.c.ServerLevel, "", 1)
		// 判断路由,如果是登录+注册+忘记密码+邀请注册，不需要验证码
		if path == "/login" || path == "/register" || path == "/verify" || path == "/invite-register" {
			// 获取 session 存储
			session := sessions.Default(context)
			// 获取 session 中保存的验证码答案
			captchaVal := session.Get("captcha")
			if captchaVal == nil {
				context.Error(common.CaptchaGetError)
				context.Abort()
				return
			}
			storedCaptcha := captchaVal.(string)
			// 清除 session 中保存的验证码答案
			session.Delete("captcha")
			session.Save()
			// 获取请求体中的验证码

			// 获取请求体中的验证码（处理 JSON 请求体）
			var requestBody map[string]any
			if err := context.ShouldBindJSON(&requestBody); err != nil {
				context.JSON(200, common.BadRequestError)
				context.Abort()
				return
			}

			inputCaptcha, exists := requestBody["captcha"].(string)
			if !exists || storedCaptcha == "" || strings.ToLower(storedCaptcha) != strings.ToLower(inputCaptcha) {
				context.Error(common.CaptchaErrorError)
				context.Abort()
				return
			}
			// 如果需要继续读取原始请求体，可以手动将其重新设置回 cm.Request.Body
			// 重新设置 cm.Request.Body 为解析后的内容，以便后续处理
			bodyBytes, _ := json.Marshal(requestBody)
			context.Request.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}
		context.Next()
	}
}

func (cm *CaptchaMiddleware) Order() int {
	return 99999
}

func NewCaptchaMiddleware(
	mwm *middlewares.MiddlewareManager,
	c *config.GConfig,
	l *log.ConsoleLogger) *CaptchaMiddleware {
	cw := &CaptchaMiddleware{}
	mwm.AddMiddleware(cw)
	cw.l = l
	cw.c = c
	return cw
}
