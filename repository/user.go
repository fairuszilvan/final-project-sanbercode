package repository

import (
	"database/sql"
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/structs"
	"time"
)

func IndexUser(db *sql.DB) (results []structs.User, err error) {
	sql := "SELECT * FROM users"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user = structs.User{}
		err = rows.Scan(&user.ID, &user.Profile_url, &user.Nama_users, &user.Email, &user.Password, &user.Role, &user.Create_at, &user.Update_at)
		if err != nil {
			panic(err)
		}
		results = append(results, user)
	}
	return
}
func InsertUser(db *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO users (profile_url,nama_users,email,password,role,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)"
	errs := db.QueryRow(sql, user.Profile_url, user.Nama_users, user.Email, user.Password, user.Role, time.Now(), time.Now())
	return errs.Err()
}
func UpdateUser(db *sql.DB, user structs.User) (err error) {
	sql := "UPDATE users SET profile_url=$2,nama_users=$3,email=$4,password=$5,role=$6,created_at=$7,updated_at=$8 WHERE id=$1"
	errs := db.QueryRow(sql, user.ID, user.Profile_url, user.Nama_users, user.Email, user.Password, user.Role, user.Create_at, time.Now())
	return errs.Err()
}
func DeleteUser(db *sql.DB, user structs.User) (err error) {
	sql := "DELETE FROM users WHERE id=$1"
	errs := db.QueryRow(sql, user.ID)
	return errs.Err()
}
