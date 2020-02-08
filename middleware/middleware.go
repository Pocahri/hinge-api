package middleware

import (
	"errors"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

// ValidateToken is a wrapper method that allows validation of the auth token
func ValidateToken(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("auth-token")
		if authHeader != "" {
			token, tokenErr := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("Error validating authorization token")
				}

				return []byte("HingeAssignment"), nil
			})
			if tokenErr != nil {
				http.Error(w, tokenErr.Error(), http.StatusNotFound)
			}

			if token.Valid {
				context.Set(r, "decoded", token.Claims)
				h(w, r)
			} else {
				http.Error(w, "Invalid authorization token", http.StatusNotFound)
			}
		} else {
			http.Error(w, "An authorization token is required", http.StatusBadRequest)
			return
		}
	})
}
