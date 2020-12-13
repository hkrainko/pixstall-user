package get_auth_url

import "pixstall-user/app/utils"

type SuccessResponse struct {
	AuthURL string `json:"authUrl"`
}

func NewSuccessResponse(authURL string) interface{} {
	return utils.NewAPIResponse(0, SuccessResponse{
		AuthURL: authURL,
	})
}
