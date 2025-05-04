package middleware

import (
	"errors"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
	"github.com/emanuel3k/playlist-transfer/pkg/web/response"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"time"
)

var invalidTokenError = web.UnauthorizedError("Invalid token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			secret := os.Getenv("JWT_SECRET_KEY")
			tokenValue := r.Header.Get("Authorization")

			if !strings.HasPrefix(tokenValue, "Bearer ") {
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
			if !ok && !token.Valid {
				response.Send(w, http.StatusUnauthorized, invalidTokenError)
				return
			}

			if exp, ok := claims["exp"].(float64); !ok || int64(exp) < time.Now().Unix() {
				response.Send(w, http.StatusUnauthorized, invalidTokenError)
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
