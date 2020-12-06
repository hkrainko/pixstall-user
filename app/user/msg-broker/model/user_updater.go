package model

import domainUserModel "pixstall-user/app/domain/user/model"

type UserUpdater struct {
	*domainUserModel.UserUpdater
}

func NewUserUpdaterFromDomainUserUpdater(u *domainUserModel.UserUpdater) *UserUpdater {
	return &UserUpdater{
		u,
	}
}