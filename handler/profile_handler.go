package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func getEditProfileRequest(r *http.Request) (map[string]string, error) {
	var accountID string

	if accountID = r.Header.Get("account-id"); accountID == "" {
		return nil, errors.New("Account ID not provided in header")
	}

	var pRequestBody map[string]string

	if bodyBytes, bodyErr := ioutil.ReadAll(r.Body); bodyErr == nil {
		marshalErr := json.Unmarshal(bodyBytes, &pRequestBody)

		if marshalErr != nil {
			return nil, errors.New("Error reading body")
		}
	}

	return pRequestBody, nil
}
