package service

import (
	"github.com/Pocahri/hinge-api/model"
	"github.com/pkg/errors"
)

// EditProfile assigns new value to the specified profile field
func EditProfile(acctID string, m map[string]string, acctList map[string]*model.Account) error {
	if val, ok := acctList[acctID]; ok {
		p := val.Profile

		for key, val := range m {
			switch key {
			case "firstName":
				p.FirstName = val
			case "lastName":
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

	}

	return errors.New("Profile does not exist")
}
