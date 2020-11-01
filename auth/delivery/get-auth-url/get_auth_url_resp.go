package get_auth_url

type Response struct {
	AuthURL string `json:"authUrl"`
}

func NewResponse(authURL string) *Response {
	return &Response{
		AuthURL: authURL,
	}
}
