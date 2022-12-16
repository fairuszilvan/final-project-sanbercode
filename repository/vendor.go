package repository

import (
	"database/sql"
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/structs"
	"time"
)

func IndexVendor(db *sql.DB) (results []structs.Vendor, err error) {
	sql := "SELECT * FROM vendor"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var vendor = structs.Vendor{}
		err = rows.Scan(&vendor.ID, &vendor.Nama_vendor, &vendor.Alamat, &vendor.Create_at, &vendor.Update_at)
		if err != nil {
			panic(err)
		}
		results = append(results, vendor)
	}
	return
}
func InsertVendor(db *sql.DB, vendor structs.Vendor) (err error) {
	sql := "INSERT INTO vendor (nama_vendor,alamat,created_at,updated_at) VALUES ($1,$2,$3,$4)"
	errs := db.QueryRow(sql, vendor.Nama_vendor, vendor.Alamat, time.Now(), time.Now())
	return errs.Err()
}
func UpdateVendor(db *sql.DB, vendor structs.Vendor) (err error) {
	sql := "UPDATE vendor SET nama_vendor=$2,alamat=$3,created_at=$6,updated_at=$5 WHERE id=$1"
	errs := db.QueryRow(sql, vendor.ID, vendor.Nama_vendor, vendor.Alamat, vendor.Create_at, vendor.Update_at)
	return errs.Err()
}
func DeleteVendor(db *sql.DB, vendor structs.Vendor) (err error) {
	sql := "DELETE FROM vendor WHERE id=$1"
	errs := db.QueryRow(sql, vendor.ID)
	return errs.Err()
}
