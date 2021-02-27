package commission

import (
	"golang.org/x/net/context"
	"pixstall-user/domain/commission/model"
)

type UseCase interface {
	HandleNewCreatedCommission(ctx context.Context, comm model.Commission) error
}
