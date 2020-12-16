package get_auth_url

import (
	"net/http"
)

type SuccessResponse struct {
	AuthURL string `json:"authUrl"`
}

func NewSuccessResponse(authURL string) (int, interface{}) {
	return http.StatusOK, SuccessResponse{
		AuthURL: authURL,
	}
}
