package home

import "family-web-server/src/web/models/eneity/home"

type HomeCardVo struct {
	Id          int     `json:"id"`          // 主键
	Title       string  `json:"title"`       // 卡片的标题
	Description *string `json:"description"` // 卡片的描述
	Image       *string `json:"image"`       // 卡片的图片链接
	Path        string  `json:"path"`        // 卡片指向的路径
}

func NewHomeCardVo(card *home.HomeCard) *HomeCardVo {
	return &HomeCardVo{
		Id:          card.Id,
		Title:       card.Title,
		Description: card.Description,
		Image:       card.Image,
		Path:        card.Path,
	}
}
