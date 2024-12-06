package impls

import (
	"family-web-server/src/log"
	"family-web-server/src/pkg/mysql"
	"family-web-server/src/web/models/eneity/home"
	"family-web-server/src/web/models/eneity/login"
	homeVo "family-web-server/src/web/models/vo/home"
	"family-web-server/src/web/services/interfaces"
)

type HomeService struct {
	gorm *mysql.GormDb
	l    *log.ConsoleLogger
}

func NewHomeService(
	gorm *mysql.GormDb,
	l *log.ConsoleLogger,
) interfaces.IHomeService {
	return &HomeService{gorm: gorm, l: l}
}

func (h *HomeService) GetHomeCardData(role *login.Role) []*homeVo.HomeCardVo {
	db := h.gorm.GetDb()
	var homeCards []*home.HomeCard
	db.Raw(`
		SELECT hc.*
			FROM role r
					 LEFT JOIN role_home_card_access rhca ON r.id = rhca.role_id
					 LEFT JOIN home_card hc ON rhca.home_card_id = hc.id
			WHERE r.id = ?
		ORDER BY hc.sort
		`, role.Id).Scan(&homeCards)
	var homeCardVos []*homeVo.HomeCardVo
	for i := range homeCards {
		homeCardVos = append(homeCardVos, homeVo.NewHomeCardVo(homeCards[i]))
	}
	return homeCardVos
}
