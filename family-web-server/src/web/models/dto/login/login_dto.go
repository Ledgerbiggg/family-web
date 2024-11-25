package login

type UserDto struct {
	Username string `json:"username" binding:"required"`      // 必填，最小3个字符，最大20个字符
	Password string `json:"password" binding:"required"`      // 必填，最小6个字符
	Captcha  string `json:"captcha" binding:"required,len=4"` // 必填，长度为4
}
