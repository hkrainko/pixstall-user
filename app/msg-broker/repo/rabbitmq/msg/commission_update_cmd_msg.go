package msg

import "pixstall-user/domain/commission/model"

type UpdateCommissionCmdMsg struct {
	model.CommissionUpdater `json:",inline"`
}
