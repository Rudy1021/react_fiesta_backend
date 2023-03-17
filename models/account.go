package models

type Account struct {
	ID       string `json:"id"`
	UserName string `json:"UserName" gorm:"unique"`
	Password string `json:"password"`
}
