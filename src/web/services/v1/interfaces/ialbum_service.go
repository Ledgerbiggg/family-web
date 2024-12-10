package interfaces

import (
	"family-web-server/src/web/models/eneity/login"
	"family-web-server/src/web/models/vo/album"
)

type IAlbumService interface {
	//
	// GetCategoryList
	//  @Description: 获取所有分类
	//  @return []*album.CategoryVo 分类
	//
	GetCategoryList(role *login.Role) []*album.CategoryVo

	//
	// GetImageBytesByName
	//  @Description: 根据图片名称获取图片
	//  @param name 图片名称
	//  @return []byte 图片字节
	//  @return error 错误
	//
	GetImageBytesByName(name string) ([]byte, error)

	//
	// GetCategoryPhotos
	//  @Description: 根据分类名称获取照片
	//  @param category 分类名称
	//  @return []*album.PhotoVo 照片
	//
	GetCategoryPhotos(category string, role *login.Role) []*album.PhotoVo
}
