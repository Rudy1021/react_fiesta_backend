package account

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"react_fiesta_backend/middleware"
	"react_fiesta_backend/models"

	"github.com/gin-gonic/gin"
)

func GetAccount(c *gin.Context) {

	type requestJSON struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var table models.Account
	var req requestJSON
	c.BindJSON(&req)
	passwordHash := sha256.Sum256([]byte(req.Password))                      // 將密碼轉為hash
	req.Password = hex.EncodeToString(passwordHash[:])                       // 因為hash是[]byte 所以轉回字串要用slice
	result := db.Debug().Where(`"user_name" = ?`, req.Username).Find(&table) // table是拿回來的資料
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 500,
			"data":   "帳號或密碼錯誤",
		})
	} else {
		if req.Password == table.Password {
			tokenString, _ := middleware.GenToken(req.Username)
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
				"data":   gin.H{"token": tokenString},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": 500,
				"data":   "帳號或密碼錯誤",
			})
		}
	}
}
