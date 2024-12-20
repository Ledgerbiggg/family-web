package impls

import (
	"family-web-server/src/config"
	"family-web-server/src/pkg/mysql"
	"family-web-server/src/web/common"
	"family-web-server/src/web/models/eneity/album"
	"family-web-server/src/web/models/eneity/login"
	albumVo "family-web-server/src/web/models/vo/album"
	"family-web-server/src/web/services/v1/interfaces"
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

func (a *AlbumService) GetCategoryPhotos(category string, role *login.Role) []*albumVo.PhotoVo {
	var photoVos []*albumVo.PhotoVo
	a.gorm.GetDb().Raw(`
		SELECT ap.id,
			   ap.name,
			   ap.description,
			   ap.sort,
			   ap.is_lock,
			   ap.format,
			   ap.category_id,
			   u.nickname,
			   ap.upload_at
		FROM album_photo ap
				 LEFT JOIN album_category ac ON ap.category_id = ac.id
				 LEFT JOIN album_category_role acr ON ac.id = acr.category_id
				 LEFT JOIN user u ON ap.upload_by = u.id
		WHERE ap.category_id = ?
		AND acr.role_id = ?;
	`, category, role.Id).Find(&photoVos)
	for i := range photoVos {
		photoVos[i].UploadTime = photoVos[i].UploadAt.Format("2006-01-02 15:04:05")
	}
	return photoVos
}

func (a *AlbumService) GetCategoryList(role *login.Role) []*albumVo.CategoryVo {
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
			   CONCAT(ap.name, '.', ap.format) AS cover_pic
		FROM album_category ac
				 LEFT JOIN album_photo ap ON ac.cover = ap.id
				 LEFT JOIN album_category_role acr ON ac.id = acr.category_id
		WHERE acr.role_id = ?;
	`, role.Id).Find(&categoryVos)
	// 转换时间
	for i := range categoryVos {
		categoryVos[i].CreatedTime = categoryVos[i].CreatedAt.Format("2006-01-02 15:04:05")
		categoryVos[i].UpdatedTime = categoryVos[i].UpdatedAt.Format("2006-01-02 15:04:05")
	}
	return categoryVos
}

func (a *AlbumService) GetImageBytesByName(pid string) ([]byte, error) {
	var photo *album.Photo
	a.gorm.GetDb().Raw(`
		SELECT ap.id,
			   ap.name,
			   ap.description,
			   ap.sort,
			   ap.is_lock,
			   ap.format,
			   ap.category_id,
			   ap.upload_by,
			   ap.upload_at
		FROM album_photo ap
		WHERE ap.id = ?;
	`, pid).Scan(&photo)
	if photo.ID == 0 {
		return nil, common.NotFoundResourceError
	}
	// 假设根目录下有 src/static/img 目录存储图片
	imagePath := fmt.Sprintf(a.c.Static.Dir+"img/%s", photo.Name+"."+photo.Format)
	file, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, common.NotFoundResourceError
	}
	return file, nil
}
