-- +migrate Up
-- +migrate StatementBegin

Create Table product(
    id SERIAL PRIMARY KEY,
    nama_produk VARCHAR(64),
    jenis_produk VARCHAR(64),
    deskripsi VARCHAR(64),
    tanggal_kadaluarsa VARCHAR(64),
    stok INT,
    harga BIGINT,
    total_harga BIGINT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    vendor_id INT, 
    FOREIGN KEY (vendor_id) REFERENCES vendor(id)
)

-- +migrate StatementEnd