package utils

import (
	"fmt"
	"testing"
)

func TestReadPathAllDir(t *testing.T) {
	ReadPathAllDir("./testdir", func(dirName string) int {
		t.Log(dirName + "1111111111111")
		return 999
	}, func(dirId int, fileName string) int {
		t.Log(fmt.Sprintf("%d", dirId) + "1111111111111" + fileName)
		return 99
	})
}
