DROP DATABASE IF EXISTS sqlxetude;

CREATE DATABASE sqlxetude CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
GRANT ALL ON sqlxetude.* TO mysql@'%';
DROP USER IF EXISTS sqlxetude@'%';
CREATE USER sqlxetude@'%' IDENTIFIED BY 'sqlxetude';
GRANT ALL PRIVILEGES ON sqlxetude.* TO sqlxetude@'%';

use sqlxetude;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INT AUTO_INCREMENT,
    name VARCHAR(200),
    birthday DATETIME,
    PRIMARY KEY (id)
);

DROP TABLE IF EXISTS comments;
CREATE TABLE comments (
    id INT AUTO_INCREMENT,
    users_id INT,
    body VARCHAR(200),
    post_at DATETIME,
    PRIMARY KEY (id)
);

INSERT INTO users (id, name, birthday)
VALUES
 (1, "hirosuzuki", "1993-10-10 00:00:00")
,(2, "armstrong", "1988-12-11 00:00:00")
,(3, "beatbeat", "2010-03-30 00:00:00")
;

INSERT INTO comments (users_id, body, post_at)
VALUES
 (1, "Great!",         "2022-06-22 12:20:05")
,(1, "Nice Nice Nice", "2022-06-22 14:03:33")
,(2, "Boo",            "2022-06-23 06:12:10")
;

