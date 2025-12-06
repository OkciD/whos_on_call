-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    api_key_hash CHAR(64) NOT NULL UNIQUE
);

CREATE INDEX IF NOT EXISTS idx_users_api_key_hash ON users (api_key_hash);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_users_api_key_hash;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
