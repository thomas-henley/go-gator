-- +goose Up
CREATE TABLE users(
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULl,
    updated_at TIMESTAMP NOT NULL,
    name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE users;
