CREATE TABLE IF NOT EXISTS containers (
    id SERIAL PRIMARY KEY,
    ip_address VARCHAR(255) NOT NULL,
    ping_time TIMESTAMP NOT NULL,
    last_checked TIMESTAMP NOT NULL
);