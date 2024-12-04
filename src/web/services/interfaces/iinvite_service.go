package interfaces

import (
	"family-web-server/src/web/models/dto/login"
	"family-web-server/src/web/models/eneity/invite"
)

type IInviteService interface {
	//
	// InviteService
	//  @Description: 邀请服务
	//  @param fromUsername 邀请者的手机号
	//  @param inviteDto 被邀请r人的信息
	//  @return string 邀请码
	//  @return error 错误
	//
	InviteService(fromUsername string, inviteDto *login.InviteDto) (string, error)
	//
	// CheckInviteInfoIsValid
	//  @Description: 检查邀请码是否有效
	//  @param uuid 邀请码唯一标识
	//  @return *entity.InviteLink 邀请链接数据
	//  @return error 错误
	//
	CheckInviteInfoIsValid(uuid string) (*invite.InviteLink, error)
	//
	// InviteRegisterService
	//  @Description: 使用邀请链接进行注册
	//  @param *Invite.InviteRegisterDto 前端的用户信息
	//  @return error 错误
	//
	InviteRegisterService(*login.InviteRegisterDto) error
}
