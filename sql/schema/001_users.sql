-- +goose Up

CREATE TABLE USERS (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(20) NOT NULL
);

-- +goose Down
DROP TABLE USERS;