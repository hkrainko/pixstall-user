package update_user

import (
	"net/http"
)

type SuccessResponse struct {
	UserID string `json:"userId"`
}

func NewSuccessResponse(userID string) (int, interface{}) {
	return http.StatusOK, SuccessResponse{userID}
}