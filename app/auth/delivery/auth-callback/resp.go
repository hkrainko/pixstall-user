package auth_callback

import (
	authModel "pixstall-user/app/domain/auth/model"
)

type Response struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
	AuthType string `json:"authType"`
	Token    string `json:"token"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
	PhotoURL string `json:"picture"`
	State    string `json:"state"`
}

func NewResponse(cb *authModel.HandledAuthCallback) *Response {
	return &Response{
		UserID:   cb.UserID,
		UserName: cb.UserName,
		AuthType: cb.AuthType,
		Token:    cb.Token,
		Email:    cb.Email,
		Birthday: cb.Birthday,
		Gender:   cb.Gender,
		PhotoURL: cb.PhotoURL,
		State:    string(cb.State),
	}
}
