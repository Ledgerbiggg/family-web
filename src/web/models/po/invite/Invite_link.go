package invite

import "time"

type InviteLinkPo struct {
	Id              int        `gorm:"primary_key;comment:'主键'"`
	Uuid            string     `gorm:"comment:'邀请链接的唯一标识'"`
	IsUsed          bool       `gorm:"comment:'链接是否已经使用'"`
	Description     *string    `gorm:"comment:'邀请链接描述'"`
	InviterId       int        `gorm:"comment:'邀请人id'"`
	InviterPhone    string     `gorm:"comment:'邀请人手机号'"`
	InviterRealName string     `gorm:"comment:'邀请人真实姓名'"`
	InvitedRealName string     `gorm:"comment:'被邀请人真实姓名'"`
	InvitedAdmin    bool       `gorm:"comment:'被邀请人角色是否是admin'"`
	ExpirationDate  time.Time  `gorm:"comment:'邀请链接过期时间'"`
	CreatedAt       time.Time  `gorm:"comment:'创建时间'"`
	UsedAt          *time.Time `gorm:"comment:'使用时间'"`
}
