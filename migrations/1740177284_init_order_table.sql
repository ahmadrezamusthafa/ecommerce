-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- [your SQL script here]
CREATE TABLE orders
(
    id          SERIAL PRIMARY KEY,
    customer_id INT            NOT NULL,
    product_id  INT            NOT NULL,
    order_date  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    amount      DECIMAL(10, 2) NOT NULL
);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- [your SQL script here]
DROP TABLE IF EXISTS orders;