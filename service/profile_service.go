package service

import (
	"strings"

	"github.com/Pocahri/hinge-api/data"
	"github.com/pkg/errors"
)

// EditProfile assigns new value to the specified profile field
func EditProfile(acctID string, m map[string]string) error {
	if val, ok := data.Accounts[acctID]; ok {
		p := val.Profile

		for key, val := range m {
			switch strings.ToLower(key) {
			case "firstname":
				p.FirstName = val
			case "lastname":
				p.LastName = val
			case "picture":
				p.Picture = val
			case "dob":
				p.DOB = val
			case "company":
				p.Company = val
			case "education":
				p.Education = val
			case "description":
				p.Description = val
			default:
				return errors.New("Invalid profile field")
			}
		}
	} else {
		return errors.New("Profile does not exist")
	}

	return nil
}
