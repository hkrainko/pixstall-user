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
	case UserErrorUnknown:
		return "UserErrorUnknown"
	default:
		return strconv.Itoa(int(e))
	}
}

const (
	UserErrorNotFound              UserError = 10
	UserErrorTerminated            UserError = 11
	UserErrorUnknown               UserError = 99
)
