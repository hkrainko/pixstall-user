package model

import (
	"strconv"
)

type UserError int

func (e UserError) Error() string {
	switch e {
	case UserErrorNotFound:
		return "UserErrorNotFound"
	case UserErrorTerminated:
		return "UserErrorTerminated"
	case UserErrorDuplicateUser:
		return "UserErrorDuplicateUser"
	case UserErrorAuthIDAlreadyRegister:
		return "UserErrorAuthIDAlreadyRegister"
	case UserErrorUnknown:
		return "UserErrorUnknown"
	default:
		return strconv.Itoa(int(e))
	}
}

const (
	UserErrorNotFound              UserError = 10
	UserErrorTerminated            UserError = 11
	UserErrorDuplicateUser         UserError = 12
	UserErrorAuthIDAlreadyRegister UserError = 13
	UserErrorUnknown               UserError = 99
)
