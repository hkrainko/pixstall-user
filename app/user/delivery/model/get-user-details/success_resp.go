package get_user_details

import (
	"net/http"
	"pixstall-user/domain/user/model"
	"time"
)

type SuccessResponse struct {
	AuthID      string    `json:"authId"`
	UserID      string    `json:"userId"`
	UserName    string    `json:"userName"`
	AuthType    string    `json:"authType"`
	APIToken    string    `json:"apiToken"`
	Email       string    `json:"email"`
	Birthday    string    `json:"birthday"`
	Gender      string    `json:"gender"`
	ProfilePath string    `json:"profilePath"`
	IsArtist    bool      `json:"isArtist"`
	RegTime     time.Time `json:"regTime"`
	State       string    `json:"state"`
}

func NewSuccessResponse(user *model.AuthUser) (int, *SuccessResponse) {
	return http.StatusOK, &SuccessResponse{
		AuthID:      user.AuthID,
		UserID:      user.UserID,
		UserName:    user.UserName,
		AuthType:    user.AuthType,
		APIToken:    user.APIToken,
		Email:       user.Email,
		Birthday:    user.Birthday,
		Gender:      user.Gender,
		ProfilePath: user.ProfilePath,
		IsArtist:    user.IsArtist,
		RegTime:     user.RegTime,
		State:       string(user.State),
	}
}
