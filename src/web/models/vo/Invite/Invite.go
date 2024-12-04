package Invite

import (
	entity "family-web-server/src/web/models/eneity/invite"
)

type InviteVo struct {
	Id              int     `json:"id"`              // 主键
	Uuid            string  `json:"uuid"`            // 邀请链接的唯一标识
	Description     *string `json:"description"`     // 邀请链接描述
	InviterUsername string  `json:"inviterUsername"` // 邀请人手机号
	InvitedAdmin    bool    `json:"invitedAdmin"`    // 被邀请人角色是否是 admin
}

func NewInviteVo(link *entity.InviteLink) *InviteVo {
	return &InviteVo{
		Id:              link.Id,
		Uuid:            link.Uuid,
		Description:     link.Description,
		InviterUsername: link.InviterUsername,
		InvitedAdmin:    link.InvitedAdmin,
	}
}
