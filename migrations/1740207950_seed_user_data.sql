-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- [your SQL script here]
INSERT INTO users (username, "name", email, "password", created_at, updated_at)
VALUES ('ahmad', 'Ahmad', 'ahmad@abc.com', '$2a$10$MGfzJzRYabIxfIMDEOfKxOoS7c8FeDK5D26PvXFRlKRub9dsW0lZO',
        '2025-02-22 14:05:05.929', '2025-02-22 14:05:05.929');
INSERT INTO users (username, "name", email, "password", created_at, updated_at)
VALUES ('rozi', 'Rozi', 'rozi@abc.com', '$2a$10$MGfzJzRYabIxfIMDEOfKxOoS7c8FeDK5D26PvXFRlKRub9dsW0lZO',
        '2025-02-22 14:05:05.929', '2025-02-22 14:05:05.929');
INSERT INTO users (username, "name", email, "password", created_at, updated_at)
VALUES ('miumiu', 'Miumiu', 'miumiu@abc.com', '$2a$10$MGfzJzRYabIxfIMDEOfKxOoS7c8FeDK5D26PvXFRlKRub9dsW0lZO',
        '2025-02-22 14:05:05.929', '2025-02-22 14:05:05.929');
INSERT INTO users(username, "name", email, "password", created_at, updated_at)
VALUES ('cici', 'Cici', 'cici@abc.com', '$2a$10$MGfzJzRYabIxfIMDEOfKxOoS7c8FeDK5D26PvXFRlKRub9dsW0lZO',
        '2025-02-22 14:05:05.929', '2025-02-22 14:05:05.929');
INSERT INTO users(username, "name", email, "password", created_at, updated_at)
VALUES ('susi', 'Susi', 'susi@abc.com', '$2a$10$MGfzJzRYabIxfIMDEOfKxOoS7c8FeDK5D26PvXFRlKRub9dsW0lZO',
        '2025-02-22 14:05:05.929', '2025-02-22 14:05:05.929');
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- [your SQL script here]
