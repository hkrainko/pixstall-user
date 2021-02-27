package msg

import domainUserModel "pixstall-user/domain/user/model"

type UserUpdater struct {
	*domainUserModel.UserUpdater
}

func NewUserUpdaterFromDomainUserUpdater(u *domainUserModel.UserUpdater) *UserUpdater {
	return &UserUpdater{
		u,
	}
}