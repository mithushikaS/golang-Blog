package util

import (
	"time"
	"github.com/golang-jwt/jwt"
)

// SecretKey is declared in jwt.go

func GenerateJwtWithIssuer(issuer string) (string,error){
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	SecretKey := "your-secret-key"
	return claims.SignedString([]byte(SecretKey))
}
func Parsejwt(cookie string) (string,error){
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		SecretKey := "your-secret-key"
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", err
	}
	return claims.Issuer, nil
}