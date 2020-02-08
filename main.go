package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pocahri/hinge-api/handler"
	"github.com/Pocahri/hinge-api/middleware"

	"github.com/gorilla/mux"
)

const (
	port = 8080
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/authenticate", handler.CreateToken).Methods(http.MethodPost)
	r.HandleFunc("/incoming-likes", middleware.ValidateToken(handler.GetAcctLikesRequest)).Methods(http.MethodGet)
	r.HandleFunc("/edit-profile", middleware.ValidateToken(handler.EditProfileRequest)).Methods(http.MethodPost)

	fmt.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
