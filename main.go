package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pocahri/hinge-api/handler"
	"github.com/gorilla/mux"
)

const (
	port = 8080
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/incoming-likes", handler.GetAcctLikesRequest).Methods(http.MethodGet)
	r.HandleFunc("/edit-profile", handler.EditProfileRequest).Methods(http.MethodPost)

	fmt.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
