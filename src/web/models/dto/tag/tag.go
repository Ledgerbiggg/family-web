package tag

type TagDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name" validate:"required"`             // 必填
	Description string `json:"description"`                          // 可选
	Type        int    `json:"type" validate:"required,min=1,max=5"` // 必填且必须在 0-3 之间
}
