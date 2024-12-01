package login

type InviteRegisterDto struct {
	InviteUid string `json:"inviteUid" binding:"required"`
	Username  string `json:"username" binding:"required"`
	RealName  string `json:"realName" binding:"required"`
}
