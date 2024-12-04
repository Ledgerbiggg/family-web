SELECT *
FROM `user`
WHERE username = '18288888888';



INSERT INTO `user` (`username`, `password`, `nickname`, `is_disabled`, `register_time`, `last_login_time`, `real_name`,
                    `avatar`, `email`, `role_id`)
VALUES ('', 'E10ADC3949BA59ABBE56E057F20F883E', NULL, FALSE, '0000-00-00 00:00:00', NULL, NULL, NULL, NULL, 2);


# 查询角色和主页卡片的对应关系
SELECT hc.*
FROM role r
         LEFT JOIN role_home_card_access rhca ON r.id = rhca.role_id
         LEFT JOIN home_card hc ON rhca.home_card_id = hc.id
WHERE r.id = 1;



