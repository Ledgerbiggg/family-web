package login

type Permission struct {
	Id          int     `gorm:"primary_key;auto_increment;comment:'主键'"`
	Path        string  `gorm:"comment:'权限路径'"`
	Description *string `gorm:"comment:'权限描述'"`
}

// TableName 设置表名
func (p *Permission) TableName() string {
	return "permission"
}
