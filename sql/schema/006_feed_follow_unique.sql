-- +goose Up
ALTER TABLE feed_follows
    ADD CONSTRAINT unique_user_feed_pair UNIQUE (user_id, feed_id);
-- +goose Down
ALTER TABLE feed_follows
DROP CONSTRAINT IF EXISTS unique_user_feed_pair;