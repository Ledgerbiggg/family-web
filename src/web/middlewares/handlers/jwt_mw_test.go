package handlers

import (
	"strings"
	"testing"
)

func TestJwtMiddleware_matchPath(t *testing.T) {
	path := matchPath("/a", "/a/*")
	t.Log(path)
}
func matchPath(path string, permissionPath string) bool {
	// 如果权限路径包含通配符
	if strings.Contains(permissionPath, "*") {
		// 获取通配符前的路径部分（即 `prefix`），假设是以 * 结尾
		prefix := strings.Split(permissionPath, "*")[0]

		// 检查路径是否以该前缀开头，且允许完全等于该前缀（即匹配 /a 和 /a/xxx 都符合）
		return strings.HasPrefix(path, prefix) || path+"/" == prefix || path == prefix
	}
	// 完全匹配
	return path == permissionPath
}
