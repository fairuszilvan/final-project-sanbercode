package repository

import (
	"database/sql"
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/structs"
	"time"
)

func IndexLaporan(db *sql.DB) (results []structs.Laporan, err error) {
	sql := `SELECT * FROM laporan`
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var laporan = structs.Laporan{}
		err = rows.Scan(&laporan.ID, &laporan.Jenis_laporan, &laporan.Bukti_pengesahan, &laporan.Create_at, &laporan.Update_at, &laporan.Vendor_id, &laporan.Users_id, &laporan.Produk_id)
		if err != nil {
			panic(err)
		}
		results = append(results, laporan)
	}
	return
}
func InsertLaporan(db *sql.DB, laporan structs.Laporan) (err error) {
	sql := "INSERT INTO laporan (jenis_laporan,bukti_pengesahan,created_at,updated_at,vendor_id,users_id,produk_id) VALUES ($1,$2,$3,$4,$5,$6,$7)"
	errs := db.QueryRow(sql, laporan.Jenis_laporan, laporan.Bukti_pengesahan, time.Now(), time.Now(), laporan.Vendor_id, laporan.Users_id, laporan.Produk_id)
	return errs.Err()
}
func UpdateLaporan(db *sql.DB, laporan structs.Laporan) (err error) {
	sql := "UPDATE laporan SET jenis_laporan=$2,bukti_pengesahan=$3,created_at=$4,updated_at=$5,vendor_id=$6,users_id=$7,produk_id=$8 WHERE id=$1"
	errs := db.QueryRow(sql, laporan.Jenis_laporan, laporan.Bukti_pengesahan, laporan.Create_at, time.Now(), laporan.Vendor_id, laporan.Users_id, laporan.Produk_id)
	return errs.Err()
}
func DeleteLaporan(db *sql.DB, laporan structs.Laporan) (err error) {
	sql := "DELETE FROM laporan WHERE id=$1"
	errs := db.QueryRow(sql, laporan.ID)
	return errs.Err()
}
func CreateView(db *sql.DB) (err error) {
	sql := `CREATE VIEW view_laporan as SELECT a.id,d.nama_produk,d.jenis_produk,d.tanggal_kadaluarsa,d.stok,d.harga,d.total_harga,d.created_at,b.nama_vendor,c.nama_users,a.jenis_laporan,a.bukti_pengesahan
	FROM laporan AS a
	LEFT JOIN vendor AS b ON a.vendor_id = b.id
	LEFT JOIN users AS c ON a.users_id = c.id
	LEFT JOIN product AS d ON a.produk_id = d.id`
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return
}

func ViewLaporan(db *sql.DB) (results []structs.ViewLaporan, err error) {
	sql := `SELECT * FROM view_laporan`
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var view = structs.ViewLaporan{}
		err = rows.Scan(&view.ID, &view.Nama_produk, &view.Jenis_produk, &view.Tanggal_kadaluarsa, &view.Stok, &view.Harga, &view.Total_harga, &view.Create_at, &view.Nama_vendor, &view.Nama_users, &view.Jenis_laporan, &view.Bukti_pengesahan)
		if err != nil {
			panic(err)
		}
		results = append(results, view)
	}
	return
}
