package structs

import "time"

type Vendor struct {
	ID          int       `json:"vendorID"`
	Nama_vendor string    `json:"nama_vendor"`
	Alamat      string    `json:"alamat"`
	Create_at   time.Time `json:"created_at"`
	Update_at   time.Time `json:"update_at"`
}
