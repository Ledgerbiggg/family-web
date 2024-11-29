-- 创建用户表
CREATE TABLE `user`
(
    `id`              INT          NOT NULL AUTO_INCREMENT COMMENT '主键',
    `username`        VARCHAR(255) NOT NULL COMMENT '用户名(手机号)',
    `password`        VARCHAR(255) NOT NULL COMMENT '密码',
    `nickname`       VARCHAR(255) DEFAULT NULL COMMENT '昵称',
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



