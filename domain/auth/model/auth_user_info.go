package model

type AuthUserInfo struct {
	ID       string `json:"id"`
	AuthType string `json:"authType"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
	PhotoURL string `json:"picture"`
}
