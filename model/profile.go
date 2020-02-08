package model

// Profile struct
type Profile struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Picture     string `json:"picture"`
	DOB         string `json:"dob"`
	Company     string `json:"company"`
	Education   string `json:"education"`
	Description string `json:"description"`
}
