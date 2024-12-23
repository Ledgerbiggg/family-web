package test5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlbumService(t *testing.T) {

}

func TestAlbumServiceGetImageBytesByName(t *testing.T) {
	name, err := albumService.GetImageBytesByName("1", "1")
	if err != nil {
		t.Error(err)
	}
	t.Log(name)
}
func TestAlbumServiceGetCategoryList(t *testing.T) {
	list := albumService.GetCategoryList(adminRole)
	assert.NotNil(t, list, "Expected list to be not nil")
}
