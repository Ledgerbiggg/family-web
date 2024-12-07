package impls

import (
	"family-web-server/src/config"
	"family-web-server/src/pkg/mysql"
	"family-web-server/src/web/common"
	albumVo "family-web-server/src/web/models/vo/album"
	"family-web-server/src/web/services/interfaces"
	"fmt"
	"os"
)

type AlbumService struct {
	c    *config.GConfig
	gorm *mysql.GormDb
}

func NewAlbumService(
	cf *config.GConfig,
	gorm *mysql.GormDb,
) interfaces.IAlbumService {
	return &AlbumService{c: cf, gorm: gorm}
}

func (a *AlbumService) GetCategoryList() []*albumVo.CategoryVo {
	var categoryVos []*albumVo.CategoryVo
	// 查询全部的相册分类+对应的封面
	a.gorm.GetDb().Raw(`
		SELECT ac.id,
			   ac.name,
			   ac.description,
			   ac.enabled,
			   ac.sort,
			   ac.view_count,
			   ac.status,
			   ac.created_by,
			   ac.created_at,
			   ac.updated_at,
			   CONCAT(ap.name,'.', ap.format) AS cover_pic
		FROM album_category ac
				 LEFT JOIN album_photo ap ON ac.cover = ap.id;
	`).Find(&categoryVos)
	return categoryVos
}

func (a *AlbumService) GetImageBytesByName(name string) ([]byte, error) {
	// 假设根目录下有 src/static/img 目录存储图片
	imagePath := fmt.Sprintf(a.c.Static.Dir+"img/%s", name)
	file, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, common.NotFoundResourceError
	}
	return file, nil
}
