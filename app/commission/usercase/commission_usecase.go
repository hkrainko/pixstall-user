package usercase

import (
	"context"
	"pixstall-user/domain/commission"
	"pixstall-user/domain/commission/model"
	msgBroker "pixstall-user/domain/msg-broker"
	"pixstall-user/domain/user"
	model2 "pixstall-user/domain/user/model"
)

type commissionUseCase struct {
	msgBrokerRepo msgBroker.Repo
	userRepo      user.Repo
}

func NewCommissionUseCase(msgBrokerRepo msgBroker.Repo, userRepo user.Repo) commission.UseCase {
	return &commissionUseCase{
		msgBrokerRepo: msgBrokerRepo,
		userRepo:      userRepo,
	}
}

func (c commissionUseCase) HandleNewCreatedCommission(ctx context.Context, comm model.Commission) error {

	requester, err := c.userRepo.GetUserDetails(ctx, comm.RequesterID)
	if err != nil {
		return c.sendCommissionUserInvalidateMsg(ctx, comm.ID, err.Error())
	}
	if requester.State != model2.UserStateActive {
		return c.sendCommissionUserInvalidateMsg(ctx, comm.ID, "Requester not active")
	}
	artist, err := c.userRepo.GetUserDetails(ctx, comm.ArtistID)
	if err != nil {
		return c.sendCommissionUserInvalidateMsg(ctx, comm.ID, err.Error())
	}
	if artist.State != model2.UserStateActive {
		return c.sendCommissionUserInvalidateMsg(ctx, comm.ID, "Artist not active")
	}

	commUsersValidation := model.CommissionUsersValidation{
		CommID:               comm.ID,
		IsValid:              true,
		InvalidReason:        nil,
		ArtistName:           &artist.UserName,
		ArtistProfilePath:    &artist.ProfilePath,
		RequesterName:        &requester.UserName,
		RequesterProfilePath: &requester.ProfilePath,
	}
	return c.msgBrokerRepo.SendCommissionUserValidationEvent(ctx, commUsersValidation)
}


// Private

func (c commissionUseCase) sendCommissionUserInvalidateMsg(ctx context.Context, commID string, reason string) error {
	commUsersValidation := model.CommissionUsersValidation{
		CommID:               commID,
		IsValid:              false,
		InvalidReason:        &reason,
	}
	return c.msgBrokerRepo.SendCommissionUserValidationEvent(ctx, commUsersValidation)
}