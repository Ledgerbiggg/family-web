package test5

import (
	"family-web-server/src/web/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlbumService(t *testing.T) {

}

func TestAlbumServiceGetImageBytesByName(t *testing.T) {
	name, err := albumService.GetImageBytesByCategoryIdAndPid("1")
	if err != nil {
		t.Error(err)
	}
	t.Log(name)
}
func TestAlbumServiceGetCategoryList(t *testing.T) {
	list := albumService.GetCategoryList(adminRole)
	assert.NotNil(t, list, "Expected list to be not nil")
}
func TestAlbumServiceSaveCategoryByCategoryName(t *testing.T) {
	categoryId, err := albumService.SaveCategoryByCategoryName("test1111")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(categoryId)
}

func TestAlbumServiceSavePhotoByCategoryIdAndPhotoName(t *testing.T) {
	err := albumService.SavePhotoByCategoryIdAndPhotoName(2, "test.png")
	if err != nil {
		t.Error(err.Error())
	}
}
func TestAlbumUtil(t *testing.T) {
	utils.ReadPathAllDir("./src/static/img/",
		albumService.SaveCategoryByCategoryName,
		albumService.SavePhotoByCategoryIdAndPhotoName)
}
