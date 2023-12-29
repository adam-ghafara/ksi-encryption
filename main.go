package ksi

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var key []byte

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// key = 32 bytes long
	key, err = hexDecodeString("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f")
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.POST("/encrypt", func(c *gin.Context) {
		plaintext := c.PostForm("plaintext")
		encryptedText, err := encrypt(plaintext)
		if err != nil {
			c.JSON(500, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"status":     "success",
			"ciphertext": encryptedText,
		})
	})

	router.POST("/decrypt", func(c *gin.Context) {
		encryptedText := c.PostForm("ciphertext")
		decryptedText, err := decrypt(encryptedText)
		if err != nil {
			c.JSON(500, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"status":    "success",
			"plaintext": decryptedText,
		})
	})

	router.Run(":8080")
}
