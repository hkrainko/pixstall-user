package auth_callback

import (
	authModel "pixstall-user/domain/auth/model"
)

type Response struct {
	AuthID     string `json:"authId"`
	UserName   string `json:"userName"`
	AuthType   string `json:"authType"`
	Token      string `json:"token"`
	Email      string `json:"email"`
	Birthday   string `json:"birthday"`
	Gender     string `json:"gender"`
	PhotoURL   string `json:"photoUrl"`
	State      string `json:"state"`
}

func NewResponse(cb *authModel.HandledAuthCallback) *Response {
	return &Response{
		AuthID:   cb.AuthID,
		UserName: cb.UserName,
		AuthType: cb.AuthType,
		Token:    cb.Token,
		Email:    cb.Email,
		Birthday: cb.Birthday,
		Gender:   cb.Gender,
		PhotoURL: cb.PhotoUrl,
		State:    string(cb.State),
	}
}
