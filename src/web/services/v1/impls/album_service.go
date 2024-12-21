package impls

import (
	"errors"
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/pkg/mysql"
	"family-web-server/src/web/common"
	"family-web-server/src/web/models/eneity/album"
	"family-web-server/src/web/models/eneity/login"
	albumPo "family-web-server/src/web/models/po/album"
	albumVo "family-web-server/src/web/models/vo/album"
	"family-web-server/src/web/services/v1/interfaces"
	"fmt"
	"gorm.io/gorm"
	"os"
)

type AlbumService struct {
	c    *config.GConfig
	gorm *mysql.GormDb
	l    *log.ConsoleLogger
}

func NewAlbumService(
	cf *config.GConfig,
	gorm *mysql.GormDb,
	l *log.ConsoleLogger,
) interfaces.IAlbumService {
	return &AlbumService{c: cf, gorm: gorm, l: l}
}
func (a *AlbumService) SaveCategoryByCategoryName(categoryName string) int {
	var category album.Category

	// 检查数据库中是否已有相同名称的记录
	db := a.gorm.GetDb()
	result := db.Where("name = ?", categoryName).First(&category)

	if result.Error == nil {
		// 已存在，返回现有记录的 ID
		return category.Id
	}

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 发生其他错误
		panic(result.Error)
	}

	// 如果不存在，插入新记录
	newCategory := &album.Category{Name: categoryName}
	db.Save(newCategory)

	return newCategory.Id
}
func (a *AlbumService) SavePhotoByCategoryIdAndPhotoName(categoryId int, photoName string) {
	var photo album.Photo

	// 检查是否已存在相同 CategoryID 和 Name 的记录
	db := a.gorm.GetDb()
	result := db.Where("category_id = ? AND name = ?", categoryId, photoName).First(&photo)

	if result.Error == nil {
		a.l.Error("相同 CategoryID 和 Name 的记录已存在")
		return
	}

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		a.l.Error(result.Error.Error())
		return
	}

	// 如果不存在，插入新记录
	newPhoto := &album.Photo{Name: photoName, CategoryID: categoryId}
	db.Save(newPhoto)
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
			   ac.cover,
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

func (a *AlbumService) GetImageBytesByName(_, pid string) ([]byte, error) {
	var photoPo *albumPo.PhotoPo
	a.gorm.GetDb().Raw(`
		SELECT ap.id,
			   ap.name,
			   ap.description,
			   ac.name category_name,
			   ap.sort,
			   ap.is_lock,
			   ap.format,
			   ap.category_id,
			   ap.upload_by,
			   ap.upload_at
		FROM album_photo ap
		LEFT JOIN album_category ac ON ap.category_id = ac.id
		WHERE ap.id = ?;
	`, pid).Scan(&photoPo)
	if photoPo.ID == 0 {
		return nil, common.NotFoundResourceError
	}
	// 假设根目录下有 src/static/img 目录存储图片
	imagePath := fmt.Sprintf(a.c.Static.Dir+"img/%s/%s", photoPo.CategoryName, photoPo.Name+"."+photoPo.Format)
	file, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, common.NotFoundResourceError
	}
	return file, nil
}
