package model

type User struct {
	UserID   string
	AuthID   string
	AuthType string
	Token    string
	Email    string
	Birthday string
	Gender   string
	PhotoURL string
	State    UserState
}

func (u User) Error() string {
	panic("implement me")
}

