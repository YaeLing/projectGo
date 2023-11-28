package auth

import (
	"amenBack/dbService/userDBService"
	"amenBack/model/authModel"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	userNotExist = iota
	getUserAccountsError
)

var jwtSecret = []byte("secret")

var errorMsgs = map[int]string{
	userNotExist:         "User not exist.",
	getUserAccountsError: "Get user accounts error",
}

// GenerateToken
// @Summary      Generate JWT Token
// @Description  Generate JWT Token
// @Tags         Authenticate
// @Produce      json
// @Param        id   path  string  true	"User ID"
// @Success      200  {string}  string   "Token string"
// @Failure      404  string  "User not found"
// @Failure      500  string  "Internal server error"
// @Router       /auth/{id} [get]
func GenerateToken(c *gin.Context) {
	userID := c.Param("userID")
	if results, err := userDBService.GetUserAccounts("_id", userID); err != nil {
		errorMsg := errorMsgs[userNotExist]
		log.Println(errorMsg)
		c.String(http.StatusInternalServerError, errorMsg)
	} else if len(results) == 0 {
		errorMsg := errorMsgs[userNotExist]
		log.Println(errorMsg)
		c.String(http.StatusNotFound, errorMsg)
	} else {
		account := results[0]
		claims := authModel.Claims{
			userID,
			account.Role,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
				Issuer:    "Amen Inc.",
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtSecret)
		fmt.Println(tokenString, err)
		c.String(http.StatusAccepted, tokenString)
	}
}

func Authenticate(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	tokenString := strings.Split(auth, "Bearer ")[1]
	if auth == "" {
		c.String(http.StatusForbidden, "No Authorization header provided")
		c.Abort()
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["userID"].(string)
	role := claims["role"].(string)
	if role != authModel.RoleAdmin {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Set("userID", userID)
	c.Set("role", role)
}

func Authorize(c *gin.Context) {
	if role, err := c.Get("role"); err == false || role != authModel.RoleAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
	}
}
