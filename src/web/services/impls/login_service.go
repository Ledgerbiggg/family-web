package impls

import (
	"family-web-server/src/pkg/mysql"
	dto "family-web-server/src/web/models/dto/login"
	entity "family-web-server/src/web/models/eneity/login"
	"family-web-server/src/web/services/interfaces"
)

type LoginService struct {
	gorm *mysql.GormDb
}

func NewLoginService(gorm *mysql.GormDb) interfaces.ILoginService {
	return &LoginService{gorm: gorm}
}

func (l *LoginService) Login(loginUser *dto.UserDto) (bool, error) {
	var user1 entity.User
	user := entity.NewUser(loginUser.Username, loginUser.Password)
	l.gorm.GetDb().Where("username = ?", loginUser.Username).Find(&user1)
	if user1.Password != user.Password {
		return false, nil
	}
	return true, nil
}

func (l *LoginService) Register(loginUser *dto.UserDto) error {
	user := entity.NewUser(loginUser.Username, loginUser.Password)
	l.gorm.GetDb().Create(user)
	return nil
}
