package v1

import (
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/web/common"
	"family-web-server/src/web/controllers"
	"family-web-server/src/web/models/dto/login"
	login2 "family-web-server/src/web/models/vo/Invite"
	"family-web-server/src/web/services/interfaces"
	"family-web-server/src/web/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type InviteController struct {
	c             *config.GConfig
	cm            *controllers.ControllerManager
	l             *log.ConsoleLogger
	loginService  interfaces.ILoginService
	inviteService interfaces.IInviteService
}

func NewInviteController(
	c *config.GConfig,
	cm *controllers.ControllerManager,
	l *log.ConsoleLogger,
	loginService interfaces.ILoginService,
	inviteService interfaces.IInviteService,
) *InviteController {
	i := &InviteController{
		c:             c,
		cm:            cm,
		l:             l,
		loginService:  loginService,
		inviteService: inviteService,
	}
	i.RegisterController()
	return i
}

func (ic *InviteController) GetRoot() string {
	return "/invite"
}

func (ic *InviteController) GetRoutes() []*controllers.Route {
	return []*controllers.Route{
		{Method: "POST", Path: "/get-link", Handle: ic.getInviteLink},
		{Method: "GET", Path: "/qr-code", Handle: ic.qrCode},
		{Method: "GET", Path: "/info", Handle: ic.inviteInfo},
		{Method: "POST", Path: "/register", Handle: ic.inviteRegister},
	}
}

// Invite 邀请注册
func (ic *InviteController) getInviteLink(context *gin.Context) {
	// 校验参数
	var i = &login.InviteDto{}
	if err := context.ShouldBindJSON(i); err != nil {
		ic.l.Error("参数绑定失败:" + err.Error())
		context.Error(err)
		return
	}
	// 获取当前用户ID
	uid, err := ic.inviteService.InviteService(context.GetInt("userId"), i)
	if err != nil {
		ic.l.Error("邀请失败:" + err.Error())
		context.Error(err)
		return
	}
	context.JSON(200, common.NewSuccessResult(map[string]string{"uid": uid}))
}

// InviteInfo 根据邀请的uuid获取邀请信息
func (ic *InviteController) inviteInfo(context *gin.Context) {
	// 校验参数
	uid := context.Query("uid")
	if uid == "" {
		ic.l.Error("uid 为空")
		context.Error(common.BadRequestError)
		return
	}
	// 获取邀请信息
	info, err := ic.inviteService.CheckInviteInfoIsValid(uid)
	if err != nil {
		ic.l.Error("获取邀请信息失败:" + err.Error())
		context.Error(err)
		return
	}
	context.JSON(200, common.NewSuccessResult(login2.NewInviteVo(info)))
}

// QrCode 获取二维码
func (ic *InviteController) qrCode(context *gin.Context) {
	uid := context.Query("uid")
	link, err := ic.inviteService.CheckInviteInfoIsValid(uid)
	if err != nil {
		ic.l.Error("邀请uid错误:" + err.Error())
		context.Error(err)
		return
	}
	if link.Id == 0 {
		ic.l.Error("邀请链接不存在")
		context.Error(common.InviteLinkNotFoundError)
		return
	}
	if link.IsUsed {
		ic.l.Error("邀请链接已被使用")
		context.Error(common.InviteLinkUsedError)
		return
	}
	// 获取二维码
	qrCode, err := utils.GenerateQRCode(fmt.Sprintf("http://%s/Invite-register?uid=%s", ic.c.Address.Domain, uid), 100)
	if err != nil {
		ic.l.Error("获取二维码失败:" + err.Error())
		context.Error(err)
		return
	}
	context.Header("Content-Type", "image/png")
	context.Writer.Write(qrCode)
}

// InviteRegister 使用邀请链接去注册(直接成为管理员)
func (ic *InviteController) inviteRegister(context *gin.Context) {
	// 校验参数
	var i = &login.InviteRegisterDto{}
	if err := context.ShouldBindJSON(i); err != nil {
		ic.l.Error("参数绑定失败:" + err.Error())
		context.Error(err)
		return
	}

	// 校验手机号是否一致
	if err := ic.loginService.ValidatePhone(i.Username); err != nil {
		ic.l.Error("手机号格式错误")
		context.Error(common.PhoneFormatError)
		return
	}

	err := ic.inviteService.InviteRegisterService(i)
	if err != nil {
		ic.l.Error("邀请注册失败:" + err.Error())
		context.Error(err)
		return
	}

	context.JSON(200, common.NewSuccessResult(nil))
}
func (ic *InviteController) RegisterController() {
	ic.cm.AddController(ic)
}
