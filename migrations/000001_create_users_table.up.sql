-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    age INTEGER NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Создаем индекс для email
CREATE UNIQUE INDEX users_email_key ON users (email) WHERE deleted_at IS NULL; 