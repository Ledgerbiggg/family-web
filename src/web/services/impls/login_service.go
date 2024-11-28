package impls

import (
	"crypto/md5"
	"family-web-server/src/pkg/mysql"
	dto "family-web-server/src/web/models/dto/login"
	entity "family-web-server/src/web/models/eneity/login"
	"family-web-server/src/web/services/interfaces"
	"family-web-server/src/web/utils"
)

type LoginService struct {
	gorm *mysql.GormDb
}

func NewLoginService(gorm *mysql.GormDb) interfaces.ILoginService {
	return &LoginService{gorm: gorm}
}

// Login 查询数据库用户数据并比较密码
func (l *LoginService) Login(loginUser *dto.UserDto) (bool, *entity.Role, []*entity.Permission, error) {
	var user1 entity.User
	var role entity.Role
	var permissions []*entity.Permission
	l.gorm.GetDb().Where("username = ?", loginUser.Username).Find(&user1)
	md5.New().Sum([]byte(loginUser.Password))
	if user1.Password != utils.Md5Encrypt(loginUser.Password) {
		return false, nil, permissions, nil
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
	return true, nil, permissions, nil
}

func (l *LoginService) Register(loginUser *dto.UserDto) error {
	user := entity.NewUser(loginUser.Username, loginUser.Password)
	l.gorm.GetDb().Create(user)
	return nil
}
