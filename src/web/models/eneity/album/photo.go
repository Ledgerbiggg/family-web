package album

import "time"

// Photo 相册照片
type Photo struct {
	ID          int64      `json:"id"`          // 照片ID
	Name        string     `json:"name"`        // 照片名称
	Description *string    `json:"description"` // 照片描述
	Sort        int        `json:"sort"`        // 照片排序
	IsLock      bool       `json:"isLock"`      // 是否锁定
	Format      string     `json:"format"`      // 照片格式（如JPEG、PNG等）
	CategoryID  int        `json:"categoryID"`  // 相册ID
	UploadBy    int        `json:"uploadBy"`    // 上传用户
	UploadTime  *time.Time `json:"uploadTime"`  // 上传时间
}

func (p *Photo) TableName() string {
	return "album_photo"
}
