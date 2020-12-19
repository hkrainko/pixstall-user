package get_user

import (
	"net/http"
	"pixstall-user/domain/user/model"
)

type SuccessResponse struct {
	UserID      string
	UserName    string
	ProfilePath string
	IsArtist    bool
}

func NewSuccessResponse(user *model.User) (int, interface{}) {
	return http.StatusOK, SuccessResponse{
		UserID:      user.UserID,
		UserName:    user.UserName,
		ProfilePath: user.ProfilePath,
		IsArtist:    user.IsArtist,
	}
}
