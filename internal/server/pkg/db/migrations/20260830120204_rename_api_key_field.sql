-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
RENAME COLUMN "api_key_hash" TO api_key_sha256;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
RENAME COLUMN "api_key_sha256" TO api_key_hash;
-- +goose StatementEnd
