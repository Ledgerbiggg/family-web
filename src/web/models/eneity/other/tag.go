package other

type Tag struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        int    `json:"type"`
}

func (t *Tag) TableName() string {
	return "tag"
}
