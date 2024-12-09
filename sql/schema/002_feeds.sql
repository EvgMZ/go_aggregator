-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name VARCHAR(255) UNIQUE NOT NULL,
    url VARCHAR(255) UNIQUE,
    user_id UUID NOT NULL,              -- Идентификатор пользователя, добавившего канал
    CONSTRAINT fk_user FOREIGN KEY (user_id) 
        REFERENCES users (id) 
        ON DELETE CASCADE              -- Удаление кана
);

-- +goose Down
DROP TABLE IF EXISTS feeds;