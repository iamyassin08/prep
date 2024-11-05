-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
  
);

CREATE TABLE user_tags (
    tag_id INTEGER NOT NULL REFERENCES tags(id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    PRIMARY KEY (tag_id, user_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS user_tags;
DROP TABLE IF EXISTS users;

-- +goose StatementEnd
