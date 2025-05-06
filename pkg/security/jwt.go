package security

import (
	"errors"
	"fmt"
	"github.com/emanuel3k/playlist-transfer/internal/domain"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

var (
	JwtSecretKey = "JWT_SECRET_KEY"
)

type Token struct {
	ID    string
	Name  string
	Email string
}

func GenerateToken(u *domain.User) (string, *web.AppError) {
	secret := os.Getenv(JwtSecretKey)

	claims := jwt.MapClaims{
		"id":    *u.ID,
		"name":  u.FirstName + " " + u.LastName,
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", web.InternalServerError(fmt.Errorf("error signing token: %w", err))
	}

	return tokenString, nil
}

func DecodeToken(tokenString string) (*Token, *web.AppError) {
	secret := os.Getenv(JwtSecretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, web.InternalServerError(fmt.Errorf("error parsing token: %w", err))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, web.InternalServerError(fmt.Errorf("invalid token"))
	}

	return &Token{
		ID:    claims["id"].(string),
		Name:  claims["name"].(string),
		Email: claims["email"].(string),
	}, nil
}
