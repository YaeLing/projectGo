package main

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 存在server的秘密字串，用來最後產生signature用的
var jwtSecret = []byte("secret")

// 當使用者成功登入後，應該是要生成一個jwt給他作為之後驗證使用
func generateToken(userID string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString
}

func authenticate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		//Context.Abort使用後會中止後續的handler繼續工作
		c.Abort()
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["userID"].(string)

	c.Set("userID", userID)
}
func main() {
	server := setupRoute()
	server.Run(":8888")
}

func setupRoute() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		fmt.Printf("ctx.Request.Header: %v\n", c.Request.Header)
	}, func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	return router
}
