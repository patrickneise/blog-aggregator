-- +goose Up
CREATE TABLE
    feed_follows (
        id UUID PRIMARY KEY,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        user_id UUID NOT NULL,
        feed_id UUID NOT NULL,
        CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id),
        CONSTRAINT fk_feed FOREIGN KEY (feed_id) REFERENCES feeds (id)
    );

-- +goose Down
DROP TABLE feed_follows;