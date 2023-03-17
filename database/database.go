package database

import (
	models "react_fiesta_backend/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Db *gorm.DB
var err error

func init() {
	Db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	Db.AutoMigrate(&models.Account{})
	// defer Db.Close()
}
