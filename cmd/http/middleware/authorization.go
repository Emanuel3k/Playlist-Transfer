package middleware

import (
	"errors"
	"github.com/emanuel3k/playlist-transfer/pkg/security"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
	"github.com/emanuel3k/playlist-transfer/pkg/web/response"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	invalidTokenError = web.UnauthorizedError("Invalid token")
	jwtSecretKey      = "JWT_SECRET_KEY"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			secret := os.Getenv(jwtSecretKey)
			tokenValue := r.Header.Get("Authorization")

			if tokenValue == "" || !strings.HasPrefix(tokenValue, "Bearer ") {
				response.Send(w, http.StatusUnauthorized, invalidTokenError)
				return
			}
			tokenValue = tokenValue[len("Bearer "):]

			token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}

				return []byte(secret), nil
			})
			if err != nil {
				response.Send(w, http.StatusUnauthorized, invalidTokenError)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				response.Send(w, http.StatusUnauthorized, invalidTokenError)
				return
			}

			if exp, ok := claims["exp"].(float64); !ok || int64(exp) < time.Now().Unix() {
				response.Send(w, http.StatusUnauthorized, invalidTokenError)
				return
			}

			decodedToken, appErr := security.DecodeToken(tokenValue)
			if appErr != nil {
				response.Send(w, http.StatusUnauthorized, invalidTokenError)
				return
			}

			r.Header.Set("user_id", decodedToken.ID)
			r.Header.Set("user_name", decodedToken.Name)
			r.Header.Set("user_email", decodedToken.Email)

			next.ServeHTTP(w, r)
		},
	)
}
