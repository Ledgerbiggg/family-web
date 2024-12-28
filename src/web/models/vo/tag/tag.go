package tag

import "time"

type TagVo struct {
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Type        int        `json:"type"`
	CreatedAt   time.Time  `json:"-"`
	CreatedTime string     `json:"createdTime"`
	UpdatedAt   *time.Time `json:"-"`
	UpdatedTime string     `json:"updatedTime"`
}
