package album

import "time"

// Category 相册分类
type Category struct {
	Id          int        `gorm:"primary_key;comment:'自增主键'"`
	Name        string     `gorm:"comment:'分类名称'"`
	Cover       int        `gorm:"comment:'封面图片ID'"`
	Description *string    `gorm:"comment:'分类描述'"`
	Enabled     bool       `gorm:"comment:'是否启用'"`
	Sort        int        `gorm:"comment:'排序字段'"`
	ViewCount   int        `gorm:"comment:'视图计数'"`
	Status      string     `gorm:"comment:'分类状态'"`
	CreatedBy   int        `gorm:"comment:'创建者ID'"`
	CreatedAt   time.Time  `gorm:"comment:'创建时间'"`
	UpdatedAt   *time.Time `gorm:"comment:'更新时间'"`
}

func (c Category) TableName() string {
	return "album_category"
}
