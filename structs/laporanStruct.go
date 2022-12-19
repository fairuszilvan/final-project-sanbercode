package structs

import (
	"time"
)

type Laporan struct {
	ID               int       `json:"laporanID"`
	Jenis_laporan    string    `json:"jenis_laporan"`
	Bukti_pengesahan string    `json:"bukti_pengesahan"`
	Create_at        time.Time `json:"created_at"`
	Update_at        time.Time `json:"updated_at"`
	Vendor_id        int       `json:"vendor_id"`
	Users_id         int       `json:"users_id"`
	Produk_id        int       `json:"produk_id"`
}
type ViewLaporan struct {
	ID                 int       `json:"laporanID"`
	Nama_produk        string    `json:"Nama_produk"`
	Jenis_produk       string    `json:"jenis_produk"`
	Tanggal_kadaluarsa string    `json:"tanggal_kadaluarsa"`
	Stok               int       `json:"stok"`
	Harga              int       `json:"harga"`
	Total_harga        int       `json:"total_harga"`
	Create_at          time.Time `json:"created_at"`
	Nama_vendor        string    `json:"nama_vendor"`
	Nama_users         string    `json:"nama_user"`
	Jenis_laporan      string    `json:"jenis_laporan"`
	Bukti_pengesahan   string    `json:"bukti_pengesahan"`
}
