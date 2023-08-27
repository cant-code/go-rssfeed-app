-- +goose Up

CREATE TABLE POSTS (
    id UUID PRIMARY KEY,
    created_at timestamp NOT NULL,
    modified_at timestamp NOT NULL,
    title TEXT NOT NULL,
    description text,
    published_at timestamp NOT NULL,
    url text NOT NULL UNIQUE,
    feed_id UUID NOT NULL REFERENCES FEEDS(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE POSTS;