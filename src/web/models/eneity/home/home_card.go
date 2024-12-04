package home

import "time"

type HomeCard struct {
	Id          int       `gorm:"primary_key;auto_increment;comment:'主键'"`
	Title       string    `gorm:"comment:'卡片的标题'"`
	Description *string   `gorm:"comment:'卡片的描述"`
	Image       *string   `gorm:"comment:'卡片的图片链接"`
	Path        string    `gorm:"comment:'卡片指向的路径"`
	CreatedAt   time.Time `gorm:"comment:'卡片的创建时间"`
	UpdatedAt   time.Time `gorm:"comment:'卡片的更新时间"`
}

func (h *HomeCard) TableName() string {
	return "home_card"
}
