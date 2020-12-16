package get_auth_url

import (
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) (int, interface{}) {
	return http.StatusInternalServerError, ErrorResponse{
		Message: "Unknown Error",
	}
}