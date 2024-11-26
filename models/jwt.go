package models

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/golang-jwt/jwt/v5"
)

type PayloadJwt struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJwtKey() error {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return fmt.Errorf("failed to generate random bytes: %w", err)
	}
	secretKey := base64.URLEncoding.EncodeToString(bytes)
	fmt.Println("Generated Secret Key:", secretKey)
	beego.AppConfig.Set("jwt_secret_key", secretKey)
	return nil
}

func GenerateToken(user *Users, Ids string, expiredSeconds int) (result map[string]interface{}, err error) {
	expirationTime := time.Now().Add(time.Hour * time.Duration(expiredSeconds))
	jwtSecret := beego.AppConfig.String("jwt_secret_key")
	payloadJwt := &PayloadJwt{
		Id: Ids,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payloadJwt) // Ubah ke SigningMethodHS256 jika tidak menggunakan ES256
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, errors.New("error: Generate token error")
	}
	result = map[string]interface{}{
		"email":        user.Email,
		"name":         user.Name,
		"access_token": tokenString,
	}

	return result, nil
}
