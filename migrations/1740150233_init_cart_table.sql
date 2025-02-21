-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- [your SQL script here]
CREATE TABLE carts
(
    id      SERIAL PRIMARY KEY,
    user_id INT NOT NULL
);

CREATE TABLE cart_items
(
    id         SERIAL PRIMARY KEY,
    cart_id    INT NOT NULL,
    product_id INT NOT NULL
);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- [your SQL script here]
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS cart_items;