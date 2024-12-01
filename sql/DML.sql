#  根据用户的名称(手机号码)查询用户对应的角色的全部的路由权限
SELECT p.id, p.path, p.description
FROM user u
         LEFT JOIN role r ON u.role_id = r.id
         LEFT JOIN role_permission rp ON r.id = rp.role_id
         LEFT JOIN permission p ON rp.permission_id = p.id
WHERE u.username = '18888888888';

# 根据用户名校验用户是否是管理员权限
SELECT EXISTS (
    SELECT 1
    FROM user u
             LEFT JOIN role r ON u.role_id = r.id
    WHERE u.username = '18888888888'
      AND (r.name = 'admin' OR r.name = 'root')
)




