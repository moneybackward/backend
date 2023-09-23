package models

type User struct {
	Base
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
