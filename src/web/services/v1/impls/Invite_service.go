package impls

import (
	"family-web-server/src/log"
	"family-web-server/src/pkg/mysql"
	"family-web-server/src/web/common"
	dto "family-web-server/src/web/models/dto/login"
	"family-web-server/src/web/models/eneity/invite"
	entity "family-web-server/src/web/models/eneity/login"
	invitePo "family-web-server/src/web/models/po/invite"
	"family-web-server/src/web/services/v1/interfaces"
	"family-web-server/src/web/utils"
	"time"
)

type InviteService struct {
	gorm *mysql.GormDb
	l    *log.ConsoleLogger
}

func NewInviteService(gorm *mysql.GormDb, l *log.ConsoleLogger) interfaces.IInviteService {
	return &InviteService{gorm: gorm, l: l}
}

func (l *InviteService) CheckInviteInfoIsValid(uuid string) (*invitePo.InviteLinkPo, error) {
	var link invitePo.InviteLinkPo
	l.gorm.GetDb().Raw(`
			SELECT il.id                AS id,
				   il.uuid              AS uuid,
				   il.is_used           AS is_used,
				   il.description       AS description,
				   il.inviter_id        AS inviter_id,
				   u.username           AS inviter_phone,
				   u.real_name          AS inviter_real_name,
				   il.invited_real_name AS invited_real_name,
				   il.invited_admin     AS invited_admin,
				   il.expiration_date   AS expiration_date,
				   il.created_at        AS created_at,
				   il.used_at           AS used_at
			FROM invite_link il
					 LEFT JOIN user u ON il.inviter_id = u.id
			WHERE uuid = ?;
		`, uuid).Scan(&link)
	if link.Uuid == "" {
		return nil, common.NotFoundResourceError
	}
	// 判断是否被使用过
	if link.IsUsed {
		return nil, common.InviteLinkUsedError
	}
	// 判断邀请链接是否过期
	if time.Now().After(link.ExpirationDate) {
		return nil, common.InviteRegisterExpiredError
	}
	return &link, nil
}

func (l *InviteService) InviteService(fromUserId int, inviteDto *dto.InviteDto) (uid string, err error) {
	// 校验fromUsername的权限是否为管理员
	isAdmin, err := l.gorm.IsAdmin(fromUserId)
	if err != nil {
		l.l.Error("获取用户权限失败:" + err.Error())
		return "", common.DatabaseError
	}
	if !isAdmin {
		l.l.Error("用户权限不足,不是管理员")
		return "", common.AdminRoleError
	}
	// 用户是管理员 创建邀请链接
	uuid := utils.GetRandomId(20)
	link := invite.InviteLink{
		Uuid:            uuid,
		IsUsed:          false,
		Description:     &inviteDto.Description,
		InviterId:       fromUserId,
		InvitedRealName: inviteDto.RealName,
		InvitedAdmin:    inviteDto.InvitedAdmin,
		ExpirationDate:  time.Now().Add(time.Hour * 24),
		CreatedAt:       time.Now(),
	}
	l.gorm.GetDb().Create(&link)
	return uuid, nil
}

func (l *InviteService) InviteRegisterService(inviteRegisterDto *dto.InviteRegisterDto) error {
	db := l.gorm.GetDb()
	tx := db.Begin()
	// 确保在出错时回滚事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 如果发生任何panic，回滚事务
		}
	}()

	uid := inviteRegisterDto.InviteUid
	// 查询邀请链接
	link, err := l.CheckInviteInfoIsValid(uid)
	if err != nil {
		return err
	}
	// 查询用户名是否一致
	if link.InvitedRealName != inviteRegisterDto.RealName {
		return common.InviteRegisterError
	}
	// 检查用户是否存在
	var u entity.User
	db.Where("username = ?", inviteRegisterDto.Username).Find(&u)
	if u.Username != "" {
		return common.UserIsExistError
	}
	// 标记邀请链接已被使用
	link.IsUsed = true
	now := time.Now()
	link.UsedAt = &now
	// 保存更新后的链接
	if err = tx.Save(invite.NewInviteLink(*link)).Error; err != nil {
		tx.Rollback() // 出错时回滚事务
		return err
	}
	// 修改密码为 md5 加密后的密码
	u.Username = inviteRegisterDto.Username
	u.Password = utils.Md5Encrypt("123456")
	u.RealName = &inviteRegisterDto.RealName
	u.RegisterTime = now
	if link.InvitedAdmin {
		u.RoleId = 2
	}
	// 存入用户
	if err = tx.Save(&u).Error; err != nil {
		tx.Rollback() // 出错时回滚事务
		return err
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback() // 提交失败时回滚事务
		return err
	}

	return nil
}
