package utils

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("18288888888", nil, nil, "test", 24, "ledger")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
