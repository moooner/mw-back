-- create db 'mw'
CREATE DATABASE mw CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- use db
USE mw;

-- create account 'dev'
CREATE USER 'dev'@'localhost' IDENTIFIED BY 'dev123';
FLUSH PRIVILEGES;
GRANT ALL PRIVILEGES ON mw.* to dev@localhost;
GRANT SUPER ON *.* TO dev@localhost;
FLUSH PRIVILEGES;

-- drop and create tables
drop table if exists users;
create table if not exists users(
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    google_user_id VARCHAR(32) NULL,
    apple_user_id VARCHAR(32) NULL,
    email CHAR(50) NULL,
    customized_value FLOAT default(0)
) ENGINE = INNODB;

drop table if exists customized_value_logs;
create table if not exists customized_value_logs(
    log_id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    user_id int NOT NULL,
    question_id VARCHAR(32) NOT NULL,
    answer boolean
);

drop table if exists questions;
create table if not exists questions(
    question_id INT PRIMARY KEY AUTO_INCREMENT,
    question_text TEXT NOT NULL,
    delta FLOAT NOT NULL
);
