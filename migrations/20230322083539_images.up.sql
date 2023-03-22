CREATE TABLE IF NOT EXISTS images
(
    id           SERIAL PRIMARY KEY,
    user_id      INTEGER      NOT NULL REFERENCES users (id),
    image_path   VARCHAR(250) NOT NULL,
    image_url    VARCHAR(250) NOT NULL,
    created_date TIMESTAMP    NOT NULL
);