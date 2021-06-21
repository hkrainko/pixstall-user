package msg

import userModel "pixstall-user/domain/user/model"

type UserUpdatedEventMsg struct {
	userModel.UserUpdater `json:",inline"`
}

func NewUserUpdatedEventMsg(updater userModel.UserUpdater) UserUpdatedEventMsg {
	return UserUpdatedEventMsg{
		updater,
	}
}
