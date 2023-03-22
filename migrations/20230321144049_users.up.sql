CREATE TABLE IF NOT EXISTS users
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(50) UNIQUE NOT NULL,
    password     VARCHAR(100)       NOT NULL,
    created_date TIMESTAMP          NOT NULL,
    updated_date TIMESTAMP          NOT NULL,
    deleted_date TIMESTAMP          NULL
);