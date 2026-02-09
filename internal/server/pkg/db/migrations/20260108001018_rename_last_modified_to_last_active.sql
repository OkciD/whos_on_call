-- +goose Up
-- +goose StatementBegin
ALTER TABLE device_features RENAME COLUMN last_modified TO last_active;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE device_features RENAME COLUMN last_active TO last_modified;
-- +goose StatementEnd
