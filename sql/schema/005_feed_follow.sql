-- +goose Up
CREATE TABLE feed_follows(
    id UUID PRIMARY KEY ,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL ,
    feed_id UUID REFERENCES feeds(id) ON DELETE CASCADE NOT NULL ,
    created_at TIMESTAMP NOT NULL ,
    updated_at TIMESTAMP NOT NULL

);
-- +goose Down
DROP TABLE feed_follow;