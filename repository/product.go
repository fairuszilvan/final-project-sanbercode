package repository

import (
	"database/sql"
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/structs"
	"time"
)

func IndexProduct(db *sql.DB) (results []structs.Product, err error) {
	sql := `SELECT * FROM product`
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var product = structs.Product{}
		err = rows.Scan(&product.ID, &product.Nama_produk, &product.Jenis_produk, &product.Deskripsi, &product.Tanggal_kadaluarsa, &product.Stok, &product.Harga, &product.Total_harga, &product.Create_at, &product.Update_at, &product.Vendor_id)
		if err != nil {
			panic(err)
		}
		results = append(results, product)
	}
	return
}
func InsertProduct(db *sql.DB, product structs.Product) (err error) {
	sql := "INSERT INTO product (nama_produk,jenis_produk,deskripsi,tanggal_kadaluarsa,stok,harga,total_harga,created_at,updated_at,vendor_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"
	s := product.Stok
	h := product.Harga
	t := s * h
	errs := db.QueryRow(sql, product.Nama_produk, product.Jenis_produk, product.Deskripsi, product.Tanggal_kadaluarsa, product.Stok, product.Harga, t, time.Now(), time.Now(), product.Vendor_id)
	return errs.Err()
}
func UpdateProduct(db *sql.DB, product structs.Product) (err error) {
	sql := "UPDATE product SET nama_produk=$2,jenis_produk=$3,deskripsi=$4,tanggal_kadaluarsa=$5,stok=$6,harga=$7,total_harga=$8,created_at=$9,updated_at,=$10,vendor_id=$11 WHERE id=$1"
	s := product.Stok
	h := product.Harga
	t := s * h
	errs := db.QueryRow(sql, product.ID, product.Nama_produk, product.Jenis_produk, product.Deskripsi, product.Tanggal_kadaluarsa, product.Stok, product.Harga, t, product.Create_at, product.Update_at, product.Vendor_id)
	return errs.Err()
}
func DeleteProduct(db *sql.DB, product structs.Product) (err error) {
	sql := "DELETE FROM product WHERE id=$1"
	errs := db.QueryRow(sql, product.ID)
	return errs.Err()
}
