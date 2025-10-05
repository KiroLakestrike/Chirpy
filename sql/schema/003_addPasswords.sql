-- +goose Up
ALTER TABLE users
    ADD COLUMN hashed_passwords TEXT NOT NULL default 'unset';

-- +goose Down
ALTER TABLE users
    DROP COLUMN hashed_passwords;
