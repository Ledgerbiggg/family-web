package login

type Role struct {
	Id          int     `gorm:"primary_key;comment:'主键'"`
	Name        string  `gorm:"comment:'角色名称'"`
	Description *string `gorm:"comment:'角色描述'"`
}

func (r *Role) TableName() string {
	return "role"
}
