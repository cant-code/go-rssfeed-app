-- +goose Up

ALTER TABLE FEEDS ADD COLUMN last_fetched_at timestamp;

-- +goose Down
ALTER TABLE FEEDS DROP COLUMN last_fetched_at;