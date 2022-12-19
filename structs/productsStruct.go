package structs

import (
	"time"
)

type Product struct {
	ID                 int       `json:"proID"`
	Nama_produk        string    `json:"Nama_produk"`
	Jenis_produk       string    `json:"jenis_produk"`
	Deskripsi          string    `json:"deskripsi"`
	Tanggal_kadaluarsa string    `json:"tanggal_kadaluarsa"`
	Stok               int       `json:"stok"`
	Harga              int       `json:"harga"`
	Total_harga        int       `json:"total_harga"`
	Create_at          time.Time `json:"created_at"`
	Update_at          time.Time `json:"updated_at"`
	Vendor_id          int       `json:"vendor_id"`
}
