package get_user_details

import (
	"net/http"
	"pixstall-user/domain/user/model"
)

type SuccessResponse struct {
	*model.User
}

func NewSuccessResponse(user *model.User) (int, interface{}) {
	return http.StatusOK, SuccessResponse{user}
}