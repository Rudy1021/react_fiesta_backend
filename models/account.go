package models

type Account struct {
	ID       uint   `json:"id"`
	UserName string `json:"UserName"`
	Password string `json:"password"`
}
