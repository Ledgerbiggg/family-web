package impls

import (
	"crypto/md5"
	"family-web-server/src/log"
	"family-web-server/src/pkg/mysql"
	"family-web-server/src/web/common"
	dto "family-web-server/src/web/models/dto/login"
	entity "family-web-server/src/web/models/eneity/login"
	"family-web-server/src/web/services/interfaces"
	"family-web-server/src/web/utils"
	"github.com/steambap/captcha"
	"regexp"
	"strings"
)

type LoginService struct {
	gorm *mysql.GormDb
	l    *log.ConsoleLogger
}

func (l *LoginService) VerifyService(verifyDto *dto.VerifyDto) error {
	var u entity.User
	db := l.gorm.GetDb()
	db.Where("username = ? and real_name = ?", verifyDto.Username, verifyDto.ReaName).Find(&u)
	if u.Username == "" {
		return common.RealNameNotMatchError
	}
	// 查看是否是管理员角色,不是的话返回错误
	if u.RoleId != 1 && u.RoleId != 2 {
		return common.AdminRoleError
	}
	// 修改密码为 md5 加密后的密码
	u.Password = utils.Md5Encrypt("123456")
	db.Save(&u)
	return nil
}

func (l *LoginService) ValidatePhone(phone string) error {
	// 正则表达式用于验证手机号格式
	phoneRegex := `^1[3-9]\d{9}$`
	matched, err := regexp.MatchString(phoneRegex, phone)
	if err != nil {
		// 如果正则匹配出现错误，返回相应的错误
		return err
	}
	if !matched {
		// 如果手机号格式不匹配，返回自定义的错误
		return common.PhoneFormatError
	}
	return nil
}

// RegisterService  注册用户
func (l *LoginService) RegisterService(registerDto *dto.RegisterDto) error {
	var u entity.User
	// 先检查数据库是否有一样的用户名(手机号存在)
	db := l.gorm.GetDb()
	db.Where("username = ?", registerDto.Username).Find(&u)
	registerUser := entity.NewUser(registerDto.Username, utils.Md5Encrypt(registerDto.Password))
	if u.Username == "" {
		db.Create(registerUser)
	} else {
		return common.UserIsExistError
	}
	return nil
}

// ValidateCaptcha 验证验证码
func (l *LoginService) ValidateCaptcha(storedCaptcha, inputCaptcha string) error {
	if storedCaptcha == "" || strings.ToLower(storedCaptcha) != strings.ToLower(inputCaptcha) {
		return common.CaptchaErrorError
	}
	return nil
}

// CaptchaService 获取验证码图像
func (l *LoginService) CaptchaService() (*captcha.Data, error) {
	return captcha.New(200, 100)
}

// LoginService  查询数据库用户数据并比较密码
func (l *LoginService) LoginService(loginUser *dto.UserDto) (int, *entity.Role, []*entity.Permission, error) {
	var user1 entity.User
	var role entity.Role
	var permissions []*entity.Permission
	l.gorm.GetDb().Where("username = ?", loginUser.Username).Find(&user1)
	md5.New().Sum([]byte(loginUser.Password))
	if user1.Password != utils.Md5Encrypt(loginUser.Password) {
		return 0, nil, permissions, common.LoginErrorError
	}
	// 查询用户角色+用户权限
	l.gorm.GetDb().Where("id = ?", user1.RoleId).Find(&role)

	l.gorm.GetDb().Raw(`
		SELECT p.id, p.path, p.description
		FROM user u
				 LEFT JOIN role r ON u.role_id = r.id
				 LEFT JOIN role_permission rp ON r.id = rp.role_id
				 LEFT JOIN permission p ON rp.permission_id = p.id
		WHERE u.username = ?`, loginUser.Username,
	).Scan(&permissions)
	return user1.Id, &role, permissions, nil
}
func NewLoginService(
	gorm *mysql.GormDb,
	l *log.ConsoleLogger,
) interfaces.ILoginService {
	return &LoginService{gorm: gorm, l: l}
}
