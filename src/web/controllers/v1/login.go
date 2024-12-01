package v1

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/controllers"
	"family-web-server/src/web/models/dto/login"
	"family-web-server/src/web/models/vo"
	"family-web-server/src/web/services/interfaces"
	"family-web-server/src/web/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
		{Method: "POST", Path: "/logout", Handle: c.Logout},
		{Method: "POST", Path: "/invite", Handle: c.Invite},
		{Method: "GET", Path: "/qr-code", Handle: c.QrCode},
		{Method: "GET", Path: "/invite-info", Handle: c.InviteInfo},
		{Method: "POST", Path: "/invite-register", Handle: c.InviteRegister},
	}
}

func (c *LoginController) RegisterController() {
	c.cm.AddController(c)
}

// Captcha 生成验证码
func (c *LoginController) Captcha(context *gin.Context) {
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

// Login 登录
func (c *LoginController) Login(context *gin.Context) {
	var u = &login.UserDto{}
	if err := context.ShouldBindJSON(u); err != nil {
		c.l.Error("参数绑定失败:" + err.Error())
		context.Error(common.BadRequestError)
		return
	}
	b, r, ps, _ := c.loginService.LoginService(u)
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
func (c *LoginController) Verify(context *gin.Context) {
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
func (c *LoginController) Logout(context *gin.Context) {
	//TODO 使用redis去删除token
	context.JSON(200, common.NewSuccessResult(nil))
}

// Invite 邀请注册
func (c *LoginController) Invite(context *gin.Context) {
	// 校验参数
	var i = &login.InviteDto{}
	if err := context.ShouldBindJSON(i); err != nil {
		c.l.Error("参数绑定失败:" + err.Error())
		context.Error(err)
		return
	}
	// 获取当前用户名称
	uid, err := c.loginService.InviteService(context.GetString("username"), i)
	if err != nil {
		c.l.Error("邀请失败:" + err.Error())
		context.Error(err)
		return
	}
	context.JSON(200, common.NewSuccessResult(map[string]string{"uid": uid}))
}

// InviteInfo 根据邀请的uuid获取邀请信息
func (c *LoginController) InviteInfo(context *gin.Context) {
	// 校验参数
	uid := context.Query("uid")
	if uid == "" {
		c.l.Error("uid 为空")
		context.Error(common.BadRequestError)
		return
	}
	// 获取邀请信息
	info, err := c.loginService.CheckInviteInfoIsValid(uid)
	if err != nil {
		c.l.Error("获取邀请信息失败:" + err.Error())
		context.Error(err)
		return
	}
	context.JSON(200, common.NewSuccessResult(vo.NewInviteVo(info)))
}

// QrCode 获取二维码
func (c *LoginController) QrCode(context *gin.Context) {
	uid := context.Query("uid")
	// 获取二维码
	qrCode, err := utils.GenerateQRCode(fmt.Sprintf("http://%s/Invite-register?uid=%s", c.c.Address.Domain, uid), 100)
	if err != nil {
		c.l.Error("获取二维码失败:" + err.Error())
		context.Error(err)
		return
	}
	context.Header("Content-Type", "image/png")
	context.Writer.Write(qrCode)
}

// InviteRegister 使用邀请链接去注册(直接成为管理员)
func (c *LoginController) InviteRegister(context *gin.Context) {
	// 校验参数
	var i = &login.InviteRegisterDto{}
	if err := context.ShouldBindJSON(i); err != nil {
		c.l.Error("参数绑定失败:" + err.Error())
		context.Error(err)
		return
	}

	// 校验手机号是否一致
	if err := c.loginService.ValidatePhone(i.Username); err != nil {
		c.l.Error("手机号格式错误")
		context.Error(common.PhoneFormatError)
		return
	}

	err := c.loginService.InviteRegisterService(i)
	if err != nil {
		c.l.Error("邀请注册失败:" + err.Error())
		context.Error(err)
		return
	}

	context.JSON(200, common.NewSuccessResult(nil))
}
