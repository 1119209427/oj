create table problem
(
    id          int auto_increment,
    identity    varchar(36)  null,
    title       varchar(255) null comment '问题的题目',
    content     text         null comment '问题的描述',
    max_runtime int          null comment '最大运行时间',
    max_mem     int          null comment '最大运行内存',
    created_at  datetime     null,
    constraint problem_pk
        primary key (id)
);

create index problem_identity_category_id_index
    on problem (identity);

create table user
(
    id         int auto_increment,
    identity   varchar(36)  null,
    name       int          null,
    password   int          null,
    salt       varchar(32)  null,
    phone      varchar(20)  null,
    mall       varchar(100) null,
    created_at datetime     null,
    constraint user_pk
        primary key (id)
);

create index user_identity_index
    on user (identity);

create table category
(
    id         int auto_increment,
    name       varchar(100) null comment '分类名称',
    parent_id  int          null comment '父级id',
    created_at datetime     null,
    constraint category_pk
        primary key (id)
);

create table submit
(
    id               int auto_increment,
    identity         varchar(37)  null,
    problem_identity int          null comment '问题的唯一标识',
    user_identity    varchar(36)  null comment '用户的唯一标识',
    path             varchar(255) null comment '代码路径',
    status           tinyint      null comment '0-待判断，1-答案正确，2-答案错误，3-答案超时，4-答案超内存',
    created_at       datetime     null,
    constraint submit_pk
        primary key (id)
);

create index submit_identity_problem_identity_user_identity_index
    on submit (identity, problem_identity, user_identity);

create table category_problem
(
    id          int auto_increment,
    category_id int null,
    problem_id  int null,
    constraint category_problem_pk
        primary key (id)
);

create index category_problem_category_id_problem_id_index
    on category_problem (category_id, problem_id);

alter table user
    modify identity varchar(36) collate latin1_swedish_ci null;

alter table user
    modify name varchar(100) null;

alter table user
    modify salt varchar(32) collate latin1_swedish_ci null;

alter table user
    modify phone varchar(20) collate latin1_swedish_ci null;

alter table user
    modify mall varchar(100) collate latin1_swedish_ci null;

create table test_case
(
    id         int auto_increment,
    problem_id int         not null comment '所属问题',
    input      varchar(36) null comment '输入',
    output     varchar(36) null comment '输出',
    constraint test_case_pk
        primary key (id)
);

create index test_case_id_problem_id_index
    on test_case (id, problem_id);