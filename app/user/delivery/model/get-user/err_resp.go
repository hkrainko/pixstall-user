package get_user

import (
	"net/http"
	"pixstall-user/domain/user/model"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) (int, interface{}) {
	if userError, isError := err.(model.UserError); isError {
		switch userError {
		case model.UserErrorNotFound:
			return http.StatusNotFound, ErrorResponse{
				Message: userError.Error(),
			}
		default:
			return http.StatusInternalServerError, ErrorResponse{
				Message: err.Error(),
			}
		}
	} else {
		return http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		}
	}
}
