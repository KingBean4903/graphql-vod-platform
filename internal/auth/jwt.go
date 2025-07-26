package auth

import (
	"errors"
	"os"
	"fmt"

	"time"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

type CustomClaims struct {
		UserID string `json:"user_id"`
		jwt.RegisteredClaims
}

func GenerateToken(userID string) (string, error) {
	
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims {
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	
	fmt.Println("Secret JWT", secret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
	
}

func ParseToken(tokenStr string) (string, error) {
	
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token)(interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
			return "", errors.New("invalid token")
	}
	claims, ok := token.Claims.(*CustomClaims)

	
	fmt.Println("Secret JWT", secret)

	if !ok {
			return "", errors.New("cannot parse claims")
	}
	
	return claims.UserID, nil
}
