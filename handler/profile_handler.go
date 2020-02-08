package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Pocahri/hinge-api/model"
	"github.com/Pocahri/hinge-api/service"
	"github.com/pkg/errors"
)

func getEditProfileRequest(r *http.Request) (string, map[string]string, error) {
	var accountID string

	if accountID = r.Header.Get("account-id"); accountID == "" {
		return "", nil, errors.New("Account ID not provided in header")
	}

	var pRequestBody map[string]string

	if bodyBytes, bodyErr := ioutil.ReadAll(r.Body); bodyErr == nil {
		marshalErr := json.Unmarshal(bodyBytes, &pRequestBody)

		if marshalErr != nil {
			return "", nil, errors.New("Error reading body")
		}
	}

	return accountID, pRequestBody, nil
}

// ModifyProfile calls getEditProfileRequest & EditProfile
func ModifyProfile(r *http.Request, acctList map[string]*model.Account) error {
	ID, m, errRequest := getEditProfileRequest(r)
	if errRequest != nil {
		return errRequest
	}

	errEdit := service.EditProfile(ID, m, acctList)
	if errEdit != nil {
		return errEdit
	}

	return nil
}
