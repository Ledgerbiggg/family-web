package v1

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/controllers"
	"family-web-server/src/web/models/dto/login"
	"family-web-server/src/web/services/interfaces"
	"family-web-server/src/web/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/steambap/captcha"
	"strings"
)

type LoginController struct {
	c            *config.GConfig
	cm           *controllers.ControllerManager
	l            *log.ConsoleLogger
	loginService interfaces.ILoginService
}

func NewController(
	cf *config.GConfig,
	cm *controllers.ControllerManager,
	l *log.ConsoleLogger,
	ls interfaces.ILoginService,
) *LoginController {
	c := &LoginController{
		c:            cf,
		cm:           cm,
		l:            l,
		loginService: ls,
	}
	c.RegisterController()
	return c
}

func (c *LoginController) GetRoutes() []*controllers.Route {
	return []*controllers.Route{
		{Method: "GET", Path: "/captcha", Handle: c.Captcha},
		{Method: "POST", Path: "/login", Handle: c.Login},
		{Method: "POST", Path: "/register", Handle: c.Register},
		{Method: "POST", Path: "/verify", Handle: c.Verify},
	}
}

func (c *LoginController) RegisterController() {
	c.cm.AddController(c)
}

// Captcha 生成验证码
func (c *LoginController) Captcha(context *gin.Context) {
	data, err := captcha.New(200, 100)
	if err != nil {
		c.l.Error("生成验证码失败:" + err.Error())
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
		c.l.Error("写入验证码失败:" + err.Error())
		context.Error(common.SystemServerErrorError)
		return
	}
}

// Login 登录 TODO 不要在controller里面写业务逻辑
func (c *LoginController) Login(context *gin.Context) {
	var u = &login.UserDto{}
	if err := context.ShouldBindJSON(u); err != nil {
		c.l.Error("参数绑定失败:" + err.Error())
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
		c.l.Error("验证码错误")
		context.Error(common.CaptchaErrorError)
		return
	}
	b, r, ps, _ := c.loginService.Login(u)
	if b {
		// 查询用户的角色
		token, err := utils.GenerateToken(u.Username, r, ps, c.c.ServiceName, c.c.Jwt.ExpireTime, c.c.Jwt.SecretKey)
		if err != nil {
			c.l.Error("生成 token 失败:" + err.Error())
			context.Error(common.SystemServerErrorError)
			return
		}
		context.Header("token", token)
		context.JSON(200, common.NewSuccessResult(nil))
	} else {
		context.Error(common.LoginErrorError)
	}

}

// Register 注册
func (c *LoginController) Register(context *gin.Context) {
}

// Verify 找回密码
func (c *LoginController) Verify(context *gin.Context) {

}
