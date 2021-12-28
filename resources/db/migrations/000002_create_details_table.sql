-- +migrate Up
CREATE TABLE IF NOT EXISTS detail_records
(
    id  serial PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    price float  NOT NULL,
    created_at  DATE NOT NULL
);
-- +migrate Down
DROP TABLE detail_records;