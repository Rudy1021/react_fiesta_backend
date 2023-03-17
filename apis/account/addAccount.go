package account

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"react_fiesta_backend/database"
	"react_fiesta_backend/models"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var db = database.Db

func AddAccount(c *gin.Context) {
	var table models.Account
	c.BindJSON(&table)
	table.ID = uuid.NewV4().String()
	passwordHash := sha256.Sum256([]byte(table.Password)) // 將密碼轉為hash
	table.Password = hex.EncodeToString(passwordHash[:])  // 因為hash是[]byte 所以轉回字串要用slice
	result := db.Debug().Create(&table)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 500,
			"data":   result.Error,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   result.Value,
		})
	}
}
