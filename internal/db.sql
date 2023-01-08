create database game charset utf8mb4;

create table users(
    id int not null auto_increment primary key,
    phone varchar(11) not null default '' comment '手机号',
    created_at int not null default 0 comment '创建时间'
);