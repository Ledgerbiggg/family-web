package album

import "time"

type CategoryVo struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	CoverPic    string     `json:"coverPic"`
	Description *string    `json:"description,omitempty"`
	ViewCount   int        `json:"viewCount"`
	Status      string     `json:"status"`
	CreatedBy   int        `json:"createdBy"`
	CreatedAt   time.Time  `json:"-"`
	CreatedTime string     `json:"createdTime"`
	UpdatedAt   *time.Time `json:"-"`
	UpdatedTime string     `json:"updatedTime"`
}
