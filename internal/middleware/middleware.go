package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"net/http"
	"vk_test_task/internal/api/models"
)

func JWTAdminAuth(secret string, logger *slog.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		var claims api_models.AuthClaims

		token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || token.Valid != true || claims.IsAdmin != true {
			w.WriteHeader(http.StatusUnauthorized)
			errText := fmt.Sprintf("%s request unathorized", r.URL)
			logger.Error(errText)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func JWTUserAuth(secret string, logger *slog.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")

		var claims api_models.AuthClaims

		token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || token.Valid != true {
			w.WriteHeader(http.StatusUnauthorized)
			errText := fmt.Sprintf("%s request unathorized", r.URL)
			logger.Error(errText)
			return
		}

		next.ServeHTTP(w, r)
	}
}
