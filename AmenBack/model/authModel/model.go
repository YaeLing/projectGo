package authModel

import "github.com/dgrijalva/jwt-go"

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type Claims struct {
	UserID string `json:"userID"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
