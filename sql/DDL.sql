-- auto-generated definition
CREATE TABLE home_card
(
    id          int AUTO_INCREMENT COMMENT '主键'
        PRIMARY KEY,
    title       varchar(255)                       NOT NULL COMMENT '卡片的标题',
    description text                               NULL COMMENT '卡片的描述',
    image       varchar(255)                       NULL COMMENT '卡片的图片新',
    path        varchar(255)                       NOT NULL COMMENT '卡片指向的路径',
    created_at  datetime DEFAULT CURRENT_TIMESTAMP NULL COMMENT '卡片的创建时间',
    updated_at  datetime DEFAULT CURRENT_TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '卡片的更新时间'
)
    COMMENT '首页的卡片信息表';

-- auto-generated definition
CREATE TABLE invite_link
(
    id                int AUTO_INCREMENT COMMENT '主键'
        PRIMARY KEY,
    uuid              varchar(255)                       NOT NULL COMMENT '邀请链接的唯一标识',
    is_used           tinyint(1)                         NOT NULL COMMENT '链接是否已经使用',
    description       text                               NULL COMMENT '邀请链接描述',
    inviter_username  varchar(50)                        NOT NULL COMMENT '邀请人手机号',
    invited_real_name varchar(100)                       NOT NULL COMMENT '被邀请人真实姓名',
    invited_admin     tinyint(1)                         NOT NULL COMMENT '被邀请人角色是否是admin',
    expiration_date   datetime                           NOT NULL COMMENT '邀请链接过期时间',
    created_at        datetime DEFAULT CURRENT_TIMESTAMP NULL COMMENT '创建时间',
    used_at           datetime                           NULL COMMENT '使用时间'
)
    COMMENT '邀请链接表';

-- auto-generated definition
CREATE TABLE permission
(
    id          int AUTO_INCREMENT COMMENT '主键'
        PRIMARY KEY,
    path        varchar(255) NOT NULL COMMENT '权限路径',
    description text         NULL COMMENT '权限描述'
)
    COMMENT '权限表';

-- auto-generated definition
CREATE TABLE role
(
    id          int AUTO_INCREMENT COMMENT '主键'
        PRIMARY KEY,
    name        varchar(255) NOT NULL COMMENT '角色名称',
    description text         NULL COMMENT '角色描述'
)
    COMMENT '角色表';

-- auto-generated definition
CREATE TABLE role_home_card_access
(
    id           int AUTO_INCREMENT COMMENT '主键'
        PRIMARY KEY,
    role_id      int          NOT NULL COMMENT '角色 ID',
    home_card_id int          NOT NULL COMMENT '主页卡片 ID',
    description  varchar(255) NOT NULL COMMENT '描述'
)
    COMMENT '角色与主页卡片访问关系表';

-- auto-generated definition
CREATE TABLE role_permission
(
    id            int AUTO_INCREMENT COMMENT '主键'
        PRIMARY KEY,
    role_id       int          NOT NULL COMMENT '角色ID，关联角色表',
    permission_id int          NOT NULL COMMENT '权限ID，关联权限表',
    description   varchar(255) NOT NULL COMMENT '授权描述'
)
    COMMENT '角色与权限关联表';


-- auto-generated definition
CREATE TABLE user
(
    id              int AUTO_INCREMENT COMMENT '主键'
        PRIMARY KEY,
    username        varchar(255)         NOT NULL COMMENT '用户名(手机号)',
    password        varchar(255)         NOT NULL COMMENT '密码',
    nickname        varchar(255)         NULL COMMENT '昵称',
    is_disabled     tinyint(1) DEFAULT 0 NULL COMMENT '是否禁用',
    register_time   datetime             NOT NULL COMMENT '注册时间',
    last_login_time datetime             NULL COMMENT '最后登录时间',
    real_name       varchar(255)         NULL COMMENT '真实姓名',
    avatar          varchar(255)         NULL COMMENT '头像',
    email           varchar(255)         NULL COMMENT '邮箱',
    role_id         int                  NOT NULL COMMENT '角色(关联角色表的id)'
)
    COMMENT '用户表';

