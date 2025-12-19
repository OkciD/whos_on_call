-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS devices (
    id INTEGER PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    type INT8 NOT NULL,
    user_id INTEGER REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_devices_user_id ON devices (user_id);

CREATE TABLE IF NOT EXISTS device_features (
    id INTEGER PRIMARY KEY,
    type INT8 NOT NULL,
    status INT8 NOT NULL,
    last_modified DATETIME,
    device_id INTEGER REFERENCES devices(id)
);

CREATE INDEX IF NOT EXISTS idx_device_feature_device_id ON device_features (device_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_device_feature_device_id;
DROP TABLE IF EXISTS device_features;

DROP INDEX IF EXISTS idx_devices_user_id;
DROP TABLE IF EXISTS devices;
-- +goose StatementEnd
