package utils

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type TokenManager interface {
	CreateToken(id uint) (string, error)
	VerifyToken(tokenString string) (jwt.MapClaims, error)
}

type tokenManager struct {
}

func NewTokenManager() TokenManager {
	return &tokenManager{}
}

func (t *tokenManager) CreateToken(id uint) (string, error) {
	payload := jwt.MapClaims{
		"id": id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (t *tokenManager) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // Replace "secret" with your actual secret key
	})

	if err != nil {
		// Token verification failed
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		// Token is invalid
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
