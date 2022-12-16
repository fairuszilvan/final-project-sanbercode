-- +migrate Up
-- +migrate StatementBegin

Create Table users(
    id SERIAL PRIMARY KEY,
    profile_url VARCHAR(64),
    nama_users VARCHAR(64),
    email VARCHAR(64),
    password VARCHAR(64),
    role VARCHAR(64),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
)

-- +migrate StatementEnd