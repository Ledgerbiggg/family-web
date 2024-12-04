package interfaces

import (
	"family-web-server/src/web/models/dto/login"
	entity "family-web-server/src/web/models/eneity/login"
	"github.com/steambap/captcha"
)

type ILoginService interface {
	//
	// CaptchaService
	//  @Description: 生成验证码
	//  @return *captcha.Data 验证码数据
	//  @return error错误
	//
	CaptchaService() (*captcha.Data, error)
	//
	// ValidatePhone
	//  @Description: 验证手机号
	//  @param phone 手机号
	//  @return error 错误
	//
	ValidatePhone(phone string) error
	//
	// LoginService
	//  @Description: 登录服务
	//  @param *login.UserDto 前端的用户信息
	//  @return bool 是否登录成功
	//  @return *entity.Role 角色
	//  @return []*entity.Permission 权限
	//  @return error 错误
	//
	LoginService(*login.UserDto) (bool, *entity.Role, []*entity.Permission, error)
	//
	// RegisterService
	//  @Description: 注册服务
	//  @param *login.RegisterDto  前端的用户信息
	//  @return error 错误
	//
	RegisterService(*login.RegisterDto) error
	//
	// VerifyService
	//  @Description: 找回密码服务(如果真实姓名和手机号一致就改密码为123456)
	//  @param *login.VerifyDto  前端的用户信息
	//  @return error 错误
	//
	VerifyService(*login.VerifyDto) error
	//
	// InviteService
	//  @Description: 邀请服务
	//  @param fromUsername 邀请者的手机号
	//  @param inviteDto 被邀请r人的信息
	//  @return string 邀请码
	//  @return error 错误
	//
	InviteService(fromUsername string, inviteDto *login.InviteDto) (string, error)
	//
	// CheckInviteInfoIsValid
	//  @Description: 检查邀请码是否有效
	//  @param uuid 邀请码唯一标识
	//  @return *entity.InviteLink 邀请链接数据
	//  @return error 错误
	//
	CheckInviteInfoIsValid(uuid string) (*entity.InviteLink, error)
	//
	// InviteRegisterService
	//  @Description: 使用邀请链接进行注册
	//  @param *login.InviteRegisterDto 前端的用户信息
	//  @return error 错误
	//
	InviteRegisterService(*login.InviteRegisterDto) error
}
