-- +goose Up
ALTER TABLE users
ADD COLUMN apikey VARCHAR(64) DEFAULT encode(sha256(random()::text::bytea), 'hex') NOT NULL;
-- +goose Down
ALTER TABLE USERS
DROP COLUMN apikey;