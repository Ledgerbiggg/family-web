package login

type InviteDto struct {
	// 邀请人的真实姓名
	RealName string `json:"realName" binding:"required"`
	// 邀请描述
	Description string `json:"description" binding:"required"`
	// 是否邀请成为管理员
	InvitedAdmin bool `json:"invitedAdmin" binding:"required"`
}
