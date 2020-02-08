package data

import (
	"github.com/Pocahri/hinge-api/model"
)

// Accounts - hardcoded data
var Accounts = map[string]*model.Account{
	"2345678": &model.Account{
		AccountID: "2345678",
		Profile: &model.Profile{
			FirstName:   "John",
			LastName:    "Doe",
			Picture:     "url1",
			DOB:         "01/01/1999",
			Company:     "ABCompany",
			Education:   "University of Nothing",
			Description: "I am very interesting please date me",
		},
		IncomingLikes: []string{
			"1234567",
			"7785640",
			"4552138",
		},
		YourLikes: []string{},
		Matches:   []string{},
	},
	"1234567": &model.Account{
		AccountID: "1234567",
		Profile: &model.Profile{
			FirstName:   "Sarah",
			LastName:    "Kerrigan",
			Picture:     "url56789",
			DOB:         "03/31/1990",
			Company:     "Mingsk and Co",
			Education:   "University of Kohral",
			Description: "Don't go terran up my heart",
		},
		IncomingLikes: []string{"7785640"},
		YourLikes: []string{
			"2345678",
			"7785640",
		},
		Matches: []string{
			"7785640",
		},
	},
	"7785640": &model.Account{
		AccountID: "7785640",
		Profile: &model.Profile{
			FirstName:   "Sett",
			LastName:    "Ionia",
			Picture:     "url567899765",
			DOB:         "01/14/1986",
			Company:     "Self-Employed",
			Description: "Being the boss is a lot better than not being the boss",
		},
		IncomingLikes: []string{},
		YourLikes: []string{
			"2345678",
			"1234567",
			"4552138",
		},
		Matches: []string{},
	},
	"4552138": &model.Account{
		AccountID: "4552138",
		Profile: &model.Profile{
			FirstName:   "Tom",
			LastName:    "Nook",
			Picture:     "url10",
			DOB:         "05/30/2001",
			Company:     "Nook Inc.",
			Description: "Yes, yes! Upstandig businessman",
		},
		IncomingLikes: []string{},
		YourLikes:     []string{},
		Matches:       []string{},
	},
}
