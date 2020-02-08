package model

// Account struct
type Account struct {
	AccountID       string   `json:"accountID"`
	Profile         *Profile `json:"profile"`
	Email           string   `json:"email"`
	PhoneNumber     string   `json:"phoneNumber"`
	Recommendations []string `json:"recommendations"`
	IncomingLikes   []string `json:"incomingLikes"`
	YourLikes       []string `json:"yourLikes"`
	Matches         []string `json:"matches"`
}
