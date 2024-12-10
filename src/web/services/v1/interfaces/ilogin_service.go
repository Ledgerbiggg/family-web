package interfaces

import (
	loginDto "family-web-server/src/web/models/dto/login"
	"family-web-server/src/web/models/eneity/login"
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
	// GetRoleAndPermissionByUserId
	//  @Description: 根据用户id获取角色和权限
	//  @param userId 用户id
	//  @return *login.Role 角色
	//  @return []*login.Permission 权限
	//  @return error
	//
	GetRoleAndPermissionByUserId(userId int) (*login.Role, []*login.Permission, error)
	//
	// LoginService
	//  @Description: 登录服务
	//  @param *Invite.UserDto 前端的用户信息
	//  @return bool 是否登录成功
	//  @return *entity.Role 角色
	//  @return []*entity.Permission 权限
	//  @return error 错误
	//
	LoginService(*loginDto.UserDto) (int, error)
	//
	// RegisterService
	//  @Description: 注册服务
	//  @param *Invite.RegisterDto  前端的用户信息
	//  @return error 错误
	//
	RegisterService(*loginDto.RegisterDto) error
	//
	// VerifyService
	//  @Description: 找回密码服务(如果真实姓名和手机号一致就改密码为123456)
	//  @param *Invite.VerifyDto  前端的用户信息
	//  @return error 错误
	//
	VerifyService(*loginDto.VerifyDto) error
}
