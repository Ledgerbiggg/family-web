package utils

import (
	"os"
	"testing"
)

func TestGenerateQRCode(t *testing.T) {
	bytes, err := GenerateQRCode("https://www.baidu.com", 100)
	if err != nil {
		return
	}
	file, err := os.OpenFile("test.png", os.O_CREATE, 0644)
	if err != nil {
		return
	}
	// 将二维码字节写入文件
	_, err = file.Write(bytes)
	if err != nil {
		t.Fatalf("Failed to write QR code to file: %v", err)
	}
}
