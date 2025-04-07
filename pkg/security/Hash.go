package security

import (
	"github.com/emanuel3k/playlist-transfer/pkg/web"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, *web.AppError) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", web.InternalServerError("Error hashing password", err)
	}

	return string(hashedPassword), nil
}
