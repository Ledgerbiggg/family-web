package tag

type TagRelation struct {
	Id       int `json:"id"`
	TagId    int `json:"tag_id"`
	LinkId   int `json:"link_id"`
	LinkType int `json:"link_type"`
	Sort     int `json:"sort"`
}

func (t *TagRelation) TableName() string {
	return "tag_relation"
}
