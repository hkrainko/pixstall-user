package get_auth_url

import (
	"pixstall-user/app/utils"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) interface{} {
	return utils.NewAPIResponse(99, ErrorResponse{
		Message: "Unknown Error",
	})
}
