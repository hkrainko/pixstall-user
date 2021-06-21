package msg

import "pixstall-user/domain/commission/model"

type CommissionUsersValidationEventMsg struct {
	model.CommissionUsersValidation `json:",inline"`
}
