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
	"time"
)

type LoginService struct {
	gorm *mysql.GormDb
	l    *log.ConsoleLogger
}

func (l *LoginService) CheckInviteInfoIsValid(uuid string) (*entity.InviteLink, error) {
	var link entity.InviteLink
	l.gorm.GetDb().Where("uuid = ?", uuid).Find(&link)
	if link.Uuid == "" {
		return nil, common.NotFoundResourceError
	}
	// 判断是否被使用过
	if link.IsUsed {
		return nil, common.InviteLinkUsedError
	}
	// 判断邀请链接是否过期
	if time.Now().After(link.ExpirationDate) {
		return nil, common.NotFoundResourceError
	}
	return &link, nil
}

func (l *LoginService) InviteService(fromUsername string, inviteDto *dto.InviteDto) (uid string, err error) {
	// 校验fromUsername的权限是否为管理员
	isAdmin, err := l.gorm.IsAdmin(fromUsername)
	if err != nil {
		l.l.Error("获取用户权限失败:" + err.Error())
		return "", common.DatabaseError
	}
	if !isAdmin {
		l.l.Error("用户权限不足,不是管理员")
		return "", common.AdminRoleError
	}
	// 用户是管理员 创建邀请链接
	uuid := utils.GetRandomId(20)
	link := entity.InviteLink{
		Uuid:            uuid,
		IsUsed:          false,
		Description:     &inviteDto.Description,
		InviterUsername: fromUsername,
		InvitedRealName: inviteDto.RealName,
		InvitedAdmin:    inviteDto.InvitedAdmin,
		ExpirationDate:  time.Now().Add(time.Hour * 24),
		CreatedAt:       time.Now(),
	}
	l.gorm.GetDb().Create(&link)
	return uuid, nil
}

func (l *LoginService) InviteRegisterService(inviteRegisterDto *dto.InviteRegisterDto) error {
	db := l.gorm.GetDb()
	tx := db.Begin()
	// 确保在出错时回滚事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 如果发生任何panic，回滚事务
		}
	}()

	uid := inviteRegisterDto.InviteUid
	// 查询邀请链接
	link, err := l.CheckInviteInfoIsValid(uid)
	if err != nil {
		return err
	}
	// 查询用户名是否一致
	if link.InvitedRealName != inviteRegisterDto.RealName {
		return common.InviteRegisterError
	}
	// 检查用户是否存在
	var u entity.User
	db.Where("username = ?", inviteRegisterDto.Username).Find(&u)
	if u.Username != "" {
		return common.UserIsExistError
	}
	// 标记邀请链接已被使用
	link.IsUsed = true
	now := time.Now()
	link.UsedAt = &now
	// 保存更新后的链接
	if err = tx.Save(link).Error; err != nil {
		tx.Rollback() // 出错时回滚事务
		return err
	}
	// 修改密码为 md5 加密后的密码
	u.Username = inviteRegisterDto.Username
	u.Password = utils.Md5Encrypt("123456")
	u.RealName = &inviteRegisterDto.RealName
	u.RegisterTime = now
	if link.InvitedAdmin {
		u.RoleId = 2
	}
	// 存入用户
	if err = tx.Save(&u).Error; err != nil {
		tx.Rollback() // 出错时回滚事务
		return err
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback() // 提交失败时回滚事务
		return err
	}

	return nil
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
func (l *LoginService) LoginService(loginUser *dto.UserDto) (bool, *entity.Role, []*entity.Permission, error) {
	var user1 entity.User
	var role entity.Role
	var permissions []*entity.Permission
	l.gorm.GetDb().Where("username = ?", loginUser.Username).Find(&user1)
	md5.New().Sum([]byte(loginUser.Password))
	if user1.Password != utils.Md5Encrypt(loginUser.Password) {
		return false, nil, permissions, common.LoginErrorError
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
	return true, &role, permissions, nil
}
func NewLoginService(
	gorm *mysql.GormDb,
	l *log.ConsoleLogger,
) interfaces.ILoginService {
	return &LoginService{gorm: gorm, l: l}
}
