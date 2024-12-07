package Invite

import (
	po "family-web-server/src/web/models/po/invite"
	"strings"
)

type InviteVo struct {
	Id              int     `json:"id"`              // 主键
	Uuid            string  `json:"uuid"`            // 邀请链接的唯一标识
	Description     *string `json:"description"`     // 邀请链接描述
	InviterRealName string  `json:"inviterRealName"` // 邀请人真实姓名
	InviterPhone    string  `json:"inviterPhone"`    // 邀请人手机号
	InvitedAdmin    bool    `json:"invitedAdmin"`    // 被邀请人角色是否是 admin
	ExpirationDate  string  `json:"expirationDate"`  // 邀请链接过期时间
}

func NewInviteVo(link *po.InviteLinkPo) *InviteVo {
	// 脱敏处理手机号
	maskedPhone := maskPhone(link.InviterPhone)
	// 脱敏处理真实姓名
	maskedRealName := maskRealName(link.InviterRealName)

	return &InviteVo{
		Id:              link.Id,
		Uuid:            link.Uuid,
		Description:     link.Description,
		InviterRealName: maskedRealName, // 使用脱敏后的真实姓名
		InviterPhone:    maskedPhone,    // 使用脱敏后的手机号
		InvitedAdmin:    link.InvitedAdmin,
		ExpirationDate:  link.ExpirationDate.Format("2006-01-02 15:04:05"),
	}
}

// maskPhone 脱敏手机号，中间四位替换为****
func maskPhone(phone string) string {
	if len(phone) < 7 {
		return phone // 如果手机号长度不足，返回原手机号
	}
	return phone[:3] + "****" + phone[7:] // 替换中间四位
}

// maskRealName 脱敏真实姓名，只保留最后一个字
func maskRealName(realName string) string {
	runes := []rune(realName) // 将字符串转换为rune切片
	if len(runes) < 2 {
		return realName // 如果姓名长度不足，返回原姓名
	}
	// 返回脱敏后的姓名，前面部分用*替换
	return "*" + strings.Repeat("*", len(runes)-2) + string(runes[len(runes)-1])
}
