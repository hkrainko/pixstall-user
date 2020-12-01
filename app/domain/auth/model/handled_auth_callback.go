package model

import "pixstall-user/app/domain/user/model"

type HandledAuthCallback struct {
	Token string
	model.User
}