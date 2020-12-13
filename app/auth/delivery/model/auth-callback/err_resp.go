package auth_callback

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
				Message: "Unknown Error",
			})
		case model.UserErrorAuthIDAlreadyRegister:
			return utils.NewAPIResponse(int(model.UserErrorAuthIDAlreadyRegister), ErrorResponse{
				Message: "Unknown Error",
			})
		default:
			return utils.NewAPIResponse(99, ErrorResponse{
				Message: "Unknown Error",
			})
		}
	} else {
		return utils.NewAPIResponse(99, ErrorResponse{
			Message: "Unknown Error",
		})
	}
}
