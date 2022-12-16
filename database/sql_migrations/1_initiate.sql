-- +migrate Up
-- +migrate StatementBegin

Create Table vendor(
    id SERIAL PRIMARY KEY,
    nama_vendor VARCHAR(64),
    alamat VARCHAR(64),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
)

-- +migrate StatementEnd