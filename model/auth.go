package model

// User struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthToken struct
type AuthToken struct {
	Token string `json:"token"`
}
