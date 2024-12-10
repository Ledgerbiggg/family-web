package invite

import (
	invitePo "family-web-server/src/web/models/po/invite"
	"time"
)

type InviteLink struct {
	Id              int        `gorm:"primary_key;comment:'主键'"`
	Uuid            string     `gorm:"comment:'邀请链接的唯一标识'"`
	IsUsed          bool       `gorm:"comment:'链接是否已经使用'"`
	Description     *string    `gorm:"comment:'邀请链接描述'"`
	InviterId       int        `gorm:"comment:'邀请人id'"`
	InvitedRealName string     `gorm:"comment:'被邀请人真实姓名'"`
	InvitedAdmin    bool       `gorm:"comment:'被邀请人角色是否是admin'"`
	ExpirationDate  time.Time  `gorm:"comment:'邀请链接过期时间'"`
	CreatedAt       time.Time  `gorm:"comment:'创建时间'"`
	UsedAt          *time.Time `gorm:"comment:'使用时间'"`
}

func NewInviteLink(link invitePo.InviteLinkPo) *InviteLink {
	return &InviteLink{
		Id:              link.Id,
		Uuid:            link.Uuid,
		IsUsed:          link.IsUsed,
		Description:     link.Description,
		InviterId:       link.InviterId,
		InvitedRealName: link.InvitedRealName,
		InvitedAdmin:    link.InvitedAdmin,
		ExpirationDate:  link.ExpirationDate,
		CreatedAt:       link.CreatedAt,
		UsedAt:          link.UsedAt,
	}
}

func (i *InviteLink) TableName() string {
	return "invite_link"
}
