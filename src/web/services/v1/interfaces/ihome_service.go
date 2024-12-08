package interfaces

import (
	"family-web-server/src/web/models/eneity/login"
	homeVo "family-web-server/src/web/models/vo/home"
)

type IHomeService interface {
	//
	// GetHomeCardData
	//  @Description: 获取主页的卡片
	//  @param role 用户角色
	//  @return []*entity.HomeCard 返回这个角色可以访问的卡片
	//
	GetHomeCardData(role *login.Role) []*homeVo.HomeCardVo
}
