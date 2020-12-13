package register

import (
	"pixstall-user/app/utils"
	"pixstall-user/domain/user/model"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) interface{} {
	if userError, isError := err.(model.UserError); isError {
		switch userError {
		case model.UserErrorDuplicateUser:
			return utils.NewAPIResponse(int(model.UserErrorDuplicateUser), ErrorResponse{
				Message: userError.Error(),
			})
		case model.UserErrorAuthIDAlreadyRegister:
			return utils.NewAPIResponse(int(model.UserErrorAuthIDAlreadyRegister), ErrorResponse{
				Message: userError.Error(),
			})
		default:
			return utils.NewAPIResponse(99, ErrorResponse{
				Message: err.Error(),
			})
		}
	} else {
		return utils.NewAPIResponse(99, ErrorResponse{
			Message: err.Error(),
		})
	}
}