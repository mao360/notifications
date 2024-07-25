package models

type User struct {
	ID          string `json:"-"`
	UserName    string `json:"user_name" mapstructure:"user_name"`
	Password    string `json:"password" mapstructure:"password"`
	DateOfBirth string `json:"date_of_birth" mapstructure:"date_of_birth"`
}
