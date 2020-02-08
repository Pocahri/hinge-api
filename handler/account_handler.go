package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Pocahri/hinge-api/data"
)

// GetAcctLikesRequest retrives slice of incoming likes
func GetAcctLikesRequest(w http.ResponseWriter, r *http.Request) {
	var acctID string
	if acctID = r.Header.Get("account-id"); acctID == "" {
		http.Error(w, "Account ID not provided in header", http.StatusBadRequest)
		return
	}

	if val, ok := data.Accounts[acctID]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(val.IncomingLikes)
		return
	}

	http.Error(w, "Must provide an existing account ID", http.StatusBadRequest)
}
