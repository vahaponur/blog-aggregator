-- +goose Up
CREATE TABLE feeds(
                      id UUID PRIMARY KEY,
                      created_at TIMESTAMP NOT NULL ,
                      updated_at TIMESTAMP NOT NULL ,
                      name VARCHAR NOT NULL,
                      url TEXT NOT NULL UNIQUE ,
                      user_id UUID REFERENCES users(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE feeds;