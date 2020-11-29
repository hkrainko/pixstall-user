package auth_callback

import (
	"pixstall-user/app/domain/user/model"
)

type Response struct {
	ID       string `json:"id"`
	AuthType string `json:"authType"`
	Token    string `json:"token"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
	PhotoURL string `json:"picture"`
	State    string `json:"state"`
}

func NewResponse(info *model.User) *Response {
	return &Response{
		ID:       info.ID,
		AuthType: info.AuthType,
		Token:    info.Token,
		Email:    info.Email,
		Birthday: info.Birthday,
		Gender:   info.Gender,
		PhotoURL: info.PhotoURL,
		State:    string(info.State),
	}
}
