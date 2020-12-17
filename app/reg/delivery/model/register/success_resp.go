package register

import (
	"net/http"
)

type SuccessResponse struct {
	APIToken string
}

func NewSuccessResponse(apiToken string) (int, interface{}) {
	return http.StatusOK, SuccessResponse{APIToken: apiToken}
}
