package security

import (
	"fmt"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, *web.AppError) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", web.InternalServerError(fmt.Errorf("error hashing password: %w", err))
	}

	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword, password string) *web.AppError {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return web.UnauthorizedError("Invalid email or password")
	}

	return nil
}
