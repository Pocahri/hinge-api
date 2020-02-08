package model

// Account struct
type Account struct {
	AccountID       string    `json:"accountID"`
	Profile         *Profile  `json:"profile"`
	Email           string    `json:"email"`
	PhoneNumber     string    `json:"phoneNumber"`
	Recommendations []Account `json:"recommendations"`
	LikesYou        []Account `json:"likesYou"`
	YourLikes       []Account `json:"yourLikes"`
	Matches         []Account `json:"matches"`
}

// AccountService interface
type AccountService interface {
	ModifyEmail(e string)
	ModifyPhoneNum(p string)
	AddMatches(p1 []Profile, p2 []Profile)
}

// ModifyEmail changes account's email address
func (a *Account) ModifyEmail(e string) {
	a.Email = e
}

// ModifyPhoneNum changes account's phone number
func (a *Account) ModifyPhoneNum(p string) {
	a.PhoneNumber = p
}

// AddMatches is an intersection of p1 (your likes) and p2 (people who like you)
func (a *Account) AddMatches() {
	m := make(map[string]int)
	for _, accounts := range a.YourLikes {
		m[accounts.AccountID]++
	}

	for _, accounts := range a.LikesYou {
		if _, ok := m[accounts.AccountID]; ok {
			a.Matches = append(a.Matches, accounts)
		}
	}
}
