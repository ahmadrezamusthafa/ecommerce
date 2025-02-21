-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- [your SQL script here]
CREATE TABLE accounts
(
    id      SERIAL PRIMARY KEY,
    user_id INT            NOT NULL,
    balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00
);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- [your SQL script here]
DROP TABLE IF EXISTS accounts;