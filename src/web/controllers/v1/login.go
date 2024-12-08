package v1

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/controllers"
	"family-web-server/src/web/models/dto/login"
	"family-web-server/src/web/services/v1/interfaces"
	"family-web-server/src/web/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	c            *config.GConfig
	cm           *controllers.ControllerManager
	l            *log.ConsoleLogger
	loginService interfaces.ILoginService
}

func NewLoginController(
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

func (c *LoginController) GetRoot() string {
	return ""
}

func (c *LoginController) GetRoutes() []*controllers.Route {
	return []*controllers.Route{
		{Method: "GET", Path: "/captcha", Handle: c.captcha},
		{Method: "POST", Path: "/login", Handle: c.login},
		{Method: "POST", Path: "/register", Handle: c.register},
		{Method: "POST", Path: "/verify", Handle: c.verify},
		{Method: "POST", Path: "/logout", Handle: c.logout},
	}
}

func (c *LoginController) RegisterController() {
	c.cm.AddController(c)
}

// captcha 生成验证码
func (c *LoginController) captcha(context *gin.Context) {
	data, err := c.loginService.CaptchaService()
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

// login 登录
func (c *LoginController) login(context *gin.Context) {
	var u = &login.UserDto{}
	if err := context.ShouldBindJSON(u); err != nil {
		c.l.Error("参数绑定失败:" + err.Error())
		context.Error(common.BadRequestError)
		return
	}
	b, r, ps, _ := c.loginService.LoginService(u)
	if b != 0 {
		// 查询用户的角色
		token, err := utils.GenerateToken(b, u.Username, r, ps, c.c.ServiceName, c.c.Jwt.ExpireTime, c.c.Jwt.SecretKey)
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
func (c *LoginController) register(context *gin.Context) {
	var r = &login.RegisterDto{}
	// 参数绑定
	if err := context.ShouldBindJSON(r); err != nil {
		c.l.Error("参数绑定失败:" + err.Error())
		context.Error(common.BadRequestError)
		return
	}
	// 校验参数
	if r.Password != r.ConfirmPassword {
		c.l.Error("两次密码不一致")
		context.Error(common.BadRequestError)
		return
	}

	// 手机号格式校验
	if err := c.loginService.ValidatePhone(r.Username); err != nil {
		c.l.Error("手机号格式错误")
		context.Error(common.PhoneFormatError)
		return
	}

	// 注册用户
	if err := c.loginService.RegisterService(r); err != nil {
		c.l.Error("注册失败:" + err.Error())
		context.Error(common.UserIsExistError)
		return
	}
	context.JSON(200, common.NewSuccessResult(nil))

}

// Verify 找回密码
func (c *LoginController) verify(context *gin.Context) {
	var v = &login.VerifyDto{}
	if err := context.ShouldBindJSON(v); err != nil {
		c.l.Error("参数绑定失败:" + err.Error())
		context.Error(err)
		return
	}

	if err := c.loginService.VerifyService(v); err != nil {
		c.l.Error("找回密码失败:" + err.Error())
		context.Error(err)
		return
	}
	context.JSON(200, common.NewSuccessResult(nil))

}

// Logout 退出登录
func (c *LoginController) logout(context *gin.Context) {
	//TODO 使用redis去删除token
	context.JSON(200, common.NewSuccessResult(nil))
}
