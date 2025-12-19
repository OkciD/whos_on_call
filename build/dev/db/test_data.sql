INSERT INTO users (id, name, api_key_hash) VALUES (
    0,
    "John Doe",
    "abb45ef89186194a3e3ee700894caeb3b86ced38db1fa3ec7fd0f5e2ff6d9ec1" -- sha256("lolkek")
);

INSERT INTO devices (id, name, type, user_id) VALUES (
    0,
    "Workstation",
    0, -- laptop
    0
);

INSERT INTO device_features (id, type, status, device_id) VALUES
(
    0,
    0, -- mic
    0, -- inactive,
    0
),
(
    1,
    1, -- camera
    0, -- inactive,
    0
);
