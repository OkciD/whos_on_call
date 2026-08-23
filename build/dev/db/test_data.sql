BEGIN TRANSACTION;

-- Вставляем пользователя
INSERT INTO users (name, api_key_hash)
VALUES (
    '__NAME__',
    '__KEY__'
);

-- Получаем ID последнего вставленного пользователя
WITH user_data AS (
    SELECT last_insert_rowid() AS user_id
)
-- Вставляем устройство, используя этот ID
INSERT INTO devices (name, type, user_id)
VALUES (
    'Workstation',
    0, -- laptop
    (SELECT user_id FROM user_data)
);

-- Получаем ID устройства
WITH device_data AS (
    SELECT last_insert_rowid() AS device_id
)
-- Вставляем функции устройства
INSERT INTO device_features (type, status, device_id)
VALUES
(
    0, -- mic
    0, -- mic, inactive
    (SELECT device_id FROM device_data)
),
(
    1, -- camera
    0, -- inactive
    (SELECT device_id FROM device_data)
);

COMMIT;
