-- create db 'mw'
CREATE DATABASE mw CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- use db
USE mw;

-- create account 'moooner'
CREATE USER 'moooner'@'localhost' IDENTIFIED BY 'dev123';
FLUSH PRIVILEGES;
GRANT ALL PRIVILEGES ON mw.* to 'moooner'@'localhost';
GRANT SUPER ON *.* TO 'moooner'@'localhost';
FLUSH PRIVILEGES;

SELECT User, Host, authentication_string FROM mysql.user;

-- drop and create tables
drop table if exists `user`;
create table if not exists `user`(
    idx                 int unsigned        primary key auto_increment,
    google_user_id      varchar(32)         null,
    apple_user_id       varchar(32)         null,
    email               varchar(50)         null,
    customized_value    float               not null default 0
);

drop table if exists question;
create table if not exists question(
    idx                 int unsigned        primary key auto_increment,
    question_text       text                not null,
    delta               float               not null
);

drop table if exists customized_value_log;
create table if not exists customized_value_log(
    idx                 int unsigned        primary key auto_increment,
    user_idx            int unsigned        not null,
    question_idx        int unsigned        not null,
    answer              boolean             not null
);

