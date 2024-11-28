package test5

import "testing"

func TestGorm(t *testing.T) {
	var results []map[string]any
	gorm.GetDb().Raw(`
		SELECT r.name n, r.description d, p.path p, p.description des
			FROM role r
					 LEFT JOIN role_permission pr ON r.id = pr.role_id
					 LEFT JOIN permission p ON pr.permission_id = p.id
			WHERE r.id = 1`).Scan(&results)
	t.Log(results)
}
