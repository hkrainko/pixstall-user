package register

import (
	"net/http"
	"pixstall-user/domain/user/model"
)

type SuccessResponse struct {
	model.AuthUser
}

func NewSuccessResponse(authUser model.AuthUser) (int, interface{}) {
	return http.StatusOK, SuccessResponse{
		AuthUser: authUser,
	}
}
