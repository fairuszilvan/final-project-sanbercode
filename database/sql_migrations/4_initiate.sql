-- +migrate Up
-- +migrate StatementBegin

Create Table laporan(
    id SERIAL PRIMARY KEY,
    jenis_laporan VARCHAR(64),
    bukti_pengesahan VARCHAR(64),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    produk_id INT,
    FOREIGN KEY (produk_id) REFERENCES product(id),
    vendor_id INT, 
    FOREIGN KEY (vendor_id) REFERENCES vendor(id),
    users_id INT,
    FOREIGN KEY (users_id) REFERENCES users(id)
)

-- +migrate StatementEnd