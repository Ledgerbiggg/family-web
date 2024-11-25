package login

import "time"

type User struct {
	Id            int        `gorm:"primary_key;comment:'主键'"`
	Username      string     `gorm:"comment:'用户名(手机号)'"`
	Password      string     `gorm:"comment:'密码'"`
	Nickname      *string    `gorm:"comment:'昵称'"`
	IsDisabled    *bool      `gorm:"default:false;comment:'是否被禁用，默认启用'"`
	RegisterTime  time.Time  `gorm:"comment:'注册时间'"`
	LastLoginTime *time.Time `gorm:"comment:'最后登录时间'"`
	RealName      *string    `gorm:"comment:'真实姓名'"`
	Avatar        *string    `gorm:"comment:'头像'"`
	Email         *string    `gorm:"comment:'邮箱'"`
	RoleId        int        `gorm:"comment:'角色(关联角色表的id)'"`
}

func NewUser(username string, password string) *User {
	user := &User{Username: username, Password: password}
	// 默认角色是游客
	user.RoleId = 3
	user.RegisterTime = time.Now()
	return user
}

func (u *User) TableName() string {
	return "user"
}
