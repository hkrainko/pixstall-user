package auth_callback

import (
	authModel "pixstall-user/domain/auth/model"
)

type Response struct {
	ID       string `json:"id"`
	AuthType string `json:"authType"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
	PhotoURL string `json:"picture"`
}

func NewResponse(info *authModel.AuthUserInfo) *Response {
	return &Response{
		ID:       info.ID,
		AuthType: info.AuthType,
		Email:    info.Email,
		Birthday: info.Birthday,
		Gender:   info.Gender,
		PhotoURL: info.PhotoURL,
	}
}
