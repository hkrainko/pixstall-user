package get_auth_url

import "pixstall-user/app/utils"

type Response struct {
	AuthURL string `json:"authUrl"`
}

func NewResponse(authURL string) interface{} {
	return utils.NewAPIResponse(0, Response{
		AuthURL: authURL,
	})
}
