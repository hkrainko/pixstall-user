package auth_callback

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
		default:
			return http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)
		}
	} else {
		return http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)
	}
}
