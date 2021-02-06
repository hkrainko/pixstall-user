package register

import (
	"net/http"
	model2 "pixstall-user/domain/reg/model"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) (int, interface{}) {
	if regError, isError := err.(model2.RegError); isError {
		switch regError {
		case model2.RegErrorDuplicateUser:
			return http.StatusConflict, ErrorResponse{
				Message: regError.Error(),
			}
		case model2.RegErrorAuthIDAlreadyRegister:
			return http.StatusForbidden, ErrorResponse{
				Message: regError.Error(),
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
