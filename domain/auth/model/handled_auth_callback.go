package model

import "pixstall-user/domain/user/model"

type HandledAuthCallback struct {
	APIToken string
	RegToken string
	model.User
	PhotoUrl string
}
