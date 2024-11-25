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

func (l *LoginService) Login(loginUser *dto.UserDto) error {
	user := entity.NewUser(loginUser.Username, loginUser.Password)
	l.gorm.GetDb().Create(user)
	return nil
}
