package util

import (
	"boot/ent"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

// GenerateToken jwt 生成
// return 生成的 Token
func GenerateToken(user ent.User) (string, error) {
	expiresTime := time.Now().Unix() + int64(20000)
	claims := jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: expiresTime,
		Id:        strconv.Itoa(user.ID),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "boot",
		NotBefore: time.Now().Unix(),
	}
	jwtSecret := []byte("boot")
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}


