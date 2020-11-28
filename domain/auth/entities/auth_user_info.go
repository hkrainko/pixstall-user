package entities

import "time"

type AuthUserInfo struct {
	ID       string    `json:"id"`
	AuthType string    `json:"authType"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
	PhotoURL string    `json:"picture"`
}