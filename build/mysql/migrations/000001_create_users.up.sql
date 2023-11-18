CREATE TABLE IF NOT EXISTS users
(
    id         INT         NOT NULL AUTO_INCREMENT,
    account    VARCHAR(64) NOT NULL,
    password   VARCHAR(64) NOT NULL,
    session_id VARCHAR(64),
    login_time TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);