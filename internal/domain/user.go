package domain

import (
	"fmt"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

var (
	JwtSecretKey = "JWT_SECRET_KEY"
)

type User struct {
	ID        *string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (u *User) GenerateToken() (string, *web.AppError) {
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
