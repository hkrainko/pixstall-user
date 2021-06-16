package get_user

import (
	"net/http"
	"pixstall-user/domain/user/model"
	"time"
)

type SuccessResponse struct {
	UserID      string    `json:"userId"`
	UserName    string    `json:"userName"`
	ProfilePath string    `json:"profilePath"`
	RegTime     time.Time `json:"regTime"`
	IsArtist    bool      `json:"isArtist"`
}

func NewSuccessResponse(user *model.User) (int, interface{}) {
	return http.StatusOK, SuccessResponse{
		UserID:      user.UserID,
		UserName:    user.UserName,
		ProfilePath: user.ProfilePath,
		RegTime:     user.RegTime,
		IsArtist:    user.IsArtist,
	}
}
