create database if not exists login_db default CHARACTER SET UTF8;

# GRANT ALL PRIVILEGES ON login_db.* TO @localhost IDENTIFIED BY 'login';
USE login_db;

drop table if exists users;
create table if not exists users(
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    google_user_id VARCHAR(32) NOT NULL,
    apple_user_id VARCHAR(32) NOT NULL,
    email CHAR(50) NOT NULL,
    customized_value FLOAT
) ENGINE = INNODB;

drop table if exists customized_value_logs;
create table if not exists customized_value_logs(
    log_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id int,
    question_id VARCHAR(32) NOT NULL,
    answer int
);

drop table if exists questions;
create table if not exists queuser_idstions(
    question_id INT PRIMARY KEY AUTO_INCREMENT,
    question_text VARCHAR(100),
    delta FLOAT
);
