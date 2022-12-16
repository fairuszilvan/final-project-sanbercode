package structs

import "time"

type User struct {
	ID          int       `json:"userID"`
	Profile_url string    `json:"profile_url"`
	Nama_users  string    `json:"nama_user"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	Create_at   time.Time `json:"created_at"`
	Update_at   time.Time `json:"update_at"`
}
