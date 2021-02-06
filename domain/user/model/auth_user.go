package model

type AuthUser struct {
	APIToken string `json:"apiToken"`
	User
}
