-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- [your SQL script here]
CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(50) UNIQUE  NOT NULL,
    name       VARCHAR(100)        NOT NULL,
    email      VARCHAR(100) UNIQUE NOT NULL,
    password   VARCHAR(255)        NOT NULL,
    created_at TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP                    DEFAULT CURRENT_TIMESTAMP
);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- [your SQL script here]
DROP TABLE IF EXISTS users;