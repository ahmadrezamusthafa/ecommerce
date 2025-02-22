-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- [your SQL script here]
ALTER TABLE orders ADD COLUMN quantity INT NOT NULL DEFAULT 0;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- [your SQL script here]
