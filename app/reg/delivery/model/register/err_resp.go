package register

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
		case model.UserErrorDuplicateUser:
			return http.StatusConflict, ErrorResponse{
				Message: userError.Error(),
			}
		case model.UserErrorAuthIDAlreadyRegister:
			return http.StatusConflict, ErrorResponse{
				Message: userError.Error(),
			}
		default:
			return http.StatusInternalServerError, ErrorResponse{
				Message: http.StatusText(http.StatusInternalServerError),
			}
		}
	} else {
		return http.StatusInternalServerError, ErrorResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}
}
