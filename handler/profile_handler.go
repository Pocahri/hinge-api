package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Pocahri/hinge-api/data"
	"github.com/Pocahri/hinge-api/service"
)

// EditProfileRequest - modifies profile
func EditProfileRequest(w http.ResponseWriter, r *http.Request) {
	var accountID string

	if accountID = r.Header.Get("account-id"); accountID == "" {
		http.Error(w, "Account ID not provided in header", http.StatusBadRequest)
		return
	}

	var pRequestBody map[string]string

	if bodyBytes, bodyErr := ioutil.ReadAll(r.Body); bodyErr == nil {
		marshalErr := json.Unmarshal(bodyBytes, &pRequestBody)

		if marshalErr != nil {
			http.Error(w, "Error reading body", http.StatusNotFound)
			return
		}
	}

	editErr := service.EditProfile(accountID, pRequestBody)
	if editErr != nil {
		http.Error(w, editErr.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Accounts[accountID].Profile)
	return
}
