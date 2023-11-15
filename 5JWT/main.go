package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 定義角色
const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

// 存在server的秘密字串，用來最後產生signature用的
var jwtSecret = []byte("secret")

type MyClaims struct {
	UserID string `json:"userID"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// 當使用者成功登入後，應該是要生成一個jwt給他作為之後驗證使用
func generateToken(userID string) string {
	claims := MyClaims{
		userID,
		//這邊在生成的時候塞入角色
		RoleUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "HAHA",
		},
	}
	//這邊注意signingMethod是SigningMethodHS256 不是 SigningMethodES256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	fmt.Println(tokenString, err)
	return tokenString
}

func authenticate(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	tokenString := strings.Split(auth, "Bearer ")[1]
	if auth == "" {
		c.String(http.StatusForbidden, "No Authorization header provided")
		c.Abort()
		return
	}

	//去檢查整個token的完整性
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
	role := claims["role"].(string)
	//確認角色是誰才能call相對應的api
	if role != RoleAdmin {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

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
