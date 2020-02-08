package handler

import (
	"net/http"

	"github.com/Pocahri/model"
	"github.com/pkg/errors"
)

func getAcctLikesRequest(r *http.Request, acctList map[string]*model.Account) ([]model.Account, error) {
	var acctID string
	if acctID = r.Header.Get("account-id"); acctID == "" {
		return nil, errors.New("Account ID not provided in header")
	}

	var a model.Account
	if val, ok := acctList[acctID]; ok {
		a = val.Account
	} else {
		return nil, errors.New("Must provide an existing account ID")
	}

	return a.IncomingLikes
}
