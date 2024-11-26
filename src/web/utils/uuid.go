package utils

import (
	"github.com/google/uuid"
	"strings"
)

func GetRandomId(length int) string {
	return strings.Replace(uuid.New().String(), "-", "", -1)[0:length]
}
