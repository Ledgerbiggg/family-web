SELECT *
FROM `user`
WHERE username = '18288888888';



INSERT INTO `user` (`username`, `password`, `nickname`, `is_disabled`, `register_time`, `last_login_time`, `real_name`,
                    `avatar`, `email`, `role_id`)
VALUES ('', 'E10ADC3949BA59ABBE56E057F20F883E', NULL, FALSE, '0000-00-00 00:00:00', NULL, NULL, NULL, NULL, 2);


# 查询角色和主页卡片的对应关系
SELECT hc.*
FROM role r
         LEFT JOIN role_home_card_access rhca ON r.id = rhca.role_id
         LEFT JOIN home_card hc ON rhca.home_card_id = hc.id
WHERE r.id = 1;


# 	Id              int        `gorm:"primary_key;comment:'主键'"`
# 	Uuid            string     `gorm:"comment:'邀请链接的唯一标识'"`
# 	IsUsed          bool       `gorm:"comment:'链接是否已经使用'"`
# 	Description     *string    `gorm:"comment:'邀请链接描述'"`
# 	InviterId       string     `gorm:"comment:'邀请人id'"`
# 	InviterPhone    string     `gorm:"comment:'邀请人手机号'"`
# 	InviterRealName string     `gorm:"comment:'邀请人真实姓名'"`
# 	InvitedRealName string     `gorm:"comment:'被邀请人真实姓名'"`
# 	InvitedAdmin    bool       `gorm:"comment:'被邀请人角色是否是admin'"`
# 	ExpirationDate  time.Time  `gorm:"comment:'邀请链接过期时间'"`
# 	CreatedAt       time.Time  `gorm:"comment:'创建时间'"`
# 	UsedAt          *time.Time `gorm:"comment:'使用时间'"`
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

SELECT ac.id,
       ac.name,
       ac.description,
       ac.enabled,
       ac.sort,
       ac.view_count,
       ac.status,
       ac.created_by,
       ac.created_at,
       ac.updated_at,
       CONCAT(ap.name, '.', ap.format) AS cover
FROM album_category ac
         LEFT JOIN album_photo ap ON ac.cover = ap.id;


# //	ID          int64   `json:"id"`          // 照片ID
# 	//	Name        string  `json:"name"`        // 照片名称
# 	//	Description *string `json:"description"` // 照片描述
# 	//	Sort        int     `json:"sort"`        // 照片排序
# 	//	IsLock      bool    `json:"isLock"`      // 是否锁定
# 	//	Format      string  `json:"format"`      // 照片格式（如JPEG、PNG等）
# 	//	CategoryID  int     `json:"categoryID"`  // 相册ID
# 	//	UploadBy    int     `json:"uploadBy"`    // 上传用户
# 	//	UploadTime  string  `json:"uploadTime"`  // 上传时间

SELECT ap.id,
       ap.name,
       ap.description,
       ap.sort,
       ap.is_lock,
       ap.format,
       ap.category_id,
       u.nickname,
       ap.upload_at
FROM album_photo ap
         LEFT JOIN album_category ac ON ap.category_id = ac.id
         LEFT JOIN album_category_role acr ON ac.id = acr.category_id
         LEFT JOIN user u ON ap.upload_by = u.id
WHERE ap.category_id = ?
  AND acr.role_id = ?;



SELECT r.id, r.name, r.description
FROM user u
         LEFT JOIN role r ON u.role_id = r.id
WHERE u.id = ?;

SELECT p.id, p.path, p.description
FROM user u
         LEFT JOIN role r ON u.role_id = r.id
         LEFT JOIN role_permission rp ON r.id = rp.role_id
         LEFT JOIN permission p ON rp.permission_id = p.id
WHERE u.id = ?;


SELECT ac.id,
       ac.name,
       ac.description,
       ac.enabled,
       ac.sort,
       ac.view_count,
       ac.status,
       ac.created_by,
       ac.created_at,
       ac.updated_at,
       CONCAT(ap.name, '.', ap.format) AS cover_pic
FROM album_category ac
         LEFT JOIN album_photo ap ON ac.cover = ap.id
         LEFT JOIN album_category_role acr ON ac.id = acr.category_id
WHERE acr.role_id = ?;



SELECT ap.id,
       ap.name,
       ap.description,
       ap.sort,
       ap.is_lock,
       ap.format,
       ap.category_id,
       ap.upload_by,
       ap.upload_at
FROM album_photo ap
WHERE ap.id = ?;




















































































