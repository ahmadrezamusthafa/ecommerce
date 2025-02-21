-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- [your SQL script here]
ALTER TABLE accounts
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- [your SQL script here]