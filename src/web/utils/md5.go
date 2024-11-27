package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encrypt(input string) string {
	// 创建一个 MD5 hash 对象
	hash := md5.New()
	// 写入要加密的数据
	hash.Write([]byte(input))
	// 获取加密后的结果（哈希值）
	hashBytes := hash.Sum(nil)
	// 将哈希值转为16进制字符串
	return strings.ToUpper(hex.EncodeToString(hashBytes))
}
