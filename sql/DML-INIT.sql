# 插入一些关键数据(导航页面需要使用)
INSERT INTO `home_card` (`title`, `description`, `image`, `path`)
VALUES ('相册', '拍下最美瞬间', '01.png', '/album'),
       ('视频', '记录美好时刻', '02.png', '/video'),
       ('日程', '别忘了那些纪念日', '03.png', '/calendar'),
       ('微信Bot', '微信群的小助手', '04.png', '/bot');

-- 插入一些必要的角色关联
INSERT INTO role_home_card_access (id, role_id, home_card_id)
VALUES (1, 1, 1),
       (2, 1, 2),
       (3, 1, 3),
       (4, 1, 4),
       (5, 2, 1),
       (6, 2, 2),
       (7, 2, 3),
       (8, 2, 4);
