package auth

import (
	"errors"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("JWT_SECRET")

type CustomClaims struct {
		UserID uint `json:"user_id"`
		jwt.RegisteredClaims
}

func GenerateToken(userID uint) (string, error) {
	
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims {
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				IssuedAt: jwt.NewNumericDate(time.Now())
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodH256, claims)
	return token.SignedString(secret)
	
}

func ParseToken(tokenStr string) (uint, error) {
	
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token)(interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
			return 0, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*CustomClaims)

	if !ok {
			return 0, errors.New("cannot parse claims")
	}
	
	return claims.UserID, nil
}
