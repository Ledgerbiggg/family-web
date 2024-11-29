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
	// ValidateCaptcha 验证验证码
	ValidateCaptcha(storedCaptcha, inputCaptcha string) error
	// LoginService 查询数据库用户数据并比较密码
	LoginService(registerDto *login.UserDto) (bool, *entity.Role, []*entity.Permission, error)
	// RegisterService  注册
	RegisterService(loginUser *login.RegisterDto) error
}
