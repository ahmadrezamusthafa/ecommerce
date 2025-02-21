-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- [your SQL script here]
ALTER TABLE carts
    ADD CONSTRAINT unique_user_id UNIQUE (user_id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- [your SQL script here]
ALTER TABLE carts
    DROP CONSTRAINT unique_user_id;