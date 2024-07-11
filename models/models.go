package models

type User struct {
	ID          string `json:"-"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
}
