package interfaces

import (
	"family-web-server/src/web/models/dto/login"
	entity "family-web-server/src/web/models/eneity/login"
)

type ILoginService interface {
	Login(loginUser *login.UserDto) (bool, *entity.Role, []*entity.Permission, error)
}
