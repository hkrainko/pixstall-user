package model

import "pixstall-user/domain/user/model"

type HandledAuthCallback struct {
	Token string
	model.User
}