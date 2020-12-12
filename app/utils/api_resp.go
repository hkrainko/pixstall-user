package utils

type APIResponse struct {
	ResultCode int        `json:"resultCode"`
	Content    interface{} `json:"content"`
}

func NewAPIResponse(resultCode int, content interface{}) *APIResponse {
	return &APIResponse{
		ResultCode: resultCode,
		Content: content,
	}
}
