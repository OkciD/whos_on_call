-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    name VARCHAR(256) NOT NULL UNIQUE,
    api_key_hash CHAR(64) NOT NULL UNIQUE -- индекс по этому полю уже будет создан под капотом
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
