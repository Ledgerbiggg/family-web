package interfaces

import "family-web-server/src/web/models/dto/login"

type ILoginService interface {
	Login(loginUser *login.UserDto) error
}
