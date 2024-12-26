package album

import "time"

type PhotoVo struct {
	ID          int64      `json:"id"`          // 照片ID
	Name        string     `json:"name"`        // 照片名称
	Description *string    `json:"description"` // 照片描述
	Sort        int        `json:"sort"`        // 照片排序
	IsLock      bool       `json:"isLock"`      // 是否锁定
	Format      string     `json:"format"`      // 照片格式（如JPEG、PNG等）
	CategoryID  int        `json:"categoryID"`  // 相册ID
	Nickname    string     `json:"nickname"`    // 上传用户
	CreatedAt   *time.Time `json:"-"`           // 上传时间
	CreatedTime string     `json:"createdTime"` // 上传时间
}
