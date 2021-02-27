package msg

import "pixstall-user/domain/commission/model"

type CommissionUserUpdater struct {
	model.CommissionUpdater `json:",inline"`
}
