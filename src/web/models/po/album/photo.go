package album

import "time"

// PhotoPo 相册照片
type PhotoPo struct {
	ID           int64      `gorm:"primaryKey;comment:照片ID"`
	Name         string     `gorm:"comment:照片名称"`
	CategoryName string     `gorm:"comment:照片名称"`
	Description  *string    `gorm:"comment:照片描述"`
	Sort         int        `gorm:"comment:排序"`
	IsLock       bool       `gorm:"comment:是否锁定"`
	Format       string     `gorm:"comment:照片格式"`
	CategoryID   int        `gorm:"comment:相册ID"`
	UploadBy     int        `gorm:"comment:上传用户"`
	UploadAt     *time.Time `gorm:"comment:上传时间"`
}
