package main

import (
	"os"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

func Sign(id uint, email string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["email"] = email
	claims["role"] = role
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// 電子署名
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}
