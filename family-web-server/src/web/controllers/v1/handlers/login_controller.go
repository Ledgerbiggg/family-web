package handlers

import (
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/controllers/v1/base"
	"family-web-server/src/web/controllers/v1/manager"
	"family-web-server/src/web/models/dto/login"
	"family-web-server/src/web/services/interfaces"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/steambap/captcha"
	"strings"
)

type Controller struct {
	cm           *manager.ControllerManager
	l            *log.ConsoleLogger
	loginService interfaces.ILoginService
}

func NewController(
	cm *manager.ControllerManager,
	l *log.ConsoleLogger,
	loginService interfaces.ILoginService,
) *Controller {
	c := &Controller{
		cm,
		l,
		loginService,
	}
	c.RegisterController()
	return c
}

func (c *Controller) GetRoutes() []*base.Route {
	return []*base.Route{
		{Method: "GET", Path: "/captcha", Handle: c.Captcha},
		{Method: "POST", Path: "/login", Handle: c.Login},
		{Method: "POST", Path: "/register", Handle: c.Register},
	}
}

func (c *Controller) RegisterController() {
	c.cm.AddController(c)
}

// Captcha 生成验证码
func (c *Controller) Captcha(context *gin.Context) {
	data, err := captcha.New(200, 100)
	if err != nil {
		context.Error(common.SystemServerErrorError)
		return
	}
	// 获取 session 存储
	session := sessions.Default(context)
	c.l.Info("生成验证码:====              " + data.Text)
	// 将正确答案存入 session
	session.Set("captcha", data.Text)
	session.Save()

	context.Header("Content-Type", "image/png")
	err = data.WriteImage(context.Writer)
	if err != nil {
		context.Error(common.SystemServerErrorError)
		return
	}
}

// Login 登录
func (c *Controller) Login(context *gin.Context) {
	var u = &login.UserDto{}
	if err := context.ShouldBindJSON(u); err != nil {
		context.Error(common.BadRequestError)
		return
	}
	// 获取 session 存储
	session := sessions.Default(context)

	// 获取 session 中保存的验证码答案
	captchaVal := session.Get("captcha")
	// 清除 session 中保存的验证码答案
	session.Delete("captcha")
	session.Save()
	if captchaVal == nil || strings.ToLower(captchaVal.(string)) != strings.ToLower(u.Captcha) {
		context.Error(common.CaptchaErrorError)
		return
	}
	b, _ := c.loginService.Login(u)
	if b {
		context.JSON(200, common.NewSuccessResult(nil))
	} else {
		context.Error(common.LoginErrorError)
	}

}

func (c *Controller) Register(context *gin.Context) {

}
