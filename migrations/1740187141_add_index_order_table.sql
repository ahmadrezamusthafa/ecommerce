-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- [your SQL script here]
CREATE INDEX idx_orders_customer_id ON orders(customer_id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- [your SQL script here]
DROP INDEX IF EXISTS idx_orders_customer_id;
