package album

type CategoryVo struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	CoverPic    string  `json:"coverPic"`
	Description *string `json:"description,omitempty"`
	ViewCount   int     `json:"viewCount"`
	Status      string  `json:"status"`
	CreatedBy   int     `json:"createdBy"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt,omitempty"`
}
