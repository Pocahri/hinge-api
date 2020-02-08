package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Pocahri/hinge-api/model"
	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken creating a token from user provided body
func CreateToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u model.User
	decodeErr := json.NewDecoder(r.Body).Decode(&u)
	if decodeErr != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"password": u.Password,
	})

	tokenString, tokenErr := token.SignedString([]byte("HingeAssignment"))
	if tokenErr != nil {
		http.Error(w, "Error creating token", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.AuthToken{Token: tokenString})
	return
}
