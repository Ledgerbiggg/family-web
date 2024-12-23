package test5

import (
	"testing"
)

func TestAlbumController(t *testing.T) {
	name, err := albumService.GetImageBytesByName("1", "1")
	if err != nil {
		t.Error(err)
	}
	t.Log(name)
}
