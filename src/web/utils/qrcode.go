package utils

import (
	"github.com/skip2/go-qrcode"
	"log"
)

// GenerateQRCode 将链接生成二维码并返回字节数组
func GenerateQRCode(url string, size int) ([]byte, error) {
	// 调用 go-qrcode 库生成二维码
	qrCode, err := qrcode.Encode(url, qrcode.Medium, size)
	if err != nil {
		log.Printf("Failed to generate QR code: %v", err)
		return nil, err
	}
	return qrCode, nil
}
