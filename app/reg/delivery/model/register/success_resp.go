package register

import "pixstall-user/app/utils"

type SuccessResponse struct {

}

func NewSuccessResponse() interface{} {
	return utils.NewAPIResponse(0, SuccessResponse{

	})
}