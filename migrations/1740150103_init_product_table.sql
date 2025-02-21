-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- [your SQL script here]
CREATE TABLE products
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100)   NOT NULL,
    description TEXT,
    price       DECIMAL(10, 2) NOT NULL,
    created_at  TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP               DEFAULT CURRENT_TIMESTAMP
);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- [your SQL script here]
DROP TABLE IF EXISTS products;