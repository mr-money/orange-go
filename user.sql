create table user
(
    id         bigint auto_increment comment 'id'
        primary key,
    uuid       varchar(50) default ''                not null comment '全局唯一标识',
    name       varchar(20) default ''                not null comment 'name',
    password   varchar(64) default ''                not null comment '密码',
    created_at timestamp   default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    updated_at timestamp   default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at timestamp                             null,
    constraint user_name_uindex
        unique (name),
    constraint user_uuid_uindex
        unique (uuid)
)
    comment '用户表' engine = InnoDB
                     charset = utf8mb4;

create index user_created_at_index
    on user (created_at);

create index user_updated_at_index
    on user (updated_at);

INSERT INTO go_study.user (id, uuid, name, password, created_at, updated_at, deleted_at) VALUES (1, 'd8aae903-d52f-4d1c-b84a-ac5ea18d0f8d', 'aaaa', '', '2022-04-27 23:14:00', '2022-04-27 23:14:00', null);
INSERT INTO go_study.user (id, uuid, name, password, created_at, updated_at, deleted_at) VALUES (2, '76c2e0c4-3bed-4d6e-be24-3fc1c289e6d1', '哈哈哈', '', '2022-04-27 23:14:00', '2022-04-27 23:14:00', null);
INSERT INTO go_study.user (id, uuid, name, password, created_at, updated_at, deleted_at) VALUES (3, 'ed27d954-2f1c-4243-a424-1d9f483f251a', '哦哦哦', '', '2022-04-27 23:14:00', '2022-04-27 23:14:00', null);
INSERT INTO go_study.user (id, uuid, name, password, created_at, updated_at, deleted_at) VALUES (4, 'c08f0481-8ee6-4420-8a4c-dc8971fd07ff', '卡扣', '', '2022-04-27 23:14:00', '2022-04-27 23:14:00', null);
INSERT INTO go_study.user (id, uuid, name, password, created_at, updated_at, deleted_at) VALUES (5, 'e5fa007b-c838-44b6-820d-2f8ab5629e4e', '卡扣2', '', '2022-04-27 23:14:00', '2022-04-27 23:14:00', '2022-04-22 21:45:52');
INSERT INTO go_study.user (id, uuid, name, password, created_at, updated_at, deleted_at) VALUES (6, 'a8044564-86e5-4f33-8bea-bd4014ee7ae1', '略略', '', '2022-04-27 23:14:00', '2022-04-27 23:14:00', null);
INSERT INTO go_study.user (id, uuid, name, password, created_at, updated_at, deleted_at) VALUES (7, '12a33124-bba7-4f1b-b946-fccf642721cb', '去玩儿', '$2a$04$ecYg4N/M8IXiXlGsJ6/YQO81MK3XjO/A6dG5Mw78pBTpouPTQjZsi', '2022-04-27 23:14:00', '2022-04-27 23:14:00', null);
INSERT INTO go_study.user (id, uuid, name, password, created_at, updated_at, deleted_at) VALUES (8, '1ecf1d25-5f0e-43bb-8f7b-d8af95fe9bf5', '去玩儿22', '$2a$04$V07.KkqXPfqWeCkvurGDte1MaiXo9LWmaijv.cZEd6flY/3akuU9u', '2022-04-27 23:14:00', '2022-04-27 23:14:00', null);
INSERT INTO go_study.user (id, uuid, name, password, created_at, updated_at, deleted_at) VALUES (9, 'dfe19073-956b-4a18-acae-d340a547dcd2', '123456', '$2a$04$J1TxVXKixwUOoKIr3FoY3uYG9csggHd/PSITxxpz3GnRMVnCCuwFm', '2022-04-27 23:14:00', '2022-04-27 23:14:00', null);
INSERT INTO go_study.user (id, uuid, name, password, created_at, updated_at, deleted_at) VALUES (10, 'cb12a069-ec9e-45d3-a6f7-c605fe6561ad', '111111', '$2a$04$.svRq5tPttaWlXKGQEFiBuLJSA4bcsKMDw/LVPfbwSid/nZ5LZyG.', '2022-04-27 23:14:00', '2022-04-27 23:14:00', null);
