-- 创建用户表
CREATE TABLE `user`
(
    `id`              INT          NOT NULL AUTO_INCREMENT COMMENT '主键',
    `username`        VARCHAR(255) NOT NULL COMMENT '用户名(手机号)',
    `password`        VARCHAR(255) NOT NULL COMMENT '密码',
    `nickname`        VARCHAR(255) DEFAULT NULL COMMENT '昵称',
    `is_disabled`     BOOLEAN      DEFAULT FALSE COMMENT '是否禁用',
    `register_time`   DATETIME     NOT NULL COMMENT '注册时间',
    `last_login_time` DATETIME     DEFAULT NULL COMMENT '最后登录时间',
    `real_name`       VARCHAR(255) DEFAULT NULL COMMENT '真实姓名',
    `avatar`          VARCHAR(255) DEFAULT NULL COMMENT '头像',
    `email`           VARCHAR(255) DEFAULT NULL COMMENT '邮箱',
    `role_id`         INT          NOT NULL COMMENT '角色(关联角色表的id)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

-- 创建角色表
CREATE TABLE `role`
(
    `id`          INT          NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`        VARCHAR(255) NOT NULL COMMENT '角色名称',
    `description` TEXT DEFAULT NULL COMMENT '角色描述',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='角色表';

-- 创建权限表
CREATE TABLE `permission`
(
    `id`          INT          NOT NULL AUTO_INCREMENT COMMENT '主键',
    `path`        VARCHAR(255) NOT NULL COMMENT '权限路径',
    `description` TEXT DEFAULT NULL COMMENT '权限描述',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='权限表';

-- 创建角色与权限关联表
CREATE TABLE `role_permission`
(
    `id`            INT NOT NULL AUTO_INCREMENT COMMENT '主键',
    `role_id`       INT NOT NULL COMMENT '角色ID，关联角色表',
    `permission_id` INT NOT NULL COMMENT '权限ID，关联权限表',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='角色与权限关联表';

-- 创建链接表(用于存放管理员生成的邀请链接数据,可以用于邀请注册其他管理员的家庭成员)
CREATE TABLE `invite_link`
(
    `id`                INT          NOT NULL AUTO_INCREMENT COMMENT '主键',
    `uuid`              VARCHAR(255) NOT NULL COMMENT '邀请链接的唯一标识',
    `is_used`           BOOLEAN      NOT NULL COMMENT '链接是否已经使用',
    `description`       TEXT     DEFAULT NULL COMMENT '邀请链接描述',
    `inviter_username`  VARCHAR(50)  NOT NULL COMMENT '邀请人手机号',
    `invited_real_name` VARCHAR(100) NOT NULL COMMENT '被邀请人真实姓名',
    `invited_admin`     BOOLEAN      NOT NULL COMMENT '被邀请人角色是否是admin',
    `expiration_date`   DATETIME     NOT NULL COMMENT '邀请链接过期时间',
    `created_at`        DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `used_at`           DATETIME DEFAULT NULL COMMENT '使用时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='邀请链接表';

-- 主页的轮播图表信息
CREATE TABLE `home_cards`
(
    `id`          INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    `title`       VARCHAR(255) NOT NULL COMMENT '卡片的标题',
    `description` TEXT COMMENT '卡片的描述',
    `image`       VARCHAR(255) COMMENT '卡片的图片新',
    `path`        VARCHAR(255) NOT NULL COMMENT '卡片指向的路径',
    `created_at`  DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '卡片的创建时间',
    `updated_at`  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '卡片的更新时间',
    `user_id`     INT COMMENT '与此卡片关联的用户 ID'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='首页的卡片信息表';





