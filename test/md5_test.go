package test5

import (
	"crypto/md5"
	"encoding/hex"
	"testing"
)

func TestMD5(t *testing.T) {
	// 要加密的字符串
	input := "biaoge666"

	// 创建一个 MD5 hash 对象
	hash := md5.New()

	// 写入要加密的数据
	hash.Write([]byte(input))

	// 获取加密后的结果（哈希值）
	hashBytes := hash.Sum(nil)

	// 将哈希值转为16进制字符串
	hashString := hex.EncodeToString(hashBytes)

	// 输出加密后的结果
	t.Log("MD5 Hash:", hashString)
}
