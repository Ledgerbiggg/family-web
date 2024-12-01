package interfaces

import (
	"family-web-server/src/web/models/dto/login"
	entity "family-web-server/src/web/models/eneity/login"
	"github.com/steambap/captcha"
)

type ILoginService interface {
	// CaptchaService 获取验证码图像
	CaptchaService() (*captcha.Data, error)
	// ValidatePhone 校验手机号格式
	ValidatePhone(phone string) error
	// LoginService 查询数据库用户数据并比较密码
	LoginService(*login.UserDto) (bool, *entity.Role, []*entity.Permission, error)
	// RegisterService  注册
	RegisterService(*login.RegisterDto) error
	// VerifyService 找回密码
	VerifyService(*login.VerifyDto) error
	// InviteService 邀请一个成语进行注册,生成邀请码
	InviteService(fromUsername string, inviteDto *login.InviteDto) (string, error)
	// CheckInviteInfoIsValid 根据邀请的uuid获取邀请信息
	CheckInviteInfoIsValid(uuid string) (*entity.InviteLink, error)
	// InviteRegisterService 根据邀请码注册
	InviteRegisterService(*login.InviteRegisterDto) error
}
