package auth_callback

import (
	"net/http"
	authModel "pixstall-user/domain/auth/model"
)

type SuccessResponse struct {
	AuthID     string `json:"authId"`
	UserID     string `json:"userId"`
	UserName   string `json:"userName"`
	AuthType   string `json:"authType"`
	APIToken   string `json:"apiToken"`
	RegToken   string `json:"regToken"`
	Email      string `json:"email"`
	Birthday   string `json:"birthday"`
	Gender     string `json:"gender"`
	ProfileURL string `json:"profileUrl"`
	IsArtist   bool   `json:"isArtist"`
	State      string `json:"state"`
}

func NewSuccessResponse(cb *authModel.HandledAuthCallback) (int, *SuccessResponse) {
	return http.StatusOK, &SuccessResponse{
		AuthID:     cb.AuthID,
		UserID:     cb.UserID,
		UserName:   cb.UserName,
		AuthType:   cb.AuthType,
		APIToken:   cb.APIToken,
		RegToken:   cb.RegToken,
		Email:      cb.Email,
		Birthday:   cb.Birthday,
		Gender:     cb.Gender,
		ProfileURL: cb.PhotoUrl,
		IsArtist:   cb.IsArtist,
		State:      string(cb.State),
	}
}
