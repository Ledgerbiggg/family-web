package login

type RegisterDto struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
	Captcha         string `json:"captcha" binding:"required,len=4"`
}
