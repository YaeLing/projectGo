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

type MyClaims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

// 當使用者成功登入後，應該是要生成一個jwt給他作為之後驗證使用
func generateToken(userID string) string {
	claims := MyClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "Flynn",
		},
	}
	//這邊注意signingMethod是SigningMethodHS256 不是 SigningMethodES256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	fmt.Println(tokenString, err)
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
	server.Run(":8080")
}

func setupRoute() *gin.Engine {
	router := gin.Default()
	router.GET("/getToken/:userID", func(c *gin.Context) {
		userID := c.Param("userID")
		fmt.Println(userID)
		fmt.Println(generateToken(userID))
		c.JSON(200, gin.H{
			"token": generateToken(userID),
		})
	})
	//用Group去分類哪些要使用到middleware去做認證
	//參考 https://gin-gonic.com/zh-tw/docs/examples/using-middleware/
	authorize := router.Group("/authenticate")
	authorize.Use(authenticate)
	//這個括弧應該是可有可無 但就是把同個Group放在一起比較好看
	{
		authorize.GET("/", func(ctx *gin.Context) {
			ctx.String(200, "OK")
		})
	}
	return router
}
