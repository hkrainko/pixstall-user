package model

type User struct {
	ID       string
	AuthID   string
	AuthType string
	Token    string
	Email    string
	Birthday string
	Gender   string
	PhotoURL string
	State    UserState
}
