package http

import "pixstall-user/domain/user"

type UserController struct {
	userUseCase user.UseCase
}

func NewUserController(useCase user.UseCase) UserController {
	return UserController{
		userUseCase: useCase,
	}
}

